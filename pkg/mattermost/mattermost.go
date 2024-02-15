package mattermost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"mattermost-tools/internal/config"
	"mattermost-tools/internal/core"
	"mattermost-tools/pkg/environment"
)

type Client struct {
	client http.Client
	env    environment.Environment
	config config.MattermostConfig
}

func NewClient(env environment.Environment, config config.MattermostConfig) *Client {
	return &Client{
		client: http.Client{},
		env:    env,
		config: config,
	}
}

type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type ResponseUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (c *Client) Send(channel, message string) error {
	if !c.env.IsProduction() {
		channel += "-uat"
	}
	messageBytes, _ := json.Marshal(
		Message{
			Channel: channel,
			Text:    message,
		},
	)
	req, err := http.NewRequest(http.MethodPost, c.config.WebhookUrl, bytes.NewReader(messageBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	_, err = c.client.Do(req)
	if err != nil {
		return fmt.Errorf("matter most client Send %w", err)
	}
	return nil
}

func (c *Client) GetEmojiList() ([]core.EmojiResponse, error) {
	emojiListTotal := make([]core.EmojiResponse, 0)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v4/emoji", c.config.MattermostApi), nil)
	if err != nil {
		return nil, fmt.Errorf("matter most client GetEmojiList:", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AdminToken))
	pageStep := int64(0)
	for true {
		resp, err := c.client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("matter most client GetEmojiList: %w", err)
		}
		defer func() {
			if scopedErr := resp.Body.Close(); scopedErr != nil {
				err = fmt.Errorf("matter most client GetEmojiList: %w", scopedErr)
			}
		}()
		emojiList := make([]core.EmojiResponse, 0)
		err = json.NewDecoder(resp.Body).Decode(&emojiList)
		if err != nil {
			return nil, fmt.Errorf("matter most client GetEmojiList decode: %w", err)
		}
		if len(emojiList) == 0 {
			break
		}
		emojiListTotal = append(emojiListTotal, emojiList...)
		pageStep++
		urlString := fmt.Sprintf("%s/api/v4/emoji?page=%s", c.config.MattermostApi, strconv.FormatInt(pageStep, 10))
		req, err = http.NewRequest(http.MethodGet, urlString, nil)
		if err != nil {
			return nil, fmt.Errorf("matter most client GetEmojiList:", err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AdminToken))
	}
	return emojiListTotal, nil
}

func (c *Client) GetEmojiAutoCompleteName(emojiName string) ([]core.EmojiResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v4/emoji/autocomplete?name=%s", c.config.MattermostApi, emojiName), nil)
	if err != nil {
		return nil, fmt.Errorf("matter most client GetEmojiAutoCompleteName:", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AdminToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("matter most client GetEmojiAutoCompleteName: %w", err)
	}
	defer func() {
		if scopedErr := resp.Body.Close(); scopedErr != nil {
			err = fmt.Errorf("matter most client GetEmojiAutoCompleteName: %w", scopedErr)
		}
	}()
	emojiNames := make([]core.EmojiResponse, 0)
	err = json.NewDecoder(resp.Body).Decode(&emojiNames)
	if err != nil {
		return nil, fmt.Errorf("matter most client GetMattermostUserInformation decode: %w", err)
	}
	return emojiNames, nil
}

func (c *Client) PostEmojiOnPost(emojiName, postId string) (core.EmojiPostResponse, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("%s/api/v4/reactions", c.config.MattermostApi),
		core.PostEmojiRequest{
			PostId:    postId,
			UserId:    "hujyrgux1t8ejdy934dpa7kkuy",
			EmojiName: emojiName,
		})
	if err != nil {
		return core.EmojiPostResponse{}, fmt.Errorf("matter most client PostEmojiOnPost:", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return core.EmojiPostResponse{}, fmt.Errorf("matter most client PostEmojiOnPost: %w", err)
	}
	defer func() {
		if scopedErr := resp.Body.Close(); scopedErr != nil {
			err = fmt.Errorf("matter most client PostEmojiOnPost: %w", scopedErr)
		}
	}()
	result := core.EmojiPostResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return core.EmojiPostResponse{}, fmt.Errorf("matter most client PostEmojiOnPost decode: %w", err)
	}
	return result, nil
}

func (c *Client) GetAllEmojiOfPost(postId string) ([]core.EmojiPostResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v4/posts/%s/reactions", c.config.MattermostApi, postId), nil)
	if err != nil {
		return nil, fmt.Errorf("matter most client GetAllEmojiOfPost:", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AdminToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("matter most client GetAllEmojiOfPost: %w", err)
	}
	defer func() {
		if scopedErr := resp.Body.Close(); scopedErr != nil {
			err = fmt.Errorf("matter most client GetAllEmojiOfPost: %w", scopedErr)
		}
	}()
	emojiList := make([]core.EmojiPostResponse, 0)
	err = json.NewDecoder(resp.Body).Decode(&emojiList)
	if err != nil {
		return nil, fmt.Errorf("matter most client GetAllEmojiOfPost decode: %w", err)
	}
	return emojiList, nil
}

func (c *Client) DeleteEmojiOnPost(emojiName, userId, postId string) (core.DeleteEmojiRequest, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v4/users/%s/posts/%s/reactions/%s", c.config.MattermostApi, userId, postId, emojiName), nil)
	if err != nil {
		return core.DeleteEmojiRequest{}, fmt.Errorf("matter most client DeleteEmojiOnPost:", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.AdminToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return core.DeleteEmojiRequest{}, fmt.Errorf("matter most client DeleteEmojiOnPost: %w", err)
	}
	defer func() {
		if scopedErr := resp.Body.Close(); scopedErr != nil {
			err = fmt.Errorf("matter most client DeleteEmojiOnPost: %w", scopedErr)
		}
	}()
	result := core.DeleteEmojiRequest{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return core.DeleteEmojiRequest{}, fmt.Errorf("matter most client DeleteEmojiOnPost decode: %w", err)
	}
	return result, nil
}

func (c *Client) newRequest(method, url string, body any) (*http.Request, error) {
	marshaledBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("newRequest %w", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(marshaledBody))
	if err != nil {
		return nil, fmt.Errorf("newRequest %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.UserToken))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
