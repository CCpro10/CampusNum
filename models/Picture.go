package models

//通知的图片
type Picture struct {
	ID          uint   `json:"id"`
	ClubId      uint   `json:"club_id"` //对应社团的Id
	PostId      uint   `json:"post_id"`
	PictureName string `json:"picture_name"` //图片在oss的完整路径及名称
	PictureAddr string `json:"picture_addr"` //图片的可访问地址
}

//头像
type Avatar struct {
	ID          uint   `json:"id"`
	ClubId      uint   `json:"club_id"`      //对应社团的Id
	PictureName string `json:"picture_name"` //图片在oss的完整路径及名称
	PictureAddr string `json:"picture_addr"` //图片的可访问地址
}
