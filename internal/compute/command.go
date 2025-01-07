package compute

import (
	"context"

	"github.com/khulnasoft/khulnasoft/internal/gitserver"
	"github.com/khulnasoft/khulnasoft/internal/search/result"
)

type Command interface {
	command()
	// Run transforms r into a computed Result.
	//
	// Note: It takes a gitserver client since the replace action needs to
	// request the full file contents.
	Run(ctx context.Context, gitserverClient gitserver.Client, r result.Match) (Result, error)
	ToSearchPattern() string
	String() string
}

var (
	_ Command = (*MatchOnly)(nil)
	_ Command = (*Replace)(nil)
	_ Command = (*Output)(nil)
)

func (MatchOnly) command() {}
func (Replace) command()   {}
func (Output) command()    {}
