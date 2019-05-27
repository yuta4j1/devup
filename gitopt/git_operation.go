package gitopt

import (
	"fmt"
	. "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// git init operation
func GitInit(path string) (*Repository, error) {
	fmt.Println("[git init]")
	repo, err := PlainInit(path, false)
	return repo, err
}

// git add operation
func GitAdd(w *Worktree, glob string) error {
	fmt.Println("[git Add]")
	err := w.AddGlob(glob)
	return err
}

// git commit operation
func GitCommit(wt *Worktree, msg string, opt *CommitOptions) (plumbing.Hash, error) {
	ph, err := wt.Commit(msg, opt)
	return ph, err
}