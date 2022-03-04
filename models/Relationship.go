package models

//表示用户和关注的社团
type Subscription struct {
	ID     uint `gorm:"primary_key"`
	ClubId uint `form:"club_id"json:"club_id"  ` //社团ID
	UserId uint `form:"user_id"json:"user_id"`   //用户ID
}

//通过userId关注社团
func CreateSubscription(clubId interface{}, userId interface{}) error {

	DB.Model(Subscription{}).Create(map[string]interface{}{
		"club_id": clubId,
		"user_id": userId,
	})
	return nil
}

func IsSubscribe(clubId interface{}, userId interface{}) bool {
	var s Subscription
	DB.Where("club_id=? AND user_id=?", clubId, userId).First(&s)
	if s.ID == 0 {
		return false
	}
	return true
}

//通过userId取消关注社团
func CancelSubscribe(clubId interface{}, userId interface{}) error {
	DB.Where(map[string]interface{}{
		"club_id": clubId,
		"user_id": userId,
	}).Delete(&Subscription{})
	return nil
}

//表示 收藏
type Collection struct {
	ID     uint `gorm:"primary_key"`
	PostId uint `form:"post_id"json:"post_id"` //活动ID
	UserId uint `form:"user_id"json:"user_id"` //用户ID
}

//表示社团的粉丝数和其帖子被收藏的总次数
type ClubCount struct {
	ID             uint `gorm:"primary_key"`
	ClubId         uint `form:"club_id"json:"club_id"  ` //社团ID
	NumOfFans      int  `json:"num_of_fans"`             //粉丝数
	NumOfFavorites int  `json:"num_of_favorites"`        //活动被收藏的总次数
}
