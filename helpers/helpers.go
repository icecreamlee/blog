package helpers

import (
	"github.com/IcecreamLee/goutils"
	"html/template"
	"regexp"
	"time"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"html":           HTML,
		"substr":         SubStr,
		"stripTags":      StripTags,
		"dateFormat":     DateFormat,
		"datetimeFormat": DatetimeFormat,
		"int2date":       Int2date,
		"int2datetime":   Int2datetime,
	}
}

// html 将页面变量内容作为HTML显示
func HTML(str string) interface{} {
	return template.HTML(str)
}

// substr 将字符串str截取一部分并返回
func SubStr(str string, start int, length int) string {
	return str[start : start+length]
}

// StripTags 将字符串str中的html标签去除并且返回
func StripTags(str string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(str, "")
}

// datetimeFormat 将时间t格式化为字符串类型返回
func DatetimeFormat(t time.Time, layout ...string) string {
	return goutils.DatetimeFormat(t, layout...)
}

// datetimeFormat 将时间t格式化为字符串类型返回
func DateFormat(t time.Time, layout ...string) string {
	return goutils.DateFormat(t, layout...)
}

// Int2date 返回一个格式化的日期字符串,
func Int2date(intDate int, layout ...string) string {
	return goutils.Int2date(intDate, layout...)
}

// Int2datetime 返回一个格式化的日期字符串,
func Int2datetime(intDate int, layout ...string) string {
	return goutils.Int2datetime(intDate, layout...)
}
