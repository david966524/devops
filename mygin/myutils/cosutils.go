package myutils

import (
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// 腾讯云 cos clent
func CosClient() *cos.Client {
	u, _ := url.Parse("https://newim-txy4-123qwer-1312026995.cos.ap-hongkong.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  tencentyun.AccessKeyId,
			SecretKey: tencentyun.AccessKeySecret,
		},
	})
	return client
}
