package net

import (
	"fmt"
	"net/http"
	"net/url"
)

type DefaultProvider struct{}

func (DefaultProvider) Post(url string, data url.Values) error {
	response, err := http.PostForm(url, data)
	if err != nil {
		return err
	}

	fmt.Println(response)
	return nil
}
