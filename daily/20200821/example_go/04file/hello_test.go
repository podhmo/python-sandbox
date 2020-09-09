package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func (c *Client) Hello() (*http.Response, error) {
	hc := c.HTTPClient
	if hc == nil {
		hc = http.DefaultClient
	}
	return hc.Get(strings.TrimSuffix(c.BaseURL, "/") + "/hello")
}

type HandlerRoundTripper struct {
	Before  func(*http.Request)
	Handler http.Handler
	After   func(*http.Response, *http.Request)
}

func (t *HandlerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if t.Before != nil {
		t.Before(r)
	}
	w := httptest.NewRecorder()
	t.Handler.ServeHTTP(w, r)
	res := w.Result()
	if t.After != nil {
		t.After(res, r)
	}
	return res, nil
}

func TestIt(t *testing.T) {
	c := &Client{
		HTTPClient: &http.Client{
			Timeout: 100 * time.Millisecond,
			Transport: &HandlerRoundTripper{
				Before: func(r *http.Request) {
					// /hello -> testdata/hello.json
					r.URL.Path = fmt.Sprintf("%s.json", strings.TrimSuffix(r.URL.Path, "/"))
				},
				Handler: http.FileServer(http.Dir("testdata")),
			},
		},
	}
	res, err := c.Hello()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	var b bytes.Buffer
	io.Copy(&b, res.Body)

	if res.StatusCode != 200 {
		t.Errorf("unexpected: %s", res.Status)
	}

	if strings.TrimSpace(b.String()) != strings.TrimSpace(`{"message": "hello"}`) {
		t.Errorf("unmatch: %s", b.String())
	}
}
