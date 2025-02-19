package main

import (
	"fmt"
	"io"
	"os"

	"github.com/knusbaum/go9p/proto"
)

const MAX_SIZE = 512 * 1024 // 512KiB

type WriteCmd struct {
	Path string `arg:"positional"`
}

func writeCommand(args *Args) {
	cli, path := NewClient(args)

	f, err := cli.Open(path, proto.Owrite)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open path '%s' for writing: %s\n", path, err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = io.CopyN(f, os.Stdin, MAX_SIZE)
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "error occured while writing to path '%s': %s\n", path, err)
		os.Exit(1)
	}
}
