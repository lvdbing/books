package model

type ArticleTag struct {
	*Model
	// ArticleId 文章ID
	ArticleId int32 `json:"article_id"`
	// TagId 标签ID
	TagId int32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
