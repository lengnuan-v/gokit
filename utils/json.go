// +----------------------------------------------------------------------
// | json
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月05日
// +----------------------------------------------------------------------

package gokit

import "encoding/json"

// 解码JSON字符串
func JSONDecode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}

// 对变量进行 JSON 编码
func JSONEncode(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}