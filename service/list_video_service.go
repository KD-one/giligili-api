package service

import (
	"singo/model"
	"singo/serializer"
)

// ListVideoService 展示视频列表服务
type ListVideoService struct {
}

func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "展示视频列表失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{Data: serializer.BuildVideos(videos)}
}
