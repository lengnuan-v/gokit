// +----------------------------------------------------------------------
// | log_test.go
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月04日
// +----------------------------------------------------------------------

package main

import gokit "github_com/gokit/log"

func main()  {
	log := gokit.Log("main")
	log.Info("noti11得分ce")
	log.Notice("noti11得分ce")
	log.Warning("warning")
	log.Error("xiaorui.cc")
	log.Critical("太严重了")
	//log.Info("info")
	//log.Notice("notice")
	//log.Warning("warning")
	//log.Error("xiaorui.cc")
	//log.Critical("太严重了1")
}
