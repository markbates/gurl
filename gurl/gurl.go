package gurl

import (
	"flag"
	"fmt"
)

func usage(f *flag.FlagSet) {
	f.Usage = func() {
		fmt.Printf("Usage:\n\n")
		fmt.Printf("gurl %s [flags] [args...]\n", f.Name())
		f.VisitAll(func(fl *flag.Flag) {
			fmt.Printf("\t-%s\t%q (%q)\n", fl.Name, fl.Usage, fl.DefValue)
		})
		fmt.Print("\n")
	}
}
