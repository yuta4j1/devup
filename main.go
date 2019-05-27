package main

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"fmt"
	"os"
	"time"
	"log"
	"gopkg.in/urfave/cli.v1"
	. "gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/plumbing"
	"./gitopt"
)

func main() {

	var targetPath string
	app := cli.NewApp()
	app.Name = "devup"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "path",
			Usage: "starting propject path",
			Destination: &targetPath,
		},
	}
	app.Action = func(c *cli.Context) error {
		// default path is current directory
		// getPath := currentDir()
		// テスト用
		getPath := "C:\\Users\\kasca\\OneDrive\\ドキュメント\\git-test\\srctest2"
		if c.NArg() > 0 {
			getPath = c.Args()[0]
		}
		log.Println("getPath: ", getPath)
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

		hash, err := gitopt.GitCommit(workTree, "first commit", &CommitOptions{
			All: true,
			Author: &object.Signature{
				Name: "yuta4j1",
				Email: "kascado.ys10@gmail.com",
				When: time.Now(),
			},
		})
		fmt.Println(hash)
		// アクセストークンを取得する
		gitopt.FetchAccessToken()
		// プロジェクト名から、同名リポジトリがリモートリポジトリに存在するかをチェックする
		// 存在する場合、リモートリポジトリに既存プロジェクトが存在するため、
		// 処理の続行 or 中断を選択するコマンドを表示する

		// 同名リポジトリがない場合、リモートリポジトリを作成し、git pushする。

		

		// TODO git commit
		// TODO create remote repository
		
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func currentDir() string {
	cur, _ := os.Getwd()
	return cur
}