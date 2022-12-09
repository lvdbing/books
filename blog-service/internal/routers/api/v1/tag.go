package v1

import (
	"github.com/lvdbing/books/blog-service/global"
	"github.com/lvdbing/books/blog-service/internal/service"
	"github.com/lvdbing/books/blog-service/pkg/app"
	"github.com/lvdbing/books/blog-service/pkg/convert"
	"github.com/lvdbing/books/blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context) {

}

// @Summary 获取标签列表
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param pagesize query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), Pagesize: app.GetPagesize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag err: %v", err)
		resp.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		resp.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	resp.ToResponseList(tags, totalRows)
}

func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		resp.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	resp.ToResponse(gin.H{})
}

func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		resp.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	resp.ToResponse(gin.H{})
}

func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	resp := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		resp.ToErrorResponse(errResp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err: %v", err)
		resp.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	resp.ToResponse(gin.H{})
}
