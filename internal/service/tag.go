package service

import (
	"github.com/xiaohuazjg/blog_service/internal/model"
	"github.com/xiaohuazjg/blog_service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int8   `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int8   `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name    string `form:"name" binding:"required,min=2,max=100"`
	CeateBy string `form:"create_by" binding:"required,min=2,max=100"`
	State   string `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"required,max=100"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
	State      string `form:"state,default=1" binding:"oneof=0 1"`
}

type DeletedTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)

}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CeateBy)

}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.Name, param.State, param.ModifiedBy)

}

func (svc *Service) DeletedTag(param *DeletedTagRequest) error {
	return svc.dao.DeletedTag(param.ID)

}
