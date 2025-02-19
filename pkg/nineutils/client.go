package nineutils

import (
	"net"
	"path/filepath"
	"strings"

	"github.com/knusbaum/go9p/client"
)

func NewTCPClient(addr, user, aname string, opts ...client.Option) (*client.Client, error) {
	sock, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	cli, err := client.NewClient(sock, user, aname, opts...)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func NewUnixClient(path, user, aname string, opts ...client.Option) (*client.Client, error) {
	namespace := Namespace()
	name := strings.Split(path, "/")[0]
	path = filepath.Join(namespace, name)

	sock, err := net.Dial("unix", path)
	if err != nil {
		return nil, err
	}

	cli, err := client.NewClient(sock, user, aname, opts...)
	if err != nil {
		return nil, err
	}

	return cli, nil
}
