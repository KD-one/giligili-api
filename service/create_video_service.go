package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 创建视频服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=1,max=300"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=9999"`
}

// setSession 设置session
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "视频创建失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideo(video)}
}
