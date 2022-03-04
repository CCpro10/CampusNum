package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/models"
	"net/http"
)

//请求的参数列表
type reqParam struct {
	StudentId string `json:"student_id"form:"student_id"validate:"required,len=10|len=11"`
	Page      int    `json:"page" form:"page" example:"1"validate:"required,min=1"`         // 页码,最小为1
	Size      int    `json:"size" form:"size" example:"10"validate:"required,min=1,max=40"` //每页数据量,最大为40
}

// @Summary 获取用户关注的社团发布的最新的多条通知和动态
// @Produce json
// @Param object query reqParam true "参数列表"
// @Success 200 {object} ResponsePosts "data内有多条post"
// @Router /user/posts_from_clubs_user_fellow [get]
func GetPostsFromClubsUserFollow(c *gin.Context) {
	var request reqParam
	if e := c.ShouldBind(&request); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不符合规范,err=" + err.Error()})
		return
	}
	userId, _ := models.GetUserIdByStudentId(request.StudentId)

	posts, _ := models.GetPostsFormClubsUserFellow(userId, request.Page, request.Size)

	var resPosts []ResponsePost
	for _, post := range *posts {
		resPosts = append(resPosts, *BindPost(&post))
	}

	c.JSON(http.StatusOK, resPosts)
	return

}
