package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteVideoService 删除视频服务
type DeleteVideoService struct {
}

func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "没找到要删除的视频",
			Error: err.Error(),
		}
	}
	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "删除视频失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{}
}
