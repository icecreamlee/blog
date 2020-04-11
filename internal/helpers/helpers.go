package helpers

import (
	"github.com/IcecreamLee/goutils"
	"html/template"
	"net/http"
	"regexp"
	"strings"
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

// GetSite 返回指定请求r的域名: [协议类型]://[服务器地址]:[端口号]/, 如https://google.com/
func GetSite(r *http.Request) string {
	if r.TLS == nil {
		return "http://" + r.Host + "/"
	}
	return "https://" + r.Host + "/"
}

// GetFullURL 返回指定请求r的完整URL地址: [协议类型]://[服务器地址]:[端口号]/[路径]?[查询参数]，如：https://www.a.com/b/c?d=e
func GetFullURL(r *http.Request) string {
	return GetSite(r) + r.RequestURI[1:]
}

// GetURL 返回指定请求r不包含查询参数的URL地址: [协议类型]://[服务器地址]:[端口号]/[路径]，如：https://www.a.com/b/c
func GetURL(r *http.Request) string {
	pos := strings.Index(r.RequestURI, "?")
	if pos >= 0 {
		return GetSite(r) + r.RequestURI[1:pos]
	}
	return GetSite(r) + r.RequestURI[1:]
}

// SetURL 根据指定请求r的当前域名和指定的路径path去设置URL
// 如: 当前r的URL为https://a.com/b/c,
// SetURL(r, "/index") => https://a.com/index
// SetURL(r, "d?e=f") => https://a.com/d?e=f
func SetURL(r *http.Request, path string) string {
	if path == "" {
		return GetURL(r)
	}
	if path[0:1] == "/" {
		return GetSite(r) + path[1:]
	}
	url := GetURL(r)
	pos := strings.LastIndex(url, "/")
	return url[0:pos+1] + path
}

// IsAjax 返回指定的请求r是否为ajax异步请求
func IsAjax(r *http.Request) bool {
	values, ok := r.Header["X-Requested-With"]
	return ok && values[0] == "XMLHttpRequest"
}
