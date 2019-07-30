package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/kyokomi/emoji"
	"github.com/motemen/go-gitconfig"
	"github.com/yuta4j1/devup/gitopt"
	"golang.org/x/crypto/ssh/terminal"
	. "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/urfave/cli.v1"
)

type User struct {
	UserName string `gitconfig:"user.name"`
	Email    string `gitconfig:"user.email"`
}

func main() {

	cli.AppHelpTemplate = `NAME:
	{{.Name}} - {{.Usage}}
 USAGE:
	{{.Name}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}}{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if len .Authors}}
 AUTHOR:
	{{range .Authors}}{{ . }}{{end}}
	{{end}}{{if .Commands}}
 COMMANDS:
 {{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
 GLOBAL OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright }}
 COPYRIGHT:
	{{.Copyright}}
	{{end}}{{if .Version}}
 VERSION:
	{{.Version}}
	{{end}}
	`
	var targetPath string
	var accessToken string
	app := cli.NewApp()
	app.Name = "devup"
	app.Usage = "Start development on GitHub with a single command line."
	app.Flags = []cli.Flag{
		// path (absolute path)
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "Target project path. If you don't specify it, target current directory.",
			Destination: &targetPath,
		},
		// github access token
		cli.StringFlag{
			Name:        "token, t",
			Usage:       "Your github account access token. This flag is required.",
			Destination: &accessToken,
		},
	}
	app.Action = func(c *cli.Context) error {
		// default path is current directory
		if targetPath == "" {
			targetPath = currentDir()
		}
		projName := projectName(targetPath)
		fmt.Println("Project path to be initialized: ", targetPath)
		// git init
		repo, err := gitopt.GitInit(targetPath)
		if err != nil {
			log.Fatal(err)
			// if target repository is already initialized, advance the process regardless.
			// return nil
		}
		fmt.Println("✔ git init")
		// get Worktree
		workTree, err := repo.Worktree()
		if err != nil {
			log.Fatal(err)
			return nil
		}
		addErr := gitopt.GitAdd(workTree, ".")
		if addErr != nil {
			log.Fatal(err)
		}
		fmt.Println("✔ git Add")
		// git commit
		// don't use hash at this point
		// get git-config info
		var config User
		err = gitconfig.Load(&config)
		_, err = gitopt.GitCommit(workTree, "first commit", &CommitOptions{
			All: true,
			Author: &object.Signature{
				Name:  config.UserName,
				Email: config.Email,
				When:  time.Now(),
			},
		})
		fmt.Println("✔ git Commit")
		// Initialize github client object
		githubClient, ctx := gitopt.InitClient(accessToken)

		// TODO validation access token
		repos, _, err := githubClient.Repositories.List(ctx, "", nil)
		// get account name
		var accountName string
		if len(repos) != 0 {
			accountName = *repos[0].Owner.Login
		}
		// verify whether there is a project with the same name as local repository in the remote repository
		for _, repo := range repos {
			if projName == *repo.Name {
				fmt.Println("[abort] A same name project already exists in remote repository.")
				return nil
			}
		}

		// create new project at remote repository
		newRepo, _, err := gitopt.GithubCreateRepository(githubClient, ctx, projName)
		fmt.Println("✔ created remote repository")
		fmt.Println("  ID: ", *newRepo.ID)
		fmt.Println("  FullName: ", *newRepo.FullName)
		fmt.Println("  CreatedAt: ", *newRepo.CreatedAt)
		fmt.Println("  CloneURL: ", *newRepo.CloneURL)

		// git remote add
		remoteRepo, err := gitopt.GitCreateRemote(*repo, "https://github.com/"+accountName+"/"+projName+".git")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Please input your GitHub account.")
		fmt.Println("Username: ")
		userName := prompt.Input(": ", makeCompleter(
			prompt.Suggest{Text: "username", Description: "your username for authentication"}))
		fmt.Println("Password: ")
		// TODO mask password input
		// password := prompt.Input(": ", makeCompleter(
		// 	prompt.Suggest{Text: "password", Description: "your passwords for authentication"}))
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		// create 'master' branch
		err = gitopt.GitCreateBranch(*repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("✔ created 'master' branch")
		// git push
		err = gitopt.GitPush(remoteRepo, ctx, userName, string(password))
		if err != nil {
			log.Fatal("[git push]", err)
		}
		fmt.Println("✔ git Push")

		fmt.Print("Done successfully!")
		emoji.Println(":tada:")

		return nil
	}
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "yuta4j1",
			Email: "kascado.ys10@gmail.com",
		},
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

// get current directory path
func currentDir() string {
	cur, _ := os.Getwd()
	return cur
}

// get project name from a given path
func projectName(dirPath string) string {
	return filepath.Base(dirPath)
}
