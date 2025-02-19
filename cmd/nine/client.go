package main

import (
	"fmt"
	"os"

	"blurrycat.dev/nine/pkg/nineutils"
	"github.com/knusbaum/go9p/client"
)

func NewClient(args *Args) (*client.Client, string) {
	username, _, err := nineutils.GetCurrentUser()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get current user: %s\n", err)
		os.Exit(1)
	}

	var cli *client.Client
	isUnix := false
	path := args.Path()

	switch args.Addr {
	case "":
		cli, err = nineutils.NewUnixClient(path, username, "")
		isUnix = true
	default:
		cli, err = nineutils.NewTCPClient(args.Addr, username, "")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to server: %s\n", err)
		os.Exit(1)
	}

	if isUnix {
		path = nineutils.PathForUnixClient(path)
	}
	if path[0] != '/' {
		path = fmt.Sprintf("/%s", path)
	}

	return cli, path
}
