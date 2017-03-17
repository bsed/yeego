/**
 * Created by angelina-zf on 17/2/25.
 */

// yeego
// 数据库处理,使用x-orm
// 依赖："github.com/go-xorm/xorm"
// TODO 目前只使用mysql,后续可以新增其他数据库类型,貌似直接引入那些包就可以了^_^
package yeego

import (
	"fmt"
	"os"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

// DBConfig
// mysql数据库配置
type DBConfig struct {
	DbUser       string // 用户
	DbPassword   string // 密码
	DbHost       string // 地址
	DbPort       string // 端口
	DbName       string // 数据库名称
	DbType       string // 数据库类型
	MaxIdleConns int    // 最大空闲连接数
	MaxOpenConns int    // 最大打开连接数
}

var xORM map[string]*xorm.Engine

func init() {
	xORM = make(map[string]*xorm.Engine)
}

// MustInitDefaultDb
// 初始化默认的数据库配置，读取配置文件.
func MustInitDefaultDb(dbConfig DBConfig) {
	var orm *xorm.Engine
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConfig.DbUser, dbConfig.DbPassword,
		dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName)
	orm, err = xorm.NewEngine(dbConfig.DbType, dsn)
	if err != nil {
		panic("数据库连接异常！请检查数据库配置文件")
	}
	err = orm.Ping()
	if err != nil {
		panic("数据库连接异常！请检查数据库配置文件")
	}
	//如果是dev模式，需要打印sql语句
	orm.ShowSQL(true)
	orm.ShowExecTime(true)
	//设置sql语句存储文件地址
	f, err := os.Create("default-sql.log")
	if err != nil {
		fmt.Println(err.Error())
	}
	orm.SetLogger(xorm.NewSimpleLogger(f))
	orm.SetMapper(core.SameMapper{})
	orm.SetColumnMapper(core.SameMapper{})
	orm.SetTableMapper(core.GonicMapper{})
	if dbConfig.MaxIdleConns != 0 {
		orm.SetMaxIdleConns(dbConfig.MaxIdleConns)
	}
	if dbConfig.MaxOpenConns != 0 {
		orm.SetMaxOpenConns(dbConfig.MaxOpenConns)
	}
	xORM["default"] = orm
}

// MustInitDb
// 初始化新的数据库实例.
func MustInitDb(instanceName string, dbConfig DBConfig) {
	var orm *xorm.Engine
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConfig.DbUser, dbConfig.DbPassword,
		dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName)
	orm, err = xorm.NewEngine(dbConfig.DbType, dsn)
	if err != nil {
		panic("数据库连接异常！请检查数据库配置文件")
	}
	err = orm.Ping()
	if err != nil {
		panic("数据库连接异常！请检查数据库配置文件")
	}
	//如果是dev模式，需要打印sql语句
	orm.ShowSQL(true)
	//设置sql语句存储文件地址
	f, err := os.Create(instanceName + "-sql.log")
	if err != nil {
		fmt.Println(err.Error())
	}
	orm.SetLogger(xorm.NewSimpleLogger(f))
	orm.SetMapper(core.SameMapper{})
	orm.SetColumnMapper(core.SameMapper{})
	orm.SetTableMapper(core.GonicMapper{})
	if dbConfig.MaxIdleConns != 0 {
		orm.SetMaxIdleConns(dbConfig.MaxIdleConns)
	}
	if dbConfig.MaxOpenConns != 0 {
		orm.SetMaxOpenConns(dbConfig.MaxOpenConns)
	}
	xORM[instanceName] = orm
}

// GetORMByName 通过实例名称获取数据库连接.
func GetORMByName(instanceName string) *xorm.Engine {
	orm, ok := xORM[instanceName]
	if !ok {
		panic("数据库实例" + instanceName + "不存在！")
	}
	return orm
}

// GetORM
// 获取默认的数据库连接.
func GetORM() *xorm.Engine {
	return xORM["default"]
}
