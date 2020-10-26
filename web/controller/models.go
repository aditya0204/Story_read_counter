package controller

type Story struct {
	StoryTitle string
	StoryContent string
	ImageLink string
	Views int
	Username string
}

type Stories struct {
	Story []Story
}