package leetcode

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path"
)

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

func getPath(f string) string {
	dir := os.TempDir()
	u, err := user.Current()
	if err == nil && u.HomeDir != "" {
		dir = path.Join(u.HomeDir, ".leetcode")
	}
	os.MkdirAll(dir, 0755)
	return path.Join(dir, f)
}

func saveCookies(cookies []*http.Cookie) {
	data, _ := json.Marshal(cookies)
	err := ioutil.WriteFile(getPath(cookiesFile), data, 0755)
	checkErr(err)
}

func getCookies() (cookies []*http.Cookie) {
	b, err := ioutil.ReadFile(getPath(cookiesFile))
	checkErr(err)
	err = json.Unmarshal(b, &cookies)
	checkErr(err)
	return
}
