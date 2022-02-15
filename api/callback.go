package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PicCallback struct {
	PictureID   uint   `json:"picture_id" form:"picture_id" binding:"required"`
	Size        uint64 `json:"size" form:"size" binding:"required"`
	PictureName uint   `json:"picture_name" form:"picture_name" binding:"required"`
}

func Callback(c *gin.Context) {
	var p PicCallback

	if e := c.ShouldBind(&p); e != nil {
		log.Println(e)
	}

	c.JSON(http.StatusOK, gin.H{"200": "OK", "picture_id": p.PictureID})

}
