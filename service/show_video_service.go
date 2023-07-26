package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowVideoService 展示视频详情服务
type ShowVideoService struct {
}

func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "展示视频详情失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideo(video)}
}
