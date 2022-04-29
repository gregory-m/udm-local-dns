package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	URL      string // base URL of form http://ipaddr:port with no trailing slash
	Username string
	Password string
	Site     string

	httpClient *http.Client
}

func New(username, password, url, site string) (*Client, error) {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore bad SSL certificates UDM cert is self signed
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("can't create cookie jar: %w", err)
	}

	httpClient := &http.Client{Transport: transCfg, Jar: jar}

	return &Client{
		URL:        url,
		Username:   username,
		Password:   password,
		Site:       site,
		httpClient: httpClient,
	}, nil
}
