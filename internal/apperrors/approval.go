package apperrors

var ErrorAlreadyApproved = AppError{
	Err:     nil,
	Code:    400_010,
	Message: "already approved",
}

var ErrorStatusApproved = AppError{
	Err:     nil,
	Code:    400_010,
	Message: "approval has already been approved",
}

var ErrorUserContext = AppError{
	Err:     nil,
	Code:    401_011,
	Message: "unauthorized",
}

var ErrorDeleteGroupRole = AppError{
	Err:     nil,
	Code:    500_012,
	Message: "Cannot delete group role",
}

var ErrorDeleteGroupUser = AppError{
	Err:     nil,
	Code:    500_013,
	Message: "Cannot delete group user",
}

var ErrorDeletePolicy = AppError{
	Err:     nil,
	Code:    500_014,
	Message: "Cannot delete policy",
}

var ErrorDeleteAction = AppError{
	Err:     nil,
	Code:    500_015,
	Message: "Cannot delete action",
}

var ErrorDeleteActionPolicy = AppError{
	Err:     nil,
	Code:    500_016,
	Message: "Cannot delete action policy",
}
