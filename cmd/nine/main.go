package main

import "github.com/alexflint/go-arg"

type Args struct {
	Addr string `arg:"-a"`

	Ls     *LsCmd     `arg:"subcommand:ls" help:"print a directory listing"`
	Read   *ReadCmd   `arg:"subcommand:read" help:"print contents of path"`
	Stat   *StatCmd   `arg:"subcommand:stat" help:"print path metadata"`
	Rm     *RmCmd     `arg:"subcommand:rm" help:"remove path"`
	Create *CreateCmd `arg:"subcommand:create" help:"create file at path"`
	Write  *WriteCmd  `arg:"subcommand:write" help:"write to file at path"`
}

func (a *Args) Path() string {
	switch {
	case a.Ls != nil:
		return a.Ls.Path
	case a.Read != nil:
		return a.Read.Path
	case a.Stat != nil:
		return a.Stat.Path
	case a.Rm != nil:
		return a.Rm.Path
	case a.Create != nil:
		return a.Create.Path
	case a.Write != nil:
		return a.Write.Path
	default:
		return ""
	}
}

func (a *Args) Exec() {
	switch {
	case a.Ls != nil:
		lsCommand(a)

	case a.Read != nil:
		readCommand(a)

	case a.Stat != nil:
		statCommand(a)

	case a.Rm != nil:
		rmCommand(a)

	case a.Create != nil:
		createCommand(a)

	case a.Write != nil:
		writeCommand(a)
	}
}

func main() {
	var args Args
	p := arg.MustParse(&args)
	if p.Subcommand() == nil {
		p.Fail("missing subcommand")
	}

	args.Exec()
}
