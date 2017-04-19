package yeego

import (
	"net/http"
	"github.com/labstack/echo"
)

const (
	// DefaultCode
	// 默认错误码0：无错误
	DefaultCode = 0
	// 默认有错误的错误码
	DefaultErrorCode = -1
	// DefaultHttpStatus
	// 默认的http状态码200
	DefaultHttpStatus = http.StatusOK
)

// ResParams
// 返回的参数封装
type ResParams struct {
	Code int         `json:"code";xml:"code"` // 错误码
	Data interface{} `json:"data";xml:"data"` // 数据
	Msg  string      `json:"msg";xml:"msg"`   // 消息
}

// Response
// 对返回的组合封装
type Response struct {
	context    echo.Context
	params     *ResParams
	httpStatus int
}

// NewResponse
// 新建一个返回
func NewResponse(c echo.Context) *Response {
	return &Response{context: c, params: new(ResParams), httpStatus: DefaultHttpStatus}
}

// Context
// 返回resp的Context
func (resp *Response) Context() echo.Context {
	return resp.context
}

// SetStatus
// 设置返回状态码
func (resp *Response) SetStatus(status int) {
	resp.httpStatus = status
}

// SetMsg
// 设置返回消息
func (resp *Response) SetMsg(msg string) {
	resp.params.Msg = msg
}

// SetData
// 设置返回数据
func (resp *Response) SetData(data interface{}) {
	resp.params.Data = data
}

// Customise(
// 自定义返回内容
func (resp *Response) Customise(code int, data interface{}, msg string) error {
	resp.params.Code = code
	resp.params.Data = data
	resp.params.Msg = msg
	return resp.Ret(resp.params)
}

// Fail
// 返回失败
func (resp *Response) Fail(err error, code int) error {
	resp.params.Code = code
	resp.params.Msg = err.Error()
	return resp.Ret(resp.params)
}

// FailWithDefaultErrorCode
// 默认错误码返回错误
func (resp *Response) FailWithDefaultErrorCode(err error) error {
	resp.params.Code = DefaultErrorCode
	resp.params.Msg = err.Error()
	return resp.Ret(resp.params)
}

// Success
// 返回成功的结果 默认code为0
func (resp *Response) Success(d interface{}) error {
	resp.params.Code = DefaultCode
	resp.params.Data = d
	return resp.Ret(resp.params)
}

// SuccessWithMsg
// 成功并返回msg
func (resp *Response) SuccessWithMsg(msg string) error {
	resp.params.Code = DefaultCode
	resp.params.Msg = msg
	return resp.Ret(resp.params)
}

// Ret
// 返回结果
func (resp *Response) Ret(par interface{}) error {
	switch ReturnType {
	case 2:
		return resp.context.XML(resp.httpStatus, par)
	default:
		return resp.context.JSON(resp.httpStatus, par)
	}
}

// Write
// 返回row数据
func (resp *Response) Write(data []byte) error {
	_, err := resp.context.Response().Write(data)
	return err
}
