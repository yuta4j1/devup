package gitopt

import (
	"fmt"
	"context"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/config"
)

// git init operation
func GitInit(path string) (*git.Repository, error) {
	fmt.Println("[git init]")
	repo, err := git.PlainInit(path, false)
	return repo, err
}

func ExistsMasterBranch(r *git.Repository) (*config.Branch, error) {
	b, err := r.Branch("master")
	return b, err
}

// git add operation
func GitAdd(w *git.Worktree, glob string) error {
	fmt.Println("[git Add]")
	err := w.AddGlob(glob)
	return err
}

// git commit operation
func GitCommit(wt *git.Worktree, msg string, opt *git.CommitOptions) (plumbing.Hash, error) {
	fmt.Println("[git commit]")
	ph, err := wt.Commit(msg, opt)
	return ph, err
}

// git 'remote add' operation
func GitCreateRemote(r git.Repository, cloneUrl string) (*git.Remote, error) {
	fmt.Println("[cloneUrl]", cloneUrl)
	remote, err := r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{cloneUrl},
	})

	return remote, err
}

func GitCreateBranch(r git.Repository) error {
	err := r.CreateBranch(&config.Branch{
		Name: "master",
		Remote: "origin",
		Merge: "refs/heads/master",
	})
	return err
}

// git 'push' operation
func GitPush(r *git.Remote, ctx context.Context, userName string, password string) error {
	err := r.Push(&git.PushOptions{
		RemoteName: r.Config().Name,
		RefSpecs: r.Config().Fetch,
		Auth: &http.BasicAuth{
			Username: userName,
			Password: password,
		},
	})

	return err
}