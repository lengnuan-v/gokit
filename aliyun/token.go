// +----------------------------------------------------------------------
// | aliyun token
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月05日
// +----------------------------------------------------------------------

package gokit

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SDK struct {
	AccessKeyId     string // AccessKey Id
	AccessKeySecret string // AccessKey Secret
}

// 阿里云Token
func (s *SDK) AliyunToken() ([]byte, error) {
	var err error
	var client *sdk.Client
	if client, err = sdk.NewClientWithAccessKey("cn-shanghai", s.AccessKeyId, s.AccessKeySecret); err != nil {
		return nil, err
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Domain = "nls-meta.cn-shanghai.aliyuncs.com"
	request.ApiName = "CreateToken"
	request.Version = "2019-02-28"
	var response *responses.CommonResponse
	if response, err = client.ProcessCommonRequest(request); err != nil {
		return nil, err
	}
	return []byte(response.GetHttpContentString()), nil
}
