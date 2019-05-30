package main

import (
	"log"
	"os"

	"github.com/markbates/gurl/gurl"
)

type runner interface {
	Run([]string) error
}

type usage interface {
	Usage()
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("no enough arguments")
	}

	var r runner
	switch args[0] {
	case "json":
		r = gurl.NewJSON()
	case "html":
		r = gurl.NewHTML()
	case "file":
		r = gurl.NewFile()
	case "ask":
		r = gurl.NewAsk()
	default:
		log.Fatalf("unknown sub-command %s", args[0])
	}
	if err := r.Run(args[1:]); err != nil {
		if u, ok := r.(usage); ok {
			u.Usage()
		}
		log.Fatal(err)
	}
}
