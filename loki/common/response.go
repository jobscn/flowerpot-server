package common

type Response struct {
	Code int
	Message string
	Data interface{}
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{Code: code, Message: message, Data: data}
}

func (p *Response) Success(data interface{}) *Response{
	return NewResponse(LokiCode.SUCCESS.Code(), LokiCode.SUCCESS.Message(), data)
}

func (p *Response) Failed(data interface{}) *Response{
	return NewResponse(LokiCode.FAILED.Code(), LokiCode.FAILED.Message(), data)
}

func (p *Response) CustomFailed(err *ResponseCode) *Response{
	return NewResponse(err.Code(), err.Message(), nil)
}

func (p *Response) Forbidden(data interface{}) *Response{
	return NewResponse(LokiCode.FORBIDDEN.Code(), LokiCode.FORBIDDEN.Message(), data)
}

func (p *Response) Unauthorized(data interface{}) *Response{
	return NewResponse(LokiCode.UNAUTHORIZED.Code(), LokiCode.UNAUTHORIZED.Message(), data)
}

func (p *Response) Unvalidated(data interface{}) *Response{
	return NewResponse(LokiCode.UNVALIDATED.Code(), LokiCode.UNVALIDATED.Message(), data)
}