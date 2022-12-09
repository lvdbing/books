package model

import (
	"github.com/lvdbing/books/blog-service/pkg/app"

	"gorm.io/gorm"
)

type Tag struct {
	*Model
	// Name 标签名称
	Name string `json:"name"`
	// State 状态 0为禁用、1为启用
	State uint8 `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State)
	err := db.Model(&t).Where("is_del=?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pagesize int) ([]*Tag, error) {
	var tags []*Tag
	var err error

	if pageOffset >= 0 && pagesize > 0 {
		db = db.Offset(pageOffset).Limit(pagesize)
	}
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State)
	err = db.Where("is_del=?", 0).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	db = db.Model(&t).Where("id=? and is_del=?", t.Id, 0)
	return db.Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del=?", t.Model.Id, 0).Delete(&t).Error
}
