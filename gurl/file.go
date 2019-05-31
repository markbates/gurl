package gurl

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type File struct {
	io.Writer
	Flags *flag.FlagSet
}

func NewFile(w io.Writer) *File {
	j := File{
		Writer: w,
	}

	f := flag.NewFlagSet("file", flag.ExitOnError)
	usage(f)
	j.Flags = f
	return &j
}

func (c *File) Run(args []string) error {
	if err := c.Flags.Parse(args); err != nil {
		return err
	}
	args = c.Flags.Args()
	if len(args) == 0 {
		return fmt.Errorf("must pass in at least one path")
	}
	for _, a := range args {
		if err := c.readFile(a); err != nil {
			return err
		}
	}
	return nil
}

func (c *File) readFile(a string) error {
	ff, err := os.Open(a)
	if err != nil {
		return err
	}
	defer ff.Close()
	_, err = io.Copy(c, ff)
	return err
}
