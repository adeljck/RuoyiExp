package utils

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

func CheckCookie(url string, cookie []*http.Cookie, Timeout int) bool {
	client := resty.New()
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	client.SetCookies(cookie)
	client.SetTimeout(time.Duration(Timeout) * time.Second)
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBaseURL(url)
	res, err := client.R().Get("/index")
	if err != nil {
		if res.StatusCode() == http.StatusFound {
			return false
		}
	}
	return true
}
