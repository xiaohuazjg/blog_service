package model

import "github.com/xiaohuazjg/blog_service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

func (tag *Tag) TableName() string {
	return "blog_name"

}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
