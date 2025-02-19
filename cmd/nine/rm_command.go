package main

import (
	"fmt"
	"os"
)

type RmCmd struct {
	Path string `arg:"positional"`
}

func rmCommand(args *Args) {
	cli, path := NewClient(args)

	err := cli.Remove(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not remove path '%s': %s\n", path, err)
		os.Exit(1)
	}
}
