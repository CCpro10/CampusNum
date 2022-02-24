package models

//创建图片,返回图片Id
func (p *PostPicture) CreatePicture() (newId uint, err error) {
	DB.Create(p)
	DB.Select("id").Last(p)
	newId = p.ID
	return newId, nil
}

//创建头像,返回图片Id
func (p *Avatar) CreatePicture() (newId uint, err error) {
	DB.Create(p)
	DB.Select("id").Last(p)
	newId = p.ID
	return newId, nil
}

//通过id检查图片是否为能上传的临时图片
func CheckPostPictureById(clubId uint, id uint) (ok bool) {
	var p PostPicture
	DB.Where("id = ?", id).Find(&p)
	//不是用户上传的临时图片
	if p.ClubId != clubId {
		return false
	}
	//图片不存在
	if p.ID == 0 {
		return false
	}
	//图片不是临时图片
	if p.PostId != 0 {
		return false
	}
	return true
}

//通知或动态的图片
type PostPicture struct {
	ID          uint   `json:"id"`
	ClubId      uint   `json:"club_id"`      //对应社团的Id(主键)
	PostId      uint   `json:"post_id"`      //帖子的Id
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
