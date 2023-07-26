package service

import (
	"singo/model"
	"singo/serializer"
)

// UpdateVideoService 修改视频服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=1,max=300"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=9999"`
}

func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "没找到要修改的视频",
			Error: err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "修改视频失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideo(video)}
}
