package errs

type AppError struct {
	Code int
	Message string
}

func (ae AppError) Error() string {
	return ae.Message
}
