package leetcode

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/user"
	"path"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func client() *http.Client {
	jar, err := cookiejar.New(nil)
	// jar.SetCookies()
	checkErr(err)
	return &http.Client{
		Jar: jar,
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

func getPath(f string) string {
	dir := os.TempDir()
	u, err := user.Current()
	if err == nil && u.HomeDir != "" {
		dir = path.Join(u.HomeDir, ".leetcode")
	}
	err = os.MkdirAll(dir, 0755)
	checkErr(err)
	return path.Join(dir, f)
}

func saveCookies(cookies []*http.Cookie) {
	data, err := json.Marshal(cookies)
	checkErr(err)
	dst := bytes.Buffer{}
	err = json.Indent(&dst, data, "", "\t")
	checkErr(err)
	err = ioutil.WriteFile(getPath(cookiesFile), dst.Bytes(), 0755)
	checkErr(err)
}

func getCookies() (cookies []*http.Cookie) {
	b, err := ioutil.ReadFile(getPath(cookiesFile))
	checkErr(err)
	err = json.Unmarshal(b, &cookies)
	checkErr(err)
	return
}

func saveCredential(data url.Values) {
	u := url.UserPassword(data.Get("login"), data.Get("password"))
	err := ioutil.WriteFile(getPath(credentialsFile), []byte(u.String()), 0755)
	checkErr(err)
}
