package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

func CreateVideo(c *gin.Context) {
	s := service.CreateVideoService{}
	err := c.ShouldBind(&s)
	if err != nil {
		c.JSON(400, ErrorResponse(err))
	} else {
		res := s.Create()
		c.JSON(200, res)
	}

}

func ListVideo(c *gin.Context) {

}

func ShowVideo(c *gin.Context) {

}

func UpdateVideo(c *gin.Context) {

}

func DeleteVideo(c *gin.Context) {

}
