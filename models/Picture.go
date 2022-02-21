package models

func (p *PostPicture) CreatePicture() (newId uint, err error) {
	DB.Create(p)
	DB.Select("id").Last(p)
	newId = p.ID
	return newId, nil
}

func (p *Avatar) CreatePicture() (newId uint, err error) {
	DB.Create(p)
	DB.Select("id").Last(p)
	newId = p.ID
	return newId, nil
}

//通知或动态的图片
type PostPicture struct {
	ID          uint   `json:"id"`
	ClubId      uint   `json:"club_id"`      //对应社团的Id
	PostId      uint   `json:"post_id"`      //帖子的Id
	PictureName string `json:"picture_name"` //图片在oss的完整路径及名称
	PictureAddr string `json:"picture_addr"` //图片的可访问地址
}

func GetNewPictureId() uint {
	var p PostPicture
	DB.Select("id").Last(&p)
	return p.ID + 1
}

//头像
type Avatar struct {
	ID          uint   `json:"id"`
	ClubId      uint   `json:"club_id"`      //对应社团的Id
	PictureName string `json:"picture_name"` //图片在oss的完整路径及名称
	PictureAddr string `json:"picture_addr"` //图片的可访问地址
}
