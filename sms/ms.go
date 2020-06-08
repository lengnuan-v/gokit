// +----------------------------------------------------------------------
// | 美圣 短信 http://www.jsmsxx.com/
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月08日
// +----------------------------------------------------------------------

package gokit

import (
	"errors"
	"fmt"
	gokit "github.com/lengnuan-v/gokit/utils"
	"net/url"
	"strconv"
)

const (
	GetAmount   = "getAmount"   // 获得余额
	QueryReport = "queryReport" // 获得状态报告,只能查询当天的，已获取的状态报告后续不会再获取
	QueryMo     = "queryMo"     // 获得上行短信,只能查询当天的，已获取的上行短信后续不会再获取
)

var (
	smsStrlen = 300                                                              // 字数限制
	smsApi    = "http://112.74.76.186:8030/service/httpService/httpInterface.do" // 服务器参数设置
)

type SMS struct {
	Username string // 短信帐号
	Password string // 登录密码
	Verycode string // 校验码/密匙
	Method   string // 如果有乱码请尝试用 sendUtf8Msg 或  sendGbkMsg
	Msgtype  string // 1-普通短信，2-模板短信
	Tempid   string // 模板编号
	Code     string // 编码
}

// mobile 多个手机以,分割，不得存在空格
// content fmt.Sprintf("%s%s", "@1@=", "服务器出现异常")
func (s *SMS) Send(mobile, content string) ([]byte, error) {
	mb_strlen := gokit.MbStrlen(content)
	if mb_strlen > smsStrlen || mb_strlen <= 0 {
		return nil, errors.New("短信内容（最多" + strconv.Itoa(smsStrlen) + "个汉字，最少1个汉字）")
	}
	// 默认method
	if gokit.IsEmpty(s.Method) == true {
		s.Method = "sendMsg"
	}
	// 默认短信
	if gokit.IsEmpty(s.Msgtype) == true {
		s.Msgtype = "2"
	}
	// 编码
	if gokit.IsEmpty(s.Code) == true {
		s.Code = "utf-8"
	}
	var info = url.Values{
		"username": {s.Username}, // 短信帐号
		"password": {s.Password}, // 登录密码
		"veryCode": {s.Verycode}, // 校验码/密匙
		"method":   {s.Method},   // 如果有乱码请尝试用 sendUtf8Msg 或  sendGbkMsg
		"mobile":   {mobile},     // 手机号码
		"content":  {content},    // 内容
		"msgtype":  {s.Msgtype},  // 1-普通短信，2-模板短信
		"tempid":   {s.Tempid},   // 模板编号 ， 在客户创建模板后会生成模板编号
		"code":     {s.Code},     // 编码
	}
	results, err := gokit.HttpRequest(nil, "GET", fmt.Sprintf("%s?%s", smsApi, info.Encode()), nil, nil)
	return results, err
}

// 获得余额、获得状态报告、获得上行短信
func (s *SMS) Query(method string) ([]byte, error) {
	var info = url.Values{
		"username": {s.Username}, // 短信帐号
		"password": {s.Password}, // 登录密码
		"veryCode": {s.Verycode}, // 校验码/密匙
		"method":   {method},     // 查询参数
	}
	results, err := gokit.HttpRequest(nil, "GET", fmt.Sprintf("%s?%s", smsApi, info.Encode()), nil, nil)
	return results, err
}
