package leetcode

import (
	"net/http"
	"net/url"
	"strings"
)

func AccountsLogin(user string, pwd string) (*http.Response, error) {
	resp, _ := http.Head(AccountsLoginUrl)
	cookies := resp.Cookies()
	saveCookies(cookies)
	csrftoken := getCsrfToken(cookies)
	data := url.Values{
		"login":               {user},
		"password":            {pwd},
		"csrfmiddlewaretoken": {csrftoken},
	}
	req, err := http.NewRequest("POST", AccountsLoginUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
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
