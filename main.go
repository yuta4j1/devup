package main

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"fmt"
	"path/filepath"
	"os"
	"time"
	"log"
	"github.com/c-bata/go-prompt"
	"gopkg.in/urfave/cli.v1"
	. "gopkg.in/src-d/go-git.v4"
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
		// git commitしたユーザ名のアカウントが存在するかどうかを確認する
		githubClient, ctx := gitopt.InitClient(accessToken)
		user, _, err := githubClient.Users.Get(ctx, "yuta4j1")
		fmt.Println("[URL]", *user.URL)
		fmt.Println("[URL]", *user.ReposURL)

		// TODO validation access token
		// user名の指定
		repos, _, err := githubClient.Repositories.List(ctx, "", nil)
		// verify whether there is a project with the same name as local repository in the remote repository
		for _, repo := range repos {
			if projName == *repo.Name {
				fmt.Println("[abort] A same name project already exists in remote repository.")
				return nil
			}
		}

		// create new project at remote repository
		newRepo, _, err := gitopt.GithubCreateRepository(githubClient, ctx, projName)
		fmt.Println("[created remote repository!]")
		fmt.Println("repo ID: ", *newRepo.ID)
		fmt.Println("repo FullName: ", *newRepo.FullName)
		fmt.Println("repo CreatedAt: ", *newRepo.CreatedAt)
		fmt.Println("repo CloneURL: ", *newRepo.CloneURL)
		

		// git remote add
		remoteRepo, err := gitopt.GitCreateRemote(*repo, "https://github.com/yuta4j1/" + projName + ".git")
		if err != nil {
			log.Fatal(err)
		}

		err = gitopt.GitCreateBranch(*repo)
		if err != nil {
			log.Fatal(err)
		}
		bConf, err := gitopt.ExistsMasterBranch(repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[branch config] Name: ", bConf.Name)
		fmt.Println("[branch config] Remote: ", bConf.Remote)


		fmt.Println("Please input your username.")
		userName := prompt.Input(": ", makeCompleter(
			prompt.Suggest{Text: "username", Description: "your username for authentication"}))
		fmt.Println("your UserName: ", userName)
		fmt.Println("Please input your password.")
		password := prompt.Input(": ", makeCompleter(
			prompt.Suggest{Text: "password", Description: "your passwords for authentication"}))
		fmt.Println("your Password: ", password)

		// git push
		err = gitopt.GitPush(remoteRepo, ctx, userName, password)
		if err != nil {
			log.Fatal("[git push]", err)
		}		

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func makeCompleter(s ...prompt.Suggest) (f func(prompt.Document) []prompt.Suggest) {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}
}

// func completer(d prompt.Document) []prompt.Suggest {
// 	s := []prompt.Suggest{
// 		{Text: "username", Description: "your username for authentication"},
// 		{Text: "password", Description: "your password for authentication"},
// 	}
// 	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
// }

// get current directory path
func currentDir() string {
	cur, _ := os.Getwd()
	return cur
}

func projectName(dirPath string) string {
	return filepath.Base(dirPath)
}