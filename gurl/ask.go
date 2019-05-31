package gurl

import (
	"flag"
	"fmt"
	"io"
)

type Ask struct {
	io.Writer
	Flags *flag.FlagSet
}

func NewAsk(w io.Writer) *Ask {
	h := Ask{
		Writer: w,
	}

	f := flag.NewFlagSet("ask", flag.ExitOnError)
	usage(f)
	h.Flags = f
	return &h
}

func (c *Ask) Run(client Client, args []string) error {
	if err := c.Flags.Parse(args); err != nil {
		return err
	}
	args = c.Flags.Args()

	fmt.Println("Which website would you like to GURL?")
	var u string
	if _, err := fmt.Scanln(&u); err != nil {
		return err
	}

	if len(u) == 0 {
		return fmt.Errorf("must pass in at least one URL")
	}

	b, err := client.Do(u, "text/html")
	if err != nil {
		return err
	}
	fmt.Fprintln(c, string(b))
	return nil
}
