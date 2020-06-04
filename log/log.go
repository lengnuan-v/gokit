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
)

var format = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05} %{shortfile} ▶ %{level:.4s} %{message}%{color:reset}`,
)

var formatFile = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05} %{shortfile} ▶ %{level:.4s} %{message}`,
)

func Log(filename ...string) *logging.Logger {
	var backend1Leveled logging.LeveledBackend
	var logger = logging.MustGetLogger("example")

	// 写入文件、目录不存在，则创建
	if len(filename) >= 1 {
		path := fmt.Sprintf("../%s", "runtime")
		_ = os.MkdirAll(path, 0777)
		file, _ := os.OpenFile(fmt.Sprintf("%s/%s.log", path, filename[0]), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
