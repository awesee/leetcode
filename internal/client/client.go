// Package client provides support for http request.
package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/openset/leetcode/internal/base"
)

var err error

// Get - client.Get
func Get(url string) []byte {
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		base.CheckErr(err)
		return body
	}
	return nil
}

// PostJSON - client.PostJSON
func PostJSON(url, jsonStr string) []byte {
	if resp, err := http.Post(url, "application/json", strings.NewReader(jsonStr)); err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		base.CheckErr(err)
		return body
	}
	return nil
}

func init() {
	http.DefaultClient.Jar, err = cookiejar.New(nil)
	base.CheckErr(err)
	http.DefaultClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		req.Header.Set("Referer", req.URL.String())
		if len(via) >= 3 {
			return errors.New("stopped after 3 redirects")
		}
		return nil
	}
}
