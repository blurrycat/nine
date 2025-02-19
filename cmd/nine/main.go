package main

import "github.com/alexflint/go-arg"

type Args struct {
	Addr string `arg:"-a"`

	Ls   *LsCmd   `arg:"subcommand:ls" help:"print a directory listing"`
	Read *ReadCmd `arg:"subcommand:read" help:"print contents of path"`
}

func main() {
	var args Args
	p := arg.MustParse(&args)
	if p.Subcommand() == nil {
		p.Fail("missing subcommand")
	}

	switch {
	case args.Ls != nil:
		lsCommand(&args)
	}
}
