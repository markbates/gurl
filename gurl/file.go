package gurl

import (
	"flag"
	"fmt"
	"io/ioutil"
)

type File struct {
	Flags *flag.FlagSet
}

func (f File) Usage() {
	f.Flags.Usage()
}

func NewFile() *File {
	j := File{}

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
		b, err := ioutil.ReadFile(a)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}
	return nil
}
