package serializer

import "singo/model"

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	return Video{
		ID:        video.ID,
		Title:     video.Title,
		Info:      video.Info,
		CreatedAt: video.CreatedAt.Unix(),
	}
}

// BuildVideo 序列化视频列表
func BuildVideos(items []model.Video) []Video {
	var videos []Video
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
