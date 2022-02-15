package models

//通知的图片
type Picture struct {
	ID          uint   `json:"id"`
	ClubId      uint   `json:"club_id"`
	PostId      uint   `json:"post_id"`
	PictureName string `json:"picture_name"`
	PictureAddr string `json:"picture_addr"` //图片的可访问地址
}
