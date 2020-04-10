package core

import (
	"github.com/IcecreamLee/goutils"
)

var (
	BasePath string
)

func InitPaths() {
	BasePath = goutils.GetCurrentPath()
}
