package myutils

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &aliyunaccess.AccessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &aliyunaccess.AccessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("ecs.cn-hongkong.aliyuncs.com")
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}
