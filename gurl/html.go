package gurl

import (
	"flag"
	"fmt"
	"io/ioutil"
)

type HTML struct {
	Flags   *flag.FlagSet
	Outpath string
}

func NewHTML() *HTML {
	h := HTML{}

	f := flag.NewFlagSet("html", flag.ExitOnError)
	f.StringVar(&h.Outpath, "out", "", "write the results to a file")
	usage(f)

	h.Flags = f
	return &h
}

func (c *HTML) Run(client Client, args []string) error {
	if err := c.Flags.Parse(args); err != nil {
		return err
	}
	args = c.Flags.Args()
	if len(args) == 0 {
		return fmt.Errorf("must pass in at least one URL")
	}
	for _, a := range args {
		b, err := client.Do(a, "text/html")
		if err != nil {
			return err
		}
		if c.Outpath != "" {
			ioutil.WriteFile(c.Outpath, b, 0644)
			continue
		}
		fmt.Println(string(b))
	}
	return nil
}
