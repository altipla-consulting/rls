package main

import (
	"fmt"
	"log"
	"os"

	"github.com/juju/errors"
	"github.com/libgit2/git2go"
)

const releaseBranchName = "release"

func main() {
	if err := run(); err != nil {
		log.Fatal(errors.ErrorStack(err))
	}
}

func run() error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Trace(err)
	}

	repo, err := git.OpenRepository(wd)
	if err != nil {
		return errors.Trace(err)
	}
	defer repo.Free()

	head, err := repo.Head()
	if err != nil {
		return errors.Trace(err)
	}
	defer head.Free()

	lastCommit, err := repo.LookupCommit(head.Target())
	if err != nil {
		return errors.Trace(err)
	}
	defer lastCommit.Free()

	branch, err := repo.LookupBranch(releaseBranchName, git.BranchLocal)
	if err != nil {
		if git.IsErrorCode(err, git.ErrNotFound) {
			branch, err = repo.CreateBranch(releaseBranchName, lastCommit, false)
			if err != nil {
				return errors.Trace(err)
			}
		} else {
			return errors.Trace(err)
		}
	}
	defer branch.Free()

	log.Printf("Message:  %s", lastCommit.Message())
	log.Printf("Commit:   %s", head.Target())
	log.Printf("Author:   %s <%s>", lastCommit.Author().Name, lastCommit.Author().Email)
	log.Printf("Date:     %s", lastCommit.Author().When.Format("Mon 02 Jan 2006 15:04:05 -0700"))

	fmt.Printf("Is this correct? (Y/n): ")
	if !askForConfirmation() {
		return nil
	}

	modifiedBranch, err := branch.SetTarget(head.Target(), lastCommit.Message())
	if err != nil {
		return errors.Trace(err)
	}
	defer modifiedBranch.Free()

	remote, err := repo.Remotes.Lookup("origin")
	if err != nil {
		return errors.Trace(err)
	}
	defer remote.Free()

	callbacks := git.RemoteCallbacks{
		CredentialsCallback: func(url string, usernameFromURL string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
			code, credentials := git.NewCredSshKeyFromAgent(usernameFromURL)
			return git.ErrorCode(code), &credentials
		},
		CertificateCheckCallback: func(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
			return 0
		},
	}
	refs := []string{
		fmt.Sprintf("refs/heads/%[1]s:refs/heads/%[1]s", releaseBranchName),
		"refs/heads/master:refs/heads/master",
	}
	if err := remote.Push(refs, &git.PushOptions{RemoteCallbacks: callbacks}); err != nil {
		return errors.Trace(err)
	}

	log.Println("New version released successfully!")

	return nil
}

func askForConfirmation() bool {
	var response string
	if _, err := fmt.Scanln(&response); err != nil {
		return true
	}

	if response == "y" {
		return true
	} else if response == "n" {
		return false
	}

	fmt.Printf("Please type y or n and then press enter: ")
	return askForConfirmation()
}
