package models

import (
	"blog/configs"
	"github.com/IcecreamLee/goutils"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"xorm.io/xorm"
)

var ORM *xorm.Engine

func init() {
	var err error
	// ("mysql", "root:123456@tcp(localhost:3333)/test?charset=utf8&parseTime=True&loc=Local")
	//log.Println(configs.DBUser+":"+configs.DBPassword+"@tcp("+configs.DBHost+":"+configs.DBPort+")/"+configs.DBName+"?charset=utf8&parseTime=True&loc=Local")
	ORM, err = xorm.NewEngine("mysql", configs.DBUser+":"+configs.DBPassword+"@tcp("+configs.DBHost+":"+configs.DBPort+")/"+configs.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}
	err = ORM.Ping()
	if err != nil {
		panic("Failed to ping database: " + err.Error())
	}
}

// paginate 分页查询
func paginate(entities interface{}, sqlStr string, args []interface{}, curPage int, limit int, mode int) interface{} {
	if limit > 0 {
		queryCount := 0
		if mode == 1 {
			queryCount = count(sqlStr, args...)
		}
		startNum := goutils.IntToString(int64((curPage - 1) * limit))
		sqlStr += " LIMIT " + startNum + "," + goutils.IntToString(int64(limit))
		_ = ORM.SQL(sqlStr, args...).Find(entities)
		return queryCount
	}
	return 0
}

// count 获取SQL查询结果的总条数
func count(sqlStr string, args ...interface{}) int {
	// 把select xxx from table 替换成 select count(*) c from table
	pos := strings.Index(strings.ToLower(sqlStr), " from ")
	if pos >= 0 {
		sqlStr = "SELECT count(*) c " + sqlStr[pos+1:]
	}

	sqlOrArgs := append([]interface{}{sqlStr}, args...)
	results, err := ORM.Query(sqlOrArgs...)
	if err != nil {
		return 0
	}
	return goutils.ToInt(string(results[0]["c"]))
}
