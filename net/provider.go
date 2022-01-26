package net

import (
	"net/url"
)

type Provider interface {
	Post(url string, data url.Values) error
}
