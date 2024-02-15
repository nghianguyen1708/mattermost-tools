package apperrors

import "fmt"

func ErrParamInvalid(param string) AppError {
	return New(nil, WithCode(400_0001), WithMessage(fmt.Sprintf("invalid param: %s", param)))
}

var (
	ErrInvalidInvestorId    = New(nil, WithCode(401_0002), WithMessage("invalid investor id"))
	ErrInvalidRequestStatus = New(nil, WithCode(401_0003), WithMessage("invalid request status"))
)
