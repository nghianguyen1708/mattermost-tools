package core

type PostEmojiRequest struct {
	PostId    string `json:"post_id"`
	UserId    string `json:"user_id"`
	EmojiName string `json:"emoji_name"`
	CreatedAt string `json:"created_at, default=0"`
}

type DeleteEmojiRequest struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
}
