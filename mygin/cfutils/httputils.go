package cfutils

import "github.com/go-resty/resty/v2"

func HttpCline() *resty.Client {
	return resty.New()
}
