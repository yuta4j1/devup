package gitopt

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
)

func GitInit(path string) error {
	fmt.Println("[git init]")
	_, err := git.PlainInit(path, true)
	return err
}