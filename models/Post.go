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
