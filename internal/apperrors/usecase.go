package apperrors

import (
	"context"
	"encoding/json"
	"fmt"

	"mattermost-tools/internal/apperrors/repository"
)

type UseCase interface {
	NotifyError(ctx context.Context, err error) error
}

type useCase struct {
	notifyWebhookRepository repository.NotifyWebhookRepository
}

type NotifyFormat struct {
	UserName  string `json:"userName"`
	RequestId string `json:"requestId"`
	Trace     string `json:"trace"`
	Error     string `json:"error"`
}

func (u *useCase) NotifyError(ctx context.Context, err error) error {
	notify, _ := json.Marshal(
		NotifyFormat{
			Trace: "",
			Error: err.Error(),
		},
	)
	if err := u.notifyWebhookRepository.Send("g2fina-error", string(notify)); err != nil {
		return fmt.Errorf("notify error: %w", err)
	}
	return nil
}

func NewUseCase(notifyWebhookRepository repository.NotifyWebhookRepository) UseCase {
	return &useCase{
		notifyWebhookRepository: notifyWebhookRepository,
	}
}
