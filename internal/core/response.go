package core

type EmojiResponse struct {
	ID        string `json:"id"`
	CreateAt  int64  `json:"create_at"`
	UpdateAt  int64  `json:"update_at"`
	DeleteAt  int64  `json:"delete_at"`
	CreatorID string `json:"creator_id"`
	Name      string `json:"name"`
}

type EmojiPostResponse struct {
	UserID    string `json:"user_id"`
	PostID    string `json:"post_id"`
	EmojiName string `json:"emoji_name"`
	CreateAt  int64  `json:"create_at"`
	UpdateAt  int64  `json:"update_at"`
	DeleteAt  int64  `json:"delete_at"`
	RemoteID  string `json:"remote_id"`
	ChannelID string `json:"channel_id"`
}

type DeleteEmojiResponse struct {
	Status string `json:"status"`
}
