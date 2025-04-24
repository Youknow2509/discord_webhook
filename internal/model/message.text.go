package model

// Message text struct
type MessageText struct {
	UserName  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	Content   string `json:"content"`
}
