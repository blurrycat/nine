package main

type ReadCmd struct {
	Path string `arg:"positional"`
}

func readCommand(args *Args) {
}
