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
	s := service.ListVideoService{}
	err := c.ShouldBind(&s)
	if err != nil {
		c.JSON(400, ErrorResponse(err))
	} else {
		res := s.List()
		c.JSON(200, res)
	}
}

func ShowVideo(c *gin.Context) {
	s := service.ShowVideoService{}
	res := s.Show(c.Param("id"))
	c.JSON(200, res)

}

func UpdateVideo(c *gin.Context) {
	s := service.UpdateVideoService{}
	err := c.ShouldBind(&s)
	if err != nil {
		c.JSON(400, ErrorResponse(err))
	} else {
		res := s.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

func DeleteVideo(c *gin.Context) {
	s := service.DeleteVideoService{}
	err := c.ShouldBind(&s)
	if err != nil {
		c.JSON(400, ErrorResponse(err))
	} else {
		res := s.Delete(c.Param("id"))
		c.JSON(200, res)
	}
}
