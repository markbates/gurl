package gurl

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type File struct {
	Flags *flag.FlagSet
}

func NewFile() *File {
	j := File{}

	f := flag.NewFlagSet("file", flag.ExitOnError)
	usage(f)
	j.Flags = f
	return &j
}

func (c *File) Run(client Client, args []string) error {
	if err := c.Flags.Parse(args); err != nil {
		return err
	}
	args = c.Flags.Args()
	if len(args) == 0 {
		return fmt.Errorf("must pass in at least one path")
	}
	for _, a := range args {
		if err := c.readFile(client, a); err != nil {
			return err
		}
	}
	return nil
}

func (c *File) readFile(client Client, a string) error {
	ff, err := os.Open(a)
	if err != nil {
		return err
	}
	defer ff.Close()
	_, err = io.Copy(client.Out, ff)
	return err
}
