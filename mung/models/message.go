package models

import (
	"database/sql"
)

type Message struct {
	Model
	Key     string `grom:"type:varchar;size:100;unique_index; not null"json:"key"` //评论的key唯一标示
	Content string `grom:"type:string;size:100" json:"content"` //评论的内容
	Uid  int    `gorm:"type:int(10)" json:"id"` //评论人id

	//User    User   `json:"user"`//评论人
	NoteKey string `grom:"type:bytes(100);"json:"note_key"`//所属文章的key，可以为空，为空代表系统留言
	Praise  int    `gorm:"type:int(1);default:0" json:"praise"`//点赞数量
}


type Art struct {
	Model
	NickName     string         `gorm:"type:varchar(150);not null;default:'';comment:昵称" json:"name"`
	Age          uint          `gorm:"size:4;comment:年龄"`
	Phone        string         `gorm:"type:char(11);index:un_phone;comment:手机号"`
	MemberNumber string         `gorm:"type:varchar(20);index:un_phone;comment:会员编号"`
	Birthday     sql.NullString `gorm:"type:varchar(10);comment:生日"`
	ActivatedAt  sql.NullTime   `gorm:"comment:激活时间"`
}
