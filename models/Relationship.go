package models

//表示用户和关注的社团
type Subscription struct {
	ID     uint `gorm:"primary_key"`
	ClubId uint `form:"club_id"json:"club_id"  ` //社团ID
	UserId uint `form:"user_id"json:"user_id"`   //用户ID
}

//表示 收藏
type Collection struct {
	ID     uint `gorm:"primary_key"`
	PostId uint `form:"post_id"json:"post_id"` //活动ID
	UserId uint `form:"user_id"json:"user_id"` //用户ID
}
