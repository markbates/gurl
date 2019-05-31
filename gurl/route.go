package gurl

import (
	"fmt"
	"io"
)

type runner interface {
	Run(Client, []string) error
}

type Router struct {
	Writer io.Writer
	Client Client
}

func (rt Router) Route(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf(generalUsage)
	}

	var r runner
	a := args[0]
	switch a {
	case "json", "j":
		r = NewJSON(rt.Writer)
	case "html", "h":
		r = NewHTML(rt.Writer)
	case "file", "f":
		r = NewFile(rt.Writer)
	case "ask", "a":
		r = NewAsk(rt.Writer)
	case "-h":
		fmt.Fprintln(rt.Writer, generalUsage)
		return nil
	default:
		return fmt.Errorf(generalUsage)
	}

	return r.Run(rt.Client, args[1:])
}

const generalUsage = `
Gurl is a tool for generating, and executing, HTTP Requests.

Usage:

	gurl <command> [arguments]
`
