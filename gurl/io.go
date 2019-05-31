package gurl

import (
	"io"
	"os"
)

type IO struct {
	In  io.Reader
	Out io.ReadWriter
	Err io.Writer
}

func STD() IO {
	return IO{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}
}
