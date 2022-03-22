package api

import (
	"net/url"
	"testing"
)

func TestInvalidLogin(t *testing.T) {
	ts := newTS("testdata/bad_login.http")
	defer ts.Close()

	c, err := New("User", "WrongPass", ts.URL, "default")
	if err != nil {
		t.Fatalf("can't crate API client %s", err)
	}

	err = c.Login()
	if err == nil {
		t.Error("client should return error on non 200 HTTP code but got nil")
	}
}

func TestLogin(t *testing.T) {
	ts := newTS("testdata/good_login.http")
	defer ts.Close()

	c, err := New("User", "Pass", ts.URL, "default")

	if err != nil {
		t.Fatalf("can't crate API client %s", err)
	}

	err = c.Login()
	if err != nil {
		t.Errorf("client should not return error on 200 HTTP code but got '%s'", err)
	}
}

func TestPersistCookie(t *testing.T) {
	ts := newTS("testdata/good_login.http")
	defer ts.Close()

	c, err := New("User", "Pass", ts.URL, "default")
	if err != nil {
		t.Fatalf("can't crate API client %s", err)
	}

	err = c.Login()
	if err != nil {
		t.Errorf("client should not return error on 200 HTTP code but got '%s'", err)
	}

	url, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("can't parse testserver URL: %s", err)
	}
	cookies := c.httpClient.Jar.Cookies(url)

	if len(cookies) == 0 {
		t.Error("Got 0 cookies from client")
	}
}
