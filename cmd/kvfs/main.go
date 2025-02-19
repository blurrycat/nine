package main

import (
	"fmt"
	"os"

	"blurrycat.dev/nine/pkg/nineutils"
	"github.com/knusbaum/go9p"
	"github.com/knusbaum/go9p/fs"
)

func main() {
	username, group, err := nineutils.GetCurrentUser()
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not get local user/group:", err)
	}

	fs, _ := fs.NewFS(username, group, 0777,
		fs.WithCreateFile(fs.CreateStaticFile),
		fs.WithCreateDir(fs.CreateStaticDir),
		fs.WithRemoveFile(fs.RMFile),
	)

	go9p.Serve("127.0.0.1:9999", fs.Server())
}
