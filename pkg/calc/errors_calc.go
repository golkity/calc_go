package calc

import "fmt"

type ErrorCalc struct {
	Code          int
	Error_message string
}

func NewErrorCalc(code int, error_message string) *ErrorCalc {
	return &ErrorCalc{
		Code:          code,
		Error_message: error_message,
	}
}

func (e *ErrorCalc) Error() string {
	return fmt.Sprintf(e.Error_message, e.Code)
}
