package dir

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type GitLocation struct {
	Git  string
	Ref  string
	Path string
}

type Location struct {
	Path string
}

var ErrNotADir = errors.New("not a dir")

func FromGit(ctx context.Context, req GitLocation) (*Location, error) {
	tempDir, err := ioutil.TempDir(os.TempDir(), "from-git")
	if err != nil {
		return nil, err
	}

	if _, err := git.PlainCloneContext(ctx, tempDir, false, &git.CloneOptions{
		ReferenceName: plumbing.ReferenceName(req.Ref),
		URL:           req.Git,
		Progress:      os.Stdout,
	}); err != nil {
		return nil, err
	}

	res := &Location{
		Path: path.Join(tempDir, req.Path),
	}
	if dirInfo, err := os.Stat(res.Path); err != nil {
		return nil, fmt.Errorf("failed check cloned git: %w", err)
	} else if !dirInfo.IsDir() {
		return nil, fmt.Errorf("failed check cloned git: %v/%v: %w", req.Git, req.Path, ErrNotADir)
	}

	return res, nil
}

type ToGitRequest struct {
	GitLocation
	DeletePattern string
	FromPath string
}

func ToGit(ctx context.Context, req ToGitRequest) error {
	return nil
}
