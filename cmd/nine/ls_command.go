package main

import (
	"fmt"
	"os"
)

type LsCmd struct {
	List bool   `arg:"-l" help:"print detailed list"`
	Path string `arg:"positional"`
}

func lsCommand(args *Args) {
	cli := NewClient(args)

	stats, err := cli.Readdir(args.Ls.Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not list directory '%s': %s", args.Ls.Path, err)
		os.Exit(1)
	}

	for _, stat := range stats {
		fmt.Println(stat.Name)
	}
}
