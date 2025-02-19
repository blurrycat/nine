package main

import (
	"fmt"
	"os"
	"strings"

	"blurrycat.dev/nine/pkg/nineutils"
	"github.com/knusbaum/go9p"
	"github.com/knusbaum/go9p/fs"
)

func main() {
	username, group, err := nineutils.GetCurrentUser()
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not get local user/group:", err)
	}

	greetFS, root := fs.NewFS(username, group, 0777)

	greetDir := fs.NewStaticDir(greetFS.NewStat("greet", username, group, 0777))

	nameFile := fs.NewStaticFile(greetFS.NewStat("name", username, group, 0666), []byte(username))
	greetDir.AddChild(nameFile)

	greetDir.AddChild(
		fs.NewDynamicFile(greetFS.NewStat("hello", username, group, 0444), func() []byte {
			name := string(nameFile.Data)
			name = strings.TrimSpace(name)
			return []byte("Hello, " + name + "!")
		}),
	)

	root.AddChild(greetDir)

	go9p.PostSrv("greetfs", greetFS.Server())
}
