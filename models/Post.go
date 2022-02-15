package models

import (
	"time"
)

//通知或动态
type Post struct {
	IsNotice  bool      `json:"is_notice"`           //是否为通知
	ID        uint      `gorm:"primarykey"json:"id"` //帖子Id
	CreatedAt time.Time `json:"created_at"`          //创建时间
	UpdatedAt time.Time `json:"updated_at"`          //更新时间(刚创建时为空)
	ClubId    uint      `json:"club_id"`             //社团Id
	Article   string    `json:"article"`             //标题
	Content   string    `json:"content"`             //内容

	ClubName string `form:"club_name"json:"club_name"` //社团名称

}
