package gurl

import (
	"encoding/json"
	"flag"
	"fmt"
)

type JSON struct {
	Flags  *flag.FlagSet
	Pretty bool
}

func (j JSON) Usage() {
	j.Flags.Usage()
}

func NewJSON() *JSON {
	j := JSON{}

	f := flag.NewFlagSet("json", flag.ExitOnError)
	f.BoolVar(&j.Pretty, "pretty", false, "pretty print JSON")
	usage(f)

	j.Flags = f
	return &j
}

func (c *JSON) Run(args []string) error {
	if err := c.Flags.Parse(args); err != nil {
		return err
	}
	args = c.Flags.Args()
	if len(args) == 0 {
		return fmt.Errorf("must pass in at least one URL")
	}
	for _, a := range args {
		b, err := do(a, "application/json")
		if err != nil {
			return err
		}
		if c.Pretty {
			m := []interface{}{}
			if err := json.Unmarshal(b, &m); err != nil {
				return err
			}
			b, err = json.MarshalIndent(m, "", "  ")
			if err != nil {
				return err
			}
		}
		fmt.Println(string(b))
	}
	return nil
}
