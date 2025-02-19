package main

import (
	"fmt"
	"os"
)

type CreateCmd struct {
	Dir  bool   `arg:"-d" help:"create directory instead of file"`
	Path string `arg:"positional"`
}

func createCommand(args *Args) {
	cli, path := NewClient(args)

	perms := 0o666
	if args.Create.Dir {
		perms = 0o777 | DMDIR
	}

	f, err := cli.Create(path, os.FileMode(perms))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file at path '%s': %s\n", path, err)
		os.Exit(1)
	}
	f.Close()
}
