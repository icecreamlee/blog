package models

import (
	"blog/core"
	"database/sql"
	"log"
)

type Blog struct {
	ID       int `xorm:"id"`
	Title    string
	Date     sql.NullTime
	Content  string
	Type     int
	Views    int
	IsDelete int `xorm:"is_delete"`
	AddAt    int `xorm:"add_at"`
	UpdateAt int `xorm:"update_at"`
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

// GetBlog 返回当前页page的blog列表数据
func GetBlogList(selects string, page int, limit int) []Blog {
	var blogs []Blog

	limitLeft, limitRight := core.GetLimits(page, limit)
	s := "select " + selects + " from blog" +
		" where is_delete = 0" +
		" order by id desc limit " + limitLeft + "," + limitRight
	err := ORM.SQL(s).Find(&blogs)
	_, _ = ORM.Where("is_delete=0").Count(&Blog{})
	if err != nil {
		log.Printf("Failed to get blog list: %s", err.Error())
	}
	return blogs
}
