package main

import (
	"fmt"
	"os"
	"os/user"

	"blurrycat.dev/nine/pkg/nineutils"
	"github.com/knusbaum/go9p/client"
)

func getUsername() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.Username, nil
}

func NewClient(args *Args) *client.Client {
	username, err := getUsername()

	var cli *client.Client

	switch args.Addr {
	case "":
		cli, err = nineutils.NewUnixClient(args.Ls.Path, username, "")
	default:
		cli, err = nineutils.NewTCPClient(args.Addr, username, "")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to server: %s", err)
		os.Exit(1)
	}

	return cli
}
