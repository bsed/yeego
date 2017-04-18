package yeego

import (
	"net/http"
	"github.com/labstack/echo"
)

const (
	// DefaultCode
	// 默认错误码0：无错误
	DefaultCode = 0
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
	Context    echo.Context
	params     *ResParams
	HttpStatus int
}

// NewResponse
// 新建一个返回
func NewResponse(c echo.Context) *Response {
	return &Response{Context: c, params: new(ResParams), HttpStatus: DefaultHttpStatus}
}

// SetRetType
// 设置返回类型
func (resp *Response) SetRetType(i int) {
	if !(i == RETJSON || i == RETXML) {
		panic("RetType 类型错误l")
	}
	ReturnType = i
}

// SetStatus
// 设置返回状态码
func (resp *Response) SetStatus(status int) {
	resp.HttpStatus = status
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

// Success
// 返回成功的结果 默认code为0
func (resp *Response) Success(d interface{}) error {
	resp.params.Code = DefaultCode
	resp.params.Data = d
	return resp.Ret(resp.params)
}

// Ret
// 返回结果
func (resp *Response) Ret(par interface{}) error {
	switch ReturnType {
	case 2:
		return resp.Context.XML(resp.HttpStatus, par)
	default:
		return resp.Context.JSON(resp.HttpStatus, par)
	}
}

// Write
// 返回row数据
func (resp *Response) Write(data []byte) error {
	_, err := resp.Context.Response().Write(data)
	return err
}
