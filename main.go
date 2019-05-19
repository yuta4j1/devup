package main

import (
	"os"
	"log"
	cli "gopkg.in/urfave/cli.v1"
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
		getPath := "C:\\Users\\kasca\\OneDrive\\ドキュメント\\sandbox-Go\\devup-test"
		if c.NArg() > 0 {
			getPath = c.Args()[0]
		}
		log.Println("getPath: ", getPath)
		// git init
		err := gitopt.GitInit(getPath)
		if err != nil {
			log.Fatal(err)
			return nil
		}
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