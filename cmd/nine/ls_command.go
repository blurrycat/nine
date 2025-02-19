package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

type LsCmd struct {
	List bool   `arg:"-l" help:"print detailed list"`
	Path string `arg:"positional"`
}

func lsCommand(args *Args) {
	cli, path := NewClient(args)

	stats, err := cli.Readdir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not list directory '%s': %s\n", path, err)
		os.Exit(1)
	}

	for _, stat := range stats {
		if !args.Ls.List {
			fmt.Println(stat.Name)
			continue
		}

		date := time.Unix(int64(stat.Mtime), 0).Format(time.RFC822Z)
		fmt.Printf("%s M %d %s %s %d %s %s\n", modeString(stat.Mode), stat.Dev, stat.Uid, stat.Gid, stat.Length, date, stat.Name)
	}
}

var modes = [...]string{
	"---",
	"--x",
	"-w-",
	"-wx",
	"r--",
	"r-x",
	"rw-",
	"rwx",
}

const (
	DMDIR       = 0x80000000 /* mode bit for directories */
	DMAPPEND    = 0x40000000 /* mode bit for append only files */
	DMEXCL      = 0x20000000 /* mode bit for exclusive use files */
	DMMOUNT     = 0x10000000 /* mode bit for mounted channel */
	DMAUTH      = 0x08000000 /* mode bit for authentication file */
	DMTMP       = 0x04000000 /* mode bit for non-backed-up file */
	DMSYMLINK   = 0x02000000 /* mode bit for symbolic link (Unix, 9P2000.u) */
	DMDEVICE    = 0x00800000 /* mode bit for device file (Unix, 9P2000.u) */
	DMNAMEDPIPE = 0x00200000 /* mode bit for named pipe (Unix, 9P2000.u) */
	DMSOCKET    = 0x00100000 /* mode bit for socket (Unix, 9P2000.u) */
	DMSETUID    = 0x00080000 /* mode bit for setuid (Unix, 9P2000.u) */
	DMSETGID    = 0x00040000 /* mode bit for setgid (Unix, 9P2000.u) */
)

func modeString(mode uint32) string {
	var buf bytes.Buffer

	switch {
	case mode&DMDIR > 0:
		buf.WriteByte('d')
	case mode&DMAPPEND > 0:
		buf.WriteByte('a')
	case mode&DMAUTH > 0:
		buf.WriteByte('A')
	case mode&DMDEVICE > 0:
		buf.WriteByte('D')
	case mode&DMSOCKET > 0:
		buf.WriteByte('S')
	case mode&DMNAMEDPIPE > 0:
		buf.WriteByte('P')
	default:
		buf.WriteByte('-')
	}

	switch {
	case mode&DMEXCL > 0:
		buf.WriteByte('l')
	case mode&DMSYMLINK > 0:
		buf.WriteByte('L')
	default:
		buf.WriteByte('-')
	}

	buf.WriteString(modes[(mode>>6)&7])
	buf.WriteString(modes[(mode>>3)&7])
	buf.WriteString(modes[(mode>>0)&7])

	return buf.String()
}
