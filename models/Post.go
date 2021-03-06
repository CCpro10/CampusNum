package models

import (
	"time"
)

//通知或动态
type Post struct {
	IsNotice  bool      `json:"is_notice"`           //是否为通知
	ID        uint      `gorm:"primarykey"json:"id"` //帖子Id
	CreatedAt time.Time `json:"created_at"`          //创建时间
	UpdatedAt time.Time `json:"updated_at"`          //更新时间(未修改时为创建时间)
	ClubId    uint      `json:"club_id"`             //社团Id
	Article   string    `json:"article"`             //标题
	Content   string    `json:"content"`             //内容

	ClubName string `form:"club_name"json:"club_name"` //社团名称
}

//通过Id获取post
func GetPostById(postId interface{}) (*Post, bool) {
	var p Post
	DB.Where("id =?", postId).First(&p)
	if p.ID == 0 {
		return nil, false
	}
	return &p, true
}

//获取最新的通知或动态
func GetPosts(isNotice bool, page int, size int) (Posts *[]Post, ok bool) {
	var posts []Post
	offsetNum := (page - 1) * size
	DB.Where("is_notice=?", isNotice).Limit(size).Offset(offsetNum).Order("id desc").Find(&posts)

	return &posts, true
}

//获取用户关注的社团发布的最新的多条通知/动态
func GetPostsFormClubsUserFellow(userId interface{}, page int, size int) (Posts *[]Post, ok bool) {
	var posts []Post
	offsetNum := (page - 1) * size
	var clubIds []uint
	DB.Model(Subscription{}).Select("club_id").Where("user_id=?", userId).Find(&clubIds)
	DB.Where("club_id IN ?", clubIds).Limit(size).Offset(offsetNum).Order("id desc").Find(&posts)

	return &posts, true
}

//发布活动或动态
func (post *Post) CreatePost(clubId uint, pictureIds []uint) (err error) {
	DB.Create(post)
	DB.Last(post)
	//将上传的临时图片和post绑定
	for _, id := range pictureIds {
		DB.Model(PostPicture{}).Where("id = ?", id).Update("post_id", post.ID)
	}
	//调用oss,删除未上传的临时图片

	//删除此社团用户所有未上传的临时图片
	DB.Where("id=? and post_id=?", clubId, 0).Delete(PostPicture{})

	return nil
}
