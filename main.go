package main

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"fmt"
	"path"
	"os"
	"time"
	"log"
	"gopkg.in/urfave/cli.v1"
	. "gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/plumbing"
	"github.com/google/go-github/github"
	"./gitopt"
)

func main() {

	var targetPath string
	var accessToken string
	app := cli.NewApp()
	app.Name = "devup"
	app.Flags = []cli.Flag {
		// path (absolute path)
		cli.StringFlag{
			Name: "path, p",
			Usage: "starting propject path",
			Destination: &targetPath,
		},
		// github access token
		cli.StringFlag{
			Name: "token, t",
			Usage: "your github account access token",
			Destination: &accessToken,
		},
	}
	app.Action = func(c *cli.Context) error {
		// default path is current directory
		// if targetPath == "" {
		// 	targetPath = currentDir()
		// }
		// 動作確認用
		getPath := "C:\\Users\\kasca\\OneDrive\\ドキュメント\\git-test\\srctest2"
		projName := projectName(getPath)
		// git init
		repo, err := gitopt.GitInit(getPath)
		if err != nil {
			log.Fatal(err)
			// if target repository is already initialized, advance the process regardless.
			// return nil
		}
		// get Worktree
		workTree, err := repo.Worktree()
		if err != nil {
			log.Fatal(err)
			return nil
		}
		// TODO .gitignoreで指定したファイルを除外できているか動作確認
		addErr := gitopt.GitAdd(workTree, ".")
		if addErr != nil {
			log.Fatal(err)
		}
		// git commit
		// don't use hash at this point
		_, err = gitopt.GitCommit(workTree, "first commit", &CommitOptions{
			All: true,
			Author: &object.Signature{
				Name: "yuta4j1",
				Email: "kascado.ys10@gmail.com",
				When: time.Now(),
			},
		})
		// gitopt.FetchAccessToken()
		// Initialize github client object
		githubClient, ctx := gitopt.InitClient(accessToken)
		repos, _, err := githubClient.Repositories.List(ctx, "", nil)
		// verify whether there is a project with the same name as local repository in the remote repository
		for _, repo := range repos {
			if projName == *repo.Name {
				fmt.Println("[abort] A same name project already exists in remote repository.")
				return nil
			}
		}

		// create new project at remote repository
		newRepo, _, err := githubClient.Repositories.Create(ctx, "", &github.Repository{
			Name: &projName,
		})
		fmt.Println("[created remote repository!]")
		fmt.Println("repo ID: ", *newRepo.ID)
		fmt.Println("repo FullName: ", *newRepo.FullName)
		fmt.Println("repo MasterBranch: ", *newRepo.MasterBranch)
		fmt.Println("repo CreatedAt: ", *newRepo.CreatedAt)
		fmt.Println("repo CloneURL: ", *newRepo.CloneURL)

		// 同名リポジトリがない場合、リモートリポジトリを作成し、git pushする。
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// get current directory path
func currentDir() string {
	cur, _ := os.Getwd()
	return cur
}

func projectName(dirPath string) string {
	return path.Base(dirPath)
}