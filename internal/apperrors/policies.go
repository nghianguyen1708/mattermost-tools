package apperrors

import "fmt"

func ErrInvalidRule(rule string) AppError {
	return New(nil, WithCode(500_1001), WithMessage(fmt.Sprintf("invalid rule: %s", rule)))
}
