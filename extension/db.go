// +----------------------------------------------------------------------
// | db
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年04月02日
// +----------------------------------------------------------------------

package gohelp

import "github.com/gohouse/gorose"

// 数据库
// dsn "账号:密码@tcp(IP:端口)/数据库?charset=utf8"
// prefix 前缀
func GetDb(dsn string, prefix string) (*gorose.Connection, error) {
	var config = &gorose.DbConfigSingle{
		Driver:          "mysql",   // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,      // 是否开启sql日志
		SetMaxOpenConns: 0,         // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: 0,         // (连接池)闲置的连接数
		Prefix:          prefix,    // 表前缀
		Dsn:             dsn, // 数据库链接
	}
	return gorose.Open(config)
}
