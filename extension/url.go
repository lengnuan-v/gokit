// +----------------------------------------------------------------------
// | url方法
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年04月02日
// +----------------------------------------------------------------------

package gohelp

import "net/url"

// 解析 URL，返回其组成部分
// str 要解析的 URL
// component -1
// ParseURL("http://username:password@hostname:9090/path?arg=value#anchor", -1)
func ParseURL(str string, component int) (map[string]string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return nil, err
	}
	if component == -1 {
		component = 1 | 2 | 4 | 8 | 16 | 32 | 64 | 128
	}
	var components = make(map[string]string)
	if (component & 1) == 1 {
		components["scheme"] = u.Scheme
	}
	if (component & 2) == 2 {
		components["host"] = u.Hostname()
	}
	if (component & 4) == 4 {
		components["port"] = u.Port()
	}
	if (component & 8) == 8 {
		components["user"] = u.User.Username()
	}
	if (component & 16) == 16 {
		components["pass"], _ = u.User.Password()
	}
	if (component & 32) == 32 {
		components["path"] = u.Path
	}
	if (component & 64) == 64 {
		components["query"] = u.RawQuery
	}
	if (component & 128) == 128 {
		components["fragment"] = u.Fragment
	}
	return components, nil
}

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

// 生成 url-encoded 之后的请求字符串
// HTTPBuildQuery(map[string][]string{"first": {"value"}, "multi": {"foo bar", "baz"}})
// first=value&multi=foo+bar&multi=baz
func HTTPBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}
