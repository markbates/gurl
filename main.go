package main

import (
	"fmt"
	"net/http"
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
	rt := gurl.Router{
		Writer: os.Stdout,
		Client: gurl.Client{Client: http.DefaultClient},
	}

	if err := rt.Route(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
