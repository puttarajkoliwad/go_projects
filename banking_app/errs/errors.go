package errs

type AppError struct {
	Code int	`json:",omitempty"`
	Message string	`json:"message"`
}

func (ae AppError) AsMessage() (*AppError) {
	return &AppError{
		Message: ae.Message,
	}
}

func (ae AppError) Error() string {
	return ae.Message
}
