// +----------------------------------------------------------------------
// | main.go
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月05日
// +----------------------------------------------------------------------

package main

import (
	"fmt"
	gokit "github.com/lengnuan-v/gokit/aliyun"
)

func main() {
	//var item = make(map[string]interface{})
	//item["host"] = "192.168.1.171"
	//item["username"] = "root"
	//item["password"] = "Eo$Tev$QFCSuwx*a"
	//item["port"] = 22

	aliyun := gokit.SDK{AccessKeyId: "LTAICeaRHntqwbWZ", AccessKeySecret: "zmyjb0ykHTqfYI5wTsR3lVHyO8HeBC"}
	d, e := aliyun.AliyunToken()
	fmt.Println(string(d), e)
	//var item = make(map[string]interface{})
	//item["name"] = 11
	//fmt.Println(db.Query("select * from ln_word where id=100"))

}
