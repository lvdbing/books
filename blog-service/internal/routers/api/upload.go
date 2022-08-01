package api

import (
	"github.com/lvdbing/books/blog-service/global"

	"github.com/lvdbing/books/blog-service/internal/service"
	"github.com/lvdbing/books/blog-service/pkg/app"
	"github.com/lvdbing/books/blog-service/pkg/convert"
	"github.com/lvdbing/books/blog-service/pkg/errcode"
	"github.com/lvdbing/books/blog-service/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	resp := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errResp := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToErrorResponse(errResp)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		resp.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		errResp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		resp.ToErrorResponse(errResp)
		return
	}
	resp.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
