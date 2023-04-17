package myutils

import "github.com/go-resty/resty/v2"

func HttpClinet() *resty.Client {
	return resty.New()
}
