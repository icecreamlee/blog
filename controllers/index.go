package controllers

import (
	"blog/configs"
	"blog/models"
	"database/sql"
	"github.com/IcecreamLee/goutils"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"time"
)

// Index 首页&文章列表页
func Index(c *gin.Context) {
	_ = c.Request.ParseForm()
	page := goutils.ToInt(c.Request.Form.Get("page"))
	blogs := models.GetBlogList("*", page, 20)
	blogs[0].Content = string(template.HTML(blogs[0].Content))

	c.HTML(200, "posts.html", gin.H{
		"title":      "首页",
		"blogs":      blogs,
		"categories": configs.Categories,
	})
}

// Article 文章页
func Article(c *gin.Context) {
	id := goutils.ToInt(c.Param("id"))
	blog := models.GetBlog(id, "*")
	log.Println(configs.Categories)
	category := configs.Categories[0]
	if len(configs.Categories) > blog.Type {
		category = configs.Categories[blog.Type]
	}
	c.HTML(200, "page.html", gin.H{
		"title":    blog.Title,
		"blog":     blog,
		"category": category,
		"content":  template.HTML(blog.Content),
	})
}

// Manage blog列表管理
func Manage(c *gin.Context) {
	blogs := models.GetBlogList("*", 1, 100)
	c.HTML(200, "manage.html", gin.H{
		"title":         "管理",
		"blogs":         blogs,
		"categories":    configs.Categories,
		"categoriesLen": len(configs.Categories),
		"password":      configs.ManagePassword,
	})
}

// Edit 编辑blog页
func Edit(c *gin.Context) {
	id := goutils.ToInt(c.Param("id"))
	blog := models.GetBlog(id, "*")
	c.HTML(200, "edit.html", gin.H{
		"title":      "编辑",
		"blog":       blog,
		"content":    template.HTML(blog.Content),
		"categories": configs.Categories,
	})
}

// Add 新增blog页
func Add(c *gin.Context) {
	c.HTML(200, "edit.html", gin.H{
		"title":      "新增",
		"blog":       models.Blog{},
		"content":    "",
		"categories": configs.Categories,
	})
}

// AddEditSubmit 新增&编辑blog表单提交处理
func AddEditSubmit(c *gin.Context) {
	id := goutils.ToInt(c.Param("id"))
	blog := models.Blog{
		ID:       id,
		Title:    c.PostForm("title"),
		Date:     sql.NullTime{Time: time.Now()},
		Content:  c.PostForm("content"),
		Type:     goutils.ToInt(c.PostForm("type")),
		AddAt:    goutils.IntDatetime(),
		UpdateAt: goutils.IntDatetime(),
	}

	models.ORM.ShowSQL(true)
	var err error
	if id == 0 {
		_, err = models.ORM.InsertOne(&blog)
	} else {
		_, err = models.ORM.Where("id = ?", id).Cols("title,content,type").Update(&blog)
	}
	if err != nil {
		log.Println("failed to save blog: ", err.Error())
	}
	c.Redirect(302, "/manage?i="+configs.ManagePassword)
}

// Del 删除blog
func Del(c *gin.Context) {
	_, _ = models.ORM.Where("id = ?", c.Param("id")).Update(&models.Blog{IsDelete: 1})
	c.Redirect(302, "/manage?i="+configs.ManagePassword)
}

// ManagePermissionValidation 验证管理相关页面的密码不正确则显示404页
func ManagePermissionValidation(c *gin.Context) {
	if c.Query("i") != configs.ManagePassword {
		c.HTML(404, "404.html", gin.H{})
		c.Abort()
	}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
