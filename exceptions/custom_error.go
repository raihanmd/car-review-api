package exceptions

type customError struct {
	Code   int    `json:"code"`
	Errors string `json:"errors"`
}

func (e *customError) Error() string {
	return e.Errors
}

func NewCustomError(code int, errors string) *customError {
	return &customError{
		Code:   code,
		Errors: errors,
	}
}
