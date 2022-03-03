package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"main/models"
	"net/http"
	"strconv"
	"time"
)

type ResponseClubInfo struct {
	ID             uint      `gorm:"primarykey"json:"id"`
	CreatedAt      time.Time `form:"created_at"json:"created_at"`
	AvatarAddr     string    `form:"avatar_addr"json:"avatar_addr"`   //社团头像url地址
	Introduction   string    `form:"introduction"json:"introduction"` //社团简介
	ClubName       string    `form:"club_name"json:"club_name"`       //社团名称
	NumOfFans      int       `json:"num_of_fans"`                     //粉丝数
	NumOfFavorites int       `json:"num_of_favorites"`                //活动被收藏的总次数
}

// @Summary 查看社团信息,可用于展示社团的页面
// @Produce json
// @Param  club_id query uint true "社团id"
// @Success 200 {object} ResponseClubInfo
// @Router /user/club_info [get]
func ShowClubInfo(c *gin.Context) {
	req := c.Query("club_id")
	reqClubId, e := strconv.ParseUint(req, 10, 32)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "club_id参数格式错误,err=" + e.Error()})
		return
	}

	if !models.ExistClub("id", reqClubId) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "此社团不存在"})
		return
	}
	//获取信息
	clubInfo := models.GetClubInfoById(reqClubId)
	var rsp ResponseClubInfo
	v, _ := json.Marshal(clubInfo)
	_ = json.Unmarshal(v, &rsp)

	c.JSON(http.StatusOK, rsp)
	return

}
