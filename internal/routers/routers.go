package routers

import (
	"blog/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	r.GET("/", controllers.Index)
	r.GET("/index", controllers.Index)
	r.GET("/posts", controllers.Index)
	r.GET("/posts/:id", controllers.Article)
	r.GET("/manage", controllers.ManagePermissionValidation, controllers.Manage)
	r.GET("/add", controllers.ManagePermissionValidation, controllers.Add)
	r.GET("/edit/:id", controllers.ManagePermissionValidation, controllers.Edit)
	r.POST("/edit/:id", controllers.ManagePermissionValidation, controllers.AddEditSubmit)
	r.GET("/del/:id", controllers.ManagePermissionValidation, controllers.Del)
	r.GET("/ueditor", controllers.UEditor)
	r.POST("/ueditor", controllers.UEditor)
	r.GET("/ping", controllers.Ping)
}
