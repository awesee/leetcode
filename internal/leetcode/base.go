package leetcode

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/client"
)

// leetcode var
var (
	authInfo        = base.AuthInfo
	checkErr        = base.CheckErr
	filePutContents = base.FilePutContents
	jsonIndent      = base.JSONIndent
	LockStr         = " 🔒"
	translationSet  = make(map[int]string)
	titleSlugID     = make(map[string]int)
)

// Clean - leetcode.Clean
func Clean() {
	dir := getCachePath("")
	err := os.RemoveAll(dir)
	checkErr(err)
}

func graphQLRequest(graphQL, jsonStr, filename string, days int, v interface{}) {
	data := remember(filename, days, func() []byte {
		return client.PostJSON(graphQL, jsonStr)
	})
	jsonDecode(data, &v)
}

func remember(filename string, days int, f func() []byte) []byte {
	filename = getCachePath(filename)
	if fi, err := os.Stat(filename); err == nil {
		t := fi.ModTime().AddDate(0, 0, days)
		if time.Now().Add(time.Duration(t.Unix()%86400) * time.Second).Before(t) {
			return fileGetContents(filename)
		} else if cts := f(); len(cts) == 0 {
			return fileGetContents(filename)
		} else {
			return filePutContents(filename, cts)
		}
	}
	return filePutContents(filename, f())
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
	dir, err := os.UserHomeDir()
	if err != nil {
		dir, err = os.UserCacheDir()
		checkErr(err)
	}
	return filepath.Join(dir, ".leetcode", f)
}

func fileGetContents(filename string) []byte {
	if cts, err := ioutil.ReadFile(filename); err == nil {
		return cts
	}
	return nil
}

func jsonEncode(v interface{}) []byte {
	data, err := json.Marshal(v)
	checkErr(err)
	return jsonIndent(data)
}

func jsonDecode(data []byte, v interface{}) {
	if len(data) > 0 {
		err := json.Unmarshal(data, v)
		checkErr(err)
	}
}

func saveCookies(cookies []*http.Cookie) {
	filePutContents(getCachePath(cookiesFile), jsonEncode(cookies))
}

func getCookies() (cookies []*http.Cookie) {
	cts := fileGetContents(getCachePath(cookiesFile))
	jsonDecode(cts, &cookies)
	return
}

func saveCredential(data url.Values) {
	filePutContents(getCachePath(credentialsFile), jsonEncode(data))
}

func getCredential() (data url.Values) {
	cts := fileGetContents(getCachePath(credentialsFile))
	jsonDecode(cts, &data)
	return
}

func slugToSnake(slug string) string {
	return strings.ReplaceAll(slug, "-", "_")
}
