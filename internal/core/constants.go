package core

import (
	"github.com/IcecreamLee/goutils"
	"path/filepath"
)

var (
	BasePath string
	RootPath string
)

func InitPaths() {
	var err error
	BasePath = goutils.GetCurrentPath()
	RootPath, err = filepath.Abs(BasePath + "../../")
	RootPath += "/"
	if err != nil {
		panic("Failed to get root path")
	}
}
