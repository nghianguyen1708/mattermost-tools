package repository

import "mattermost-tools/pkg/mattermost"

type NotifyWebhookRepository interface {
	Send(channel, message string) error
	SendDirectMessage(username, message string) error
	GetMattermostUserInformation(userEmail string) (mattermost.ResponseUser, error)
}
