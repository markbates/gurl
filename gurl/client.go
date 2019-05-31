package gurl

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func (c Client) Do(u string, ct string) ([]byte, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", ct)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
