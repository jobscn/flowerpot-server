package common

import "fmt"

type Error interface {
	GetCode() int
	GetMessage() string
}

type _Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *_Error {
	err := &_Error{Code: code, Message: message}
	RegisterErrorCode(err)
	return err
}

func (p *_Error) GetCode() int {
	if p == nil {
		return 0
	}
	return p.Code
}

func (p *_Error) GetMessage() string {
	if p == nil {
		return ""
	}
	return p.Message
}

func (p *_Error) Error() string {
	return fmt.Sprintf("code = %d; msg = %s", p.GetCode(), p.GetMessage())
}

var GlobalError map[int]Error

func RegisterErrorCode(errs ...Error) {
	if GlobalError == nil {
		GlobalError = make(map[int]Error)
	}
	for _, err := range errs {
		if value, ok := GlobalError[err.GetCode()]; ok {
			panic(fmt.Errorf("already exists error code: %v, registed by: %+v, %+v", err.GetCode(), value, err))
		}

		GlobalError[err.GetCode()] = err
	}
}
