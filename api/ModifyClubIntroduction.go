package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/models"
	"net/http"
	"strings"
)

type Response struct {
	Msg string `json:"msg"` //返回的信息
}

// @Summary 修改社团简介
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌 例:Bearer fbhraewifvg43uwerfaewobf"
// @Param introduction formData string true "社团简介"
// @Success 200 {object} Response
// @Router /club/introduction [put]
func ModifyClubIntroduction(c *gin.Context) {
	//   绑定参数
	reqIntroduction := c.DefaultPostForm("introduction", "")
	//验证参数
	validate := validator.New()
	reqIntroduction = " " + strings.TrimSpace(reqIntroduction)
	e := validate.Var(reqIntroduction, "max=200")
	if e != nil {
		c.JSON(http.StatusBadRequest, Response{
			Msg: "请求参数格式错误,检查是否过长",
		})
		return
	}
	//获取用户信息
	ClubId, ok := c.Get("ClubId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}

	//根据club_id修改数据
	ok = models.ModifyClubIntroductionById(ClubId, reqIntroduction)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "简介修改失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "修改成功"})
	return
}
