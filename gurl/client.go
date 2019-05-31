package gurl

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HTTP *http.Client
	In   io.Reader
	Out  io.Writer
}

func (c Client) Do(u string, ct string) ([]byte, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", ct)

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
