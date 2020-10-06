package cmd

import (
	"context"
	"io"
)

type contextKey int

const (
	cmdReaderKey contextKey = iota
)

// ContextWithReader returns a Context with a given io.Reader.
func ContextWithReader(ctx context.Context, r io.Reader) context.Context {
	return context.WithValue(ctx, cmdReaderKey, r)
}
