package gurl

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func do(u string, ct string) ([]byte, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", ct)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

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
