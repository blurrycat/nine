package main

import (
	"fmt"
	"os"
)

type CreateCmd struct {
	Path string `arg:"positional"`
}

func createCommand(args *Args) {
	cli, path := NewClient(args)

	f, err := cli.Create(path, 0o666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file at path '%s': %s\n", path, err)
		os.Exit(1)
	}
	f.Close()
}
