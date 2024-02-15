package handler

import (
	"encoding/json"
	"net/http"

	"mattermost-tools/internal/core"
	"mattermost-tools/internal/handler"
	"mattermost-tools/pkg/mattermost"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type MattermostHandler struct {
	logger *zerolog.Logger
	client *mattermost.Client
}

func NewMattermostHandler(logger *zerolog.Logger, client *mattermost.Client) *MattermostHandler {
	return &MattermostHandler{
		logger: logger,
		client: client,
	}
}

func (h *MattermostHandler) PostEmojiOnPost(c *gin.Context) {
	req := core.PostEmojiRequest{}
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		h.logger.Error().Err(err)
		return
	}
	emojis, err := h.client.GetEmojiAutoCompleteName(req.EmojiName)
	if err != nil {
		h.logger.Error().Err(err)
		return
	}
	for _, emoji := range emojis {
		_, err = h.client.PostEmojiOnPost(emoji.Name, req.PostId)
		h.logger.Error().Err(err)
	}
	c.JSON(
		http.StatusOK, handler.BaseResponse[string]{
			Data: "emoji posted",
		},
	)
}

//func (h *MattermostHandler) PostAllEmoji(c *gin.Context) {
//	req := core.PostEmojiRequest{}
//	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
//		h.logger.Error().Err(err)
//		return
//	}
//	emojis, err := h.client.GetEmojiList()
//	if err != nil {
//		h.logger.Error().Err(err)
//		return
//	}
//	for _, emoji := range emojis {
//		_, err = h.client.PostEmojiOnPost(emoji.Name, req.PostId)
//		h.logger.Error().Err(err)
//	}
//	c.JSON(
//		http.StatusOK, handler.BaseResponse[string]{
//			Data: "emoji posted",
//		},
//	)
//}

func (h *MattermostHandler) DeleteAllEmojiOnPost(c *gin.Context) {
	req := core.DeleteEmojiRequest{}
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		h.logger.Error().Err(err)
		return
	}
	emojis, err := h.client.GetAllEmojiOfPost(req.PostId)
	if err != nil {
		h.logger.Error().Err(err)
		return
	}
	for _, emoji := range emojis {
		_, err = h.client.DeleteEmojiOnPost(emoji.EmojiName, req.UserId, req.PostId)
		h.logger.Error().Err(err)
	}

	c.JSON(
		http.StatusOK, handler.BaseResponse[string]{
			Data: "emoji deleted",
		},
	)
}
