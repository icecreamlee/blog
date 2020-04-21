package models

import (
	"blog/internal/core"
	"log"
	"strconv"
	"time"
)

type Blog struct {
	ID       int       `xorm:"id" json:"id"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	Content  string    `json:"content"`
	Type     int       `json:"type"`
	Views    int       `json:"views"`
	IsDelete int       `xorm:"is_delete" json:"is_delete"`
	AddAt    time.Time `xorm:"add_at" json:"add_at"`
	UpdateAt time.Time `xorm:"update_at" json:"update_at"`
}

// GetBlog 返回根据id查询的单个blog信息
func GetBlog(id int, selects string) Blog {
	var blog = Blog{}
	s := "select " + selects + " from blog" +
		" where id = ? limit 1"
	_, err := ORM.SQL(s, id).Get(&blog)
	if err != nil {
		log.Printf("Failed to get blog: %s", err.Error())
	}
	return blog
}

// AddBlogViews 更新blog阅读量+1
func AddBlogViews(id int) {
	s := "update blog set views = views + 1 where id = ?"
	_, _ = ORM.Exec(s, id)
}

// GetBlog 返回当前页page的blog列表数据
func GetBlogList(category int, selects string, page int, limit int) []Blog {
	s := getBlogList(category, selects, page, limit)
	var blogs []Blog
	err := ORM.SQL(s).Find(&blogs)
	if err != nil {
		log.Printf("Failed to get blog list: %s", err.Error())
	}
	return blogs
}

// GetBlog 返回当前页page的blog列表数据
func GetBlogMapList(category int, selects string, page int, limit int) []map[string]string {
	s := getBlogList(category, selects, page, limit)
	blogs, err := ORM.QueryString(s)
	if err != nil {
		log.Printf("Failed to get blog list: %s", err.Error())
	}
	return blogs
}

func getBlogList(category int, selects string, page int, limit int) string {
	andWhere := ""
	if category > 0 {
		andWhere = " and type = " + strconv.Itoa(category)
	}
	limitLeft, limitRight := core.GetLimits(page, limit)
	s := "select " + selects + " from blog" +
		" where is_delete = 0" + andWhere +
		" order by date desc,id desc limit " + limitLeft + "," + limitRight
	return s
}
