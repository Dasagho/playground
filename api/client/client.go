package client

import (
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	*http.Client
}

func New() *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic("Error building cookiejar" + err.Error())
	}

	return &Client{
		Client: &http.Client{
			Jar: jar,
		},
	}
}
