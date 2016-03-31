package rls

import (
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/juju/errors"
	"github.com/libgit2/git2go"
)

const (
	Patch = "patch"
	Minor = "minor"
	Major = "major"
)

func Release(kind string) error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Trace(err)
	}

	repo, err := git.OpenRepository(wd)
	if err != nil {
		return errors.Trace(err)
	}
	defer repo.Free()

	tags, err := repo.Tags.List()
	if err != nil {
		return errors.Trace(err)
	}

	last := semver.MustParse("0.0.0")
	for _, tag := range tags {
		version, err := semver.Parse(tag)
		if err != nil {
			return errors.Trace(err)
		}

		if version.GT(last) {
			last = version
		}
	}

	switch kind {
	case Major:
		last.Major++

	case Minor:
		last.Minor++

	case Patch:
		last.Patch++

	default:
		return errors.Errorf("undefined release kind: %s", kind)
	}

	head, err := repo.Head()
	if err != nil {
		return errors.Trace(err)
	}
	defer head.Free()

	commit, err := repo.LookupCommit(head.Target())
	if err != nil {
		return errors.Trace(err)
	}
	defer commit.Free()

	log.Printf("Release:  %s", kind)
	log.Printf("Version:  %s", last)
	log.Printf("Commit:   %s", head.Target())
	log.Printf("Message:  %s", commit.Message())
	log.Printf("Author:   %s <%s>", commit.Author().Name, commit.Author().Email)
	log.Printf("Date:     %s", commit.Author().When.Format("Mon 02 Jan 2006 15:04:05 -0700"))

	fmt.Printf("Is this correct? (Y/n): ")
	if !askForConfirmation() {
		return nil
	}

	signature, err := repo.DefaultSignature()
	if err != nil {
		return errors.Trace(err)
	}

	if _, err := repo.Tags.Create(last.String(), commit, signature, fmt.Sprintf("Release v%s\n", last)); err != nil {
		return errors.Trace(err)
	}

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
		"refs/heads/master:refs/heads/master",
		fmt.Sprintf("refs/tags/%[1]s:refs/tags/%[1]s", last),
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
