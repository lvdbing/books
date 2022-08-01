package model

import "github.com/lvdbing/books/blog-service/pkg/app"

type Article struct {
	*Model
	// Title 文章标题
	Title string `json:"title"`
	// Desc 文章简述
	Desc string `json:"desc"`
	// CoverImageUrl 封面图片地址
	CoverImageUrl string `json:"cover_image_url"`
	// Content 文章内容
	Content string `json:"content"`
	// State 状态 0为禁用、1为启用
	State int8 `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
