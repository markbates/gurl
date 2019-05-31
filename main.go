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
	client := gurl.Client{
		HTTP: http.DefaultClient,
		In:   os.Stdin,
		Out:  os.Stdout,
	}

	if err := gurl.Route(client, os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
