package common

type IResponse interface {
	Error
	GetData() interface{}
}

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{Code: code, Message: message, Data: data}
}

func (p *Response) GetCode() int {
	if p == nil {
		return 0
	}
	return p.Code
}

func (p *Response) GetMessage() string {
	if p == nil {
		return ""
	}
	return p.Message
}

func (p *Response) GetData() interface{} {
	if p == nil {
		return nil
	}
	return p.Data
}

func Success(data interface{}) *Response {
	return NewResponse(LokiCode.SUCCESS.GetCode(), LokiCode.SUCCESS.GetMessage(), data)
}

func Failed(data interface{}) *Response {
	return NewResponse(LokiCode.FAILED.GetCode(), LokiCode.FAILED.GetMessage(), data)
}

func CustomFailed(err Error) *Response {
	return NewResponse(err.GetCode(), err.GetMessage(), nil)
}

func Forbidden(data interface{}) *Response {
	return NewResponse(LokiCode.FORBIDDEN.GetCode(), LokiCode.FORBIDDEN.GetMessage(), data)
}

func Unauthorized(data interface{}) *Response {
	return NewResponse(LokiCode.UNAUTHORIZED.GetCode(), LokiCode.UNAUTHORIZED.GetMessage(), data)
}

func Unvalidated(data interface{}) *Response {
	return NewResponse(LokiCode.UNVALIDATED.GetCode(), LokiCode.UNVALIDATED.GetMessage(), data)
}
