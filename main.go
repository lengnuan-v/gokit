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
	sms2 "github.com/lengnuan-v/gokit/sms"
)

func main() {
	//var item = make(map[string]interface{})
	//item["host"] = "192.168.1.171"
	//item["username"] = "root"
	//item["password"] = "Eo$Tev$QFCSuwx*a"
	//item["port"] = 22

	//tts := aliyun.TTS{
	//	Appkey:     "8OVia795NNWNw9sa",
	//	Text:       "122222",
	//	Token:      "fcc98edb53754ff191a6555ca3233623",
	//}
	////aliyun := gokit.SDK{AccessKeyId: "LTAICeaRHntqwbWZ", AccessKeySecret: "zmyjb0ykHTqfYI5wTsR3lVHyO8HeBC"}
	////d, e := aliyun.AliyunToken()
	//dd, err := tts.Aliyuntts()
	//
	//utils.Tracefile(dd, "1.mp3", true)
	//fmt.Println(string(dd), err)
	//var item = make(map[string]interface{})
	//item["name"] = 11
	//fmt.Println(db.Query("select * from ln_word where id=100"))
	sms := sms2.SMS{
		Username: "JSM41733", // 短信帐号
		Password: "b21fk3rj", // 登录密码
		Verycode: "o2e8ckbl2uz7", // 校验码/密匙
		Tempid:   "JSM41733-0004", // 模版编号
	}
	dd, err := sms.Query("queryReport")
	fmt.Println(string(dd), err)
}
