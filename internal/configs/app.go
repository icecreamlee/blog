package configs

import (
	"github.com/IcecreamLee/goutils"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	ENV            = "dev"
	Port           = "8000"
	ManagePassword = ""
	Categories     []string
)

func init() {
	RootPath := goutils.GetCurrentPath()
	iniConfig, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, RootPath+"app.ini")
	if err != nil {
		panic("Load config file failure: " + err.Error())
	}
	section := iniConfig.Section("app")
	ENV = section.Key("env").String()
	Port = section.Key("port").String()
	ManagePassword = section.Key("manage_password").String()
	Categories = strings.Split(section.Key("categories").String(), ",")
	section = iniConfig.Section("db")
	DBHost = section.Key("host").String()
	DBPort = section.Key("port").String()
	DBUser = section.Key("user").String()
	DBPassword = section.Key("password").String()
	DBName = section.Key("db_name").String()
}
