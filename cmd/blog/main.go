package main

import (
	"blog/internal/configs"
	"blog/internal/core"
	"blog/internal/helpers"
	"blog/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitPaths()
	gin.SetMode(configs.ENV)
	router := gin.Default()
	routers.SetRouters(router)
	router.Static("/static", core.RootPath+"web/statics")
	router.StaticFile("/favicon.ico", core.RootPath+"web/statics/favicon.ico")
	router.SetFuncMap(helpers.FuncMap())
	router.LoadHTMLGlob(core.RootPath + "web/views/*")
	router.Run(":" + configs.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
