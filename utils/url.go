// +----------------------------------------------------------------------
// | url
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月05日
// +----------------------------------------------------------------------

package gokit

import "net/url"

// 编码 URL 字符串
// str 要编码的字符串
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

// 解码URL编码的字符串
// str 要解码的字符串
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}
