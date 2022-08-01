package app

import (
	"github.com/lvdbing/books/blog-service/global"
	"github.com/lvdbing/books/blog-service/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		page = 1
	}
	return page
}

func GetPagesize(c *gin.Context) int {
	pagesize := convert.StrTo(c.Query("pagesize")).MustInt()
	if pagesize <= 0 {
		return global.AppSetting.DefaultPagesize
	}
	if pagesize > global.AppSetting.MaxPagesize {
		return global.AppSetting.MaxPagesize
	}
	return pagesize
}

func GetPageOffset(page, pagesize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pagesize
	}
	return result
}
