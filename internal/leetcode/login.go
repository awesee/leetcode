package leetcode

import (
	"net/http"
	"net/url"
	"strings"
)

func AccountsLogin(user string, pwd string) (*http.Response, error) {
	resp, err := http.Head(AccountsLoginUrl)
	checkErr(err)
	cookies := resp.Cookies()
	saveCookies(cookies)
	csrftoken := getCsrfToken(cookies)
	data := url.Values{
		"login":               {user},
		"password":            {pwd},
		"csrfmiddlewaretoken": {csrftoken},
	}
	req, err := http.NewRequest("POST", AccountsLoginUrl, strings.NewReader(data.Encode()))
	checkErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", AccountsLoginUrl)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err = http.DefaultClient.Do(req)
	checkErr(err)
	saveCookies(resp.Cookies())
	return resp, err
}
