// +----------------------------------------------------------------------
// | 日志
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月04日
// +----------------------------------------------------------------------

package gokit

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
	"time"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05} %{shortfile} ▶ %{level}%{color:reset} %{message}`,
)

var formatFile = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05} %{shortfile} ▶ %{level} %{message}`,
)

// dirname 目录名称
func Log(dirname ...string) *logging.Logger {
	var backend1Leveled logging.LeveledBackend
	var logger = logging.MustGetLogger("logging")

	// 写入文件、目录不存在，则创建
	if len(dirname) >= 1 {
		wd, _ := os.Getwd()
		path := fmt.Sprintf("%s/%s/%s", wd, "runtime", "log")
		_ = os.MkdirAll(path, 0777)
		file, _ := os.OpenFile(fmt.Sprintf("%s/%s_%s.log", path, dirname[0], time.Now().Format("2006-01-02")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		backend1 := logging.NewLogBackend(file, "", 0)
		backend1Formatter := logging.NewBackendFormatter(backend1, formatFile)
		backend1Leveled = logging.AddModuleLevel(backend1Formatter)
		backend1Leveled.SetLevel(logging.DEBUG, "")
	}

	// 控制台输出
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	if backend1Leveled == nil {
		logging.SetBackend(backend2Formatter)
	} else {
		logging.SetBackend(backend1Leveled, backend2Formatter)
	}
	return logger
}
