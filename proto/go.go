package proto

import (
	"context"
)

type Location struct {
	Path string
}

type ToGoRequest struct {
	Location
	Package string
}

func ToGo(ctx context.Context, req ToGoRequest) (*Location, error) {
	return nil, nil
}
