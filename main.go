package main

import (
	"blog/configs"
	"blog/core"
	"blog/helpers"
	"blog/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitPaths()
	gin.SetMode(configs.ENV)
	router := gin.Default()
	routers.SetRouters(router)
	router.Static("/static", "./statics")
	router.StaticFile("/favicon.ico", "./statics/favicon.ico")
	router.SetFuncMap(helpers.FuncMap())
	router.LoadHTMLGlob("views/*")
	router.Run(":" + configs.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
