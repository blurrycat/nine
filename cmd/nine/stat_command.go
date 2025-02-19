package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/knusbaum/go9p/proto"
)

const (
	QTDIR    = 0x80
	QTAPPEND = 0x40
	QTEXCL   = 0x20
	QTAUTH   = 0x08
)

func qidTypeString(qid *proto.Qid) string {
	var buf bytes.Buffer

	if qid.Qtype&QTDIR != 0 {
		buf.WriteByte('d')
	}

	if qid.Qtype&QTAPPEND != 0 {
		buf.WriteByte('a')
	}

	if qid.Qtype&QTEXCL != 0 {
		buf.WriteByte('l')
	}

	if qid.Qtype&QTAUTH != 0 {
		buf.WriteByte('A')
	}

	if buf.Len() > 0 {
		return fmt.Sprintf(" (%s)", buf.String())
	}

	return ""
}

type StatCmd struct {
	Path string `arg:"positional"`
}

func statCommand(args *Args) {
	cli, path := NewClient(args)

	stat, err := cli.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not stat path '%s': %s\n", path, err)
		os.Exit(1)
	}

	fmt.Printf(
		"name:\t%s\n"+
			"uid:\t%s\n"+
			"gid:\t%s\n"+
			"muid:\t%s\n"+
			"qid\t[%s]%s\n"+
			"mode:\t%O\n"+
			"atime:\t%d\n"+
			"mtime:\t%d\n"+
			"length:\t%d\n"+
			"type:\t%d\n"+
			"dev:\t%d\n",
		stat.Name,
		stat.Uid,
		stat.Gid,
		stat.Muid,
		&stat.Qid, qidTypeString(&stat.Qid),
		stat.Mode,
		stat.Atime,
		stat.Mtime,
		stat.Length,
		stat.Type,
		stat.Dev,
	)
}
