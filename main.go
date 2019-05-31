package main

import (
	"fmt"
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
	if err := gurl.Route(os.Stdout, os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
