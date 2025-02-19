package main

import (
	"fmt"
	"io"
	"os"

	"github.com/knusbaum/go9p/proto"
)

type ReadCmd struct {
	Path string `arg:"positional"`
}

func readCommand(args *Args) {
	cli, path := NewClient(args)

	f, err := cli.Open(path, proto.Oread)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open path '%s' for reading: %s\n", path, err)
		os.Exit(1)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	fmt.Print(string(bytes))
}
