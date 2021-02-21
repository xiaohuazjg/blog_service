package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

func (tag *Tag) TableName() string {
	return "blog_name"

}
