package gurl

import (
	"fmt"
)

type runner interface {
	Run(Client, []string) error
}

func Route(client Client, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf(generalUsage)
	}

	var r runner
	a := args[0]
	switch a {
	case "json", "j":
		r = NewJSON()
	case "html", "h":
		r = NewHTML()
	case "file", "f":
		r = NewFile()
	case "ask", "a":
		r = NewAsk()
	case "-h":
		fmt.Fprintln(client.Out, generalUsage)
		return nil
	default:
		return fmt.Errorf(generalUsage)
	}

	return r.Run(client, args[1:])
}

const generalUsage = `
Gurl is a tool for generating, and executing, HTTP Requests.

Usage:

	gurl <command> [arguments]
`
