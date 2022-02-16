package notemodel

import (
	"FirstProject/common"
	"FirstProject/modules/user/usermodel"
)

type Note struct {
	Id      int            `json:"id,omitempty" gorm:"column:id;"`
	Title   string         `json:"title" gorm:"column:title;"`
	Content string         `json:"content" gorm:"column:content;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`

	User   *usermodel.User `json:"user" gorm:"preload:false;"`
	UserId int             `json:"-" gorm:"column:owner_id;"`
}

func (Note) TableName() string {
	return "Note"
}

type NoteCreate struct {
	Id      int            `json:"id,omitempty" gorm:"column:id;"`
	Title   string         `json:"title" gorm:"column:title;"`
	UserId  int            `json:"-" gorm:"column:owner_id;"`
	Content string         `json:"content" gorm:"column:content;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (NoteCreate) TableName() string {
	return Note{}.TableName()
}
