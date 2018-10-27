package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("Flavour"); err != nil {
		http.SetCookie(w, &http.Cookie{Name: "Flavour", Value: "Chocolate Chip"})
	} else {
		cookie.Value = "Oatmeal Raisin"
		http.SetCookie(w, cookie)
	}
}

func main() {
	ts := httptest.NewServer(http.HandlerFunc(setCookie))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		panic(err)
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Jar: jar,
	}
	fmt.Println("After 1st Request")
	if _, err := client.Get(u.String()); err != nil {
		panic(err)
	}

	for _, cookie := range client.Jar.Cookies(u) {
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
	fmt.Println("After 2nd Request")
	if _, err := client.Get(u.String()); err != nil {
		panic(err)
	}

	for _, cookie := range client.Jar.Cookies(u) {
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
}
