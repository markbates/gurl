package gurl

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	IO
	HTTP *http.Client
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

func NewClient(io IO, rt RoundTripper) Client {
	return Client{
		IO: io,
		HTTP: &http.Client{
			Transport: rt,
		},
	}
}

type RoundTripper func(req *http.Request) (*http.Response, error)

func (r RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}
