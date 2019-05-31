package gurl

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Equal(t *testing.T, a, b interface{}) {
	t.Helper()
	res := cmp.Equal(a, b, cmpopts.IgnoreUnexported())
	if !res {
		t.Fatalf("expected %q got %q", a, b)
	}
}

func Err(t *testing.T, err error, expected bool) bool {
	t.Helper()
	if expected {
		if err == nil {
			t.Fatalf("error expected, but none was found")
		}
		return true
	}

	if err != nil {
		t.Fatalf("no error expected, but one found %v", err)
	}
	return true
}

func BufIO(r io.Reader) IO {
	return IO{
		In:  r,
		Out: &bytes.Buffer{},
		Err: &bytes.Buffer{},
	}
}
