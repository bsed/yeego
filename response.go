package yeego

import (
	"net/http"

	"github.com/labstack/echo"
)

type ReParams struct {
	Code int         `json:"code";xml:"code"`
	Data interface{} `json:"data";xml:"data"`
	Msg  string      `json:"msg";xml:"msg"`
}

const DefaultCode = 0

var HttpStatus = http.StatusOK

var (
	//返回类型，1是json，2是xml
	ReturnType int = 1
)

type Response struct {
	Context echo.Context
	params  *ReParams
}

func NewResponse(c echo.Context) *Response {
	return &Response{Context: c, params: new(ReParams)}
}

//设置返回状态
func (resp *Response) SetStatus(status int) {
	HttpStatus = status
}

//设置返回消息
func (resp *Response) SetMsg(msg string) {
	resp.params.Msg = msg
}

//设置返回数据
func (resp *Response) SetData(data interface{}) {
	resp.params.Data = data
}

//自定义返回内容
func (resp *Response) Customsize(code int, data interface{}, msg string) error {
	resp.params.Code = code
	resp.params.Data = data
	resp.params.Msg = msg
	return resp.Ret(resp.params)
}

//返回失败
func (resp *Response) Fail(err error, code int) error {
	resp.params.Code = code
	resp.params.Msg = err.Error()
	return resp.Ret(resp.params)
}

// 返回成功的结果 默认code为1
func (resp *Response) Success(d interface{}) error {

	resp.params.Code = DefaultCode
	resp.params.Data = d

	return resp.Ret(resp.params)
}

// 返回结果
func (resp *Response) Ret(par interface{}) error {
	switch ReturnType {
	case 2:
		return resp.Context.XML(HttpStatus, par)
	default:
		return resp.Context.JSON(HttpStatus, par)
	}
}

//返回数据
func (resp *Response) Write(data []byte) {
	if _, err := resp.Context.Response().Write(data); err != nil {
		println(err)
	}
}
