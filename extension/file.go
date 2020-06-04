// +----------------------------------------------------------------------
// | 目录文件函数
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月04日
// +----------------------------------------------------------------------

package gokit

import (
	"os"
	"path"
	"strings"
)

// 创建并写入文件
// content 写入的内容
// filename 规定要写入的文件名称。
// keep true 已有的数据会被保留 false 已有的数据会被清除
func Tracefile(content, filename string, keep bool) (int, error) {
	dir, _ := path.Split(filename)
	// 目录不存在，则创建
	_ = os.MkdirAll(dir, 0777)
	var file *os.File
	if keep == true {
		file, _ = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	} else {
		file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	}
	defer file.Close()
	fileContent := strings.Join([]string{content, "\n"}, "")
	buf := []byte(fileContent)
	return file.Write(buf)
}
