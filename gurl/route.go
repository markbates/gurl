package gurl

import (
	"fmt"
	"io"
)

type runner interface {
	Run([]string) error
}

func Route(w io.Writer, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf(generalUsage)
	}

	var r runner
	a := args[0]
	switch a {
	case "json", "j":
		r = NewJSON(w)
	case "html", "h":
		r = NewHTML(w)
	case "file", "f":
		r = NewFile(w)
	case "ask", "a":
		r = NewHTML(w)
	case "-h":
		fmt.Fprintln(w, generalUsage)
		return nil
	default:
		return fmt.Errorf(generalUsage)
	}

	return r.Run(args[1:])
}

const generalUsage = `
Gurl is a tool for generating, and executing, HTTP Requests.

Usage:

	gurl <command> [arguments]
`
