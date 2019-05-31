package gurl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Ask(t *testing.T) {
	table := []struct {
		in     []string
		site   string
		result []byte
		err    bool
	}{
		{[]string{}, "https://example.com", []byte("hello!"), false},
		{[]string{}, "2234dfv9dsfvjasdf", []byte("hello!"), true},
	}

	for _, tt := range table {
		t.Run(fmt.Sprint(tt), func(st *testing.T) {

			ask := NewAsk()

			in := strings.NewReader(tt.site + "\n")

			client := NewClient(BufIO(in), func(req *http.Request) (*http.Response, error) {
				if tt.err {
					return nil, fmt.Errorf("oops")
				}

				res := httptest.NewRecorder()

				res.Write(tt.result)

				return res.Result(), nil
			})

			err := ask.Run(client, tt.in)

			if Err(t, err, tt.err) == tt.err {
				return
			}

			b, err := ioutil.ReadAll(client.IO.Out)
			Err(t, err, false)

			Equal(t, tt.result, bytes.TrimSpace(b))

		})
	}
}
