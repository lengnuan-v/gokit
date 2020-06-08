// +----------------------------------------------------------------------
// | mysql
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2020年06月05日
// +----------------------------------------------------------------------

package gokit

import (
	"github.com/gohouse/gorose"
	gokit "github.com/lengnuan-v/gokit/utils"
)

type DB struct {
	Dsn    string // dsn
	Prefix string // 前缀
}

// 数据库
// dsn "账号:密码@tcp(IP:端口)/数据库?charset=utf8"
// prefix 前缀
func (d *DB) GetDb() (*gorose.Connection, error) {
	var config = &gorose.DbConfigSingle{
		Driver:          "mysql",  // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,     // 是否开启sql日志
		SetMaxOpenConns: 0,        // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: 0,        // (连接池)闲置的连接数
		Prefix:          d.Prefix, // 表前缀
		Dsn:             d.Dsn,    // 数据库链接
	}
	return gorose.Open(config)
}

// 新增数据
func (d *DB) Insert(tanleName string, data interface{}) error {
	if connection, err := d.GetDb(); err != nil {
		return err
	} else {
		defer connection.Close()
		db := connection.NewSession()
		_, err := db.Table(tanleName).Data(data).Insert()
		return err
	}
}

// 更新数据
func (d *DB) Update(tanleName string, set map[string]interface{}, cond string) error {
	if connection, err := d.GetDb(); err != nil {
		return err
	} else {
		defer connection.Close()
		db := connection.NewSession()
		_, err := db.Table(tanleName).Data(set).Where(cond).Update()
		return err
	}
}

// 删除数据
func (d *DB) Delete(tanleName string, cond string) error {
	if connection, err := d.GetDb(); err != nil {
		return err
	} else {
		defer connection.Close()
		_, err := connection.NewSession().Table(tanleName).Where(cond).Delete()
		return err
	}
}

// 获取条数
func (d *DB) Count(tanleName string, cond string) (int64, error) {
	if connection, err := d.GetDb(); err != nil {
		return 0, err
	} else {
		defer connection.Close()
		return connection.NewSession().Table(tanleName).Where(cond).Count()
	}
}

// 查询数据
func (d *DB) Select(tanleName string, fields []string, cond string, order string) ([]map[string]interface{}, error) {
	if connection, err := d.GetDb(); err != nil {
		return nil, err
	} else {
		defer connection.Close()
		return connection.NewSession().Table(tanleName).Fields(gokit.Implode(",", fields)).Where(cond).OrderBy(order).Get()
	}
}

// 分页查询数据
func (d *DB) SelectAll(tanleName string, fields []string, cond string, order string, limit, offset int) ([]map[string]interface{}, error) {
	if connection, err := d.GetDb(); err != nil {
		return nil, err
	} else {
		defer connection.Close()
		return connection.NewSession().Table(tanleName).Fields(gokit.Implode(",", fields)).Where(cond).OrderBy(order).Limit(limit).Offset(offset).Get()
	}
}
