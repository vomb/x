package git

import (
	"context"
)

type Location struct {
	Git      string
	Ref      string
	PathGlob string
}

type Change struct {
	Git string
	Ref string
}

func Subscribe(ctx context.Context, req Location) (*Change, error) {
	return nil, nil
}
