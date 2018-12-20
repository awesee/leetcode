package leetcode

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/user"
	"path"
)

var err error

func init() {
	http.DefaultClient.Jar, err = cookiejar.New(nil)
	checkErr(err)
	http.DefaultClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		req.Header.Set("Referer", req.URL.String())
		fmt.Println(req.URL.String())
		if len(via) >= 3 {
			return errors.New("stopped after 3 redirects")
		}
		return nil
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCsrfToken(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == "csrftoken" {
			return cookie.Value
		}
	}
	return ""
}

func getCachePath(f string) string {
	dir, err := os.UserCacheDir()
	checkErr(err)
	u, err := user.Current()
	if err == nil && u.HomeDir != "" {
		dir = path.Join(u.HomeDir, ".leetcode")
	}
	filename := getFilePath(path.Join(dir, f))
	return filename
}

func getFilePath(f string) string {
	dir := path.Dir(f)
	if dir != "" {
		err := os.MkdirAll(dir, 0755)
		checkErr(err)
	}
	return f
}

func saveCookies(cookies []*http.Cookie) {
	data, err := json.Marshal(cookies)
	checkErr(err)
	dst := bytes.Buffer{}
	err = json.Indent(&dst, data, "", "\t")
	checkErr(err)
	err = ioutil.WriteFile(getCachePath(cookiesFile), dst.Bytes(), 0644)
	checkErr(err)
}

func getCookies() (cookies []*http.Cookie) {
	b, err := ioutil.ReadFile(getCachePath(cookiesFile))
	checkErr(err)
	err = json.Unmarshal(b, &cookies)
	checkErr(err)
	return
}

func saveCredential(v url.Values) {
	data, err := json.Marshal(v)
	checkErr(err)
	dst := bytes.Buffer{}
	err = json.Indent(&dst, data, "", "\t")
	checkErr(err)
	err = ioutil.WriteFile(getCachePath(credentialsFile), dst.Bytes(), 0644)
	checkErr(err)
}

func getCredential() (data url.Values) {
	b, err := ioutil.ReadFile(getCachePath(credentialsFile))
	checkErr(err)
	err = json.Unmarshal(b, &data)
	checkErr(err)
	return
}
