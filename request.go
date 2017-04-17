/**
 * Created by WillkYang on 17/2/25.
 */
package yeego

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/labstack/echo"
	"github.com/yeeyuntech/yeego/validation"
)

type Request struct {
	//请求上下文
	context echo.Context
	//请求参数
	params *Param
	//请求校验
	valid validation.Validation
	//jsonTag
	jsonTag bool
	//Json
	json *Json
	//jsonParam
	jsonParam *JsonParam
}

type Param struct {
	key string
	val string
}

type JsonParam struct {
	key string
	val Json
}

//新建请求
func NewRequest(c echo.Context) *Request {
	r := &Request{context: c}
	return r
}

//通用获取参数方式，尝试get方法获取参数失败后转为post方式获取
func (req *Request) Param(key string) *Request {
	if req.GetParam(key); req.params.val == "" {
		req.PostParam(key)
	}
	return req
}

//get方式获取参数
func (req *Request) GetParam(key string) *Request {
	req.CleanParams()
	req.params.key = key
	req.params.val = req.context.QueryParam(key)
	req.jsonTag = false
	return req
}

//post方式获取参数
func (req *Request) PostParam(key string) *Request {
	req.CleanParams()
	req.params.key = key
	req.params.val = req.context.FormValue(key)
	req.jsonTag = false
	return req
}

//Json 相关
func (req *Request) SetJson(json string) {
	req.json = InitJson(json)
}

func (req *Request) GetJson() Json {
	return req.jsonParam.val
}

//获取json参数
func (req *Request) JsonParam(keys ...string) *Request {
	req.CleanParams()

	var tmpJson Json
	if req.json != nil {
		tmpJson = *(req.json)
	}

	for _, v := range keys {
		tmpJson.Get(v)
		req.jsonParam.key += v
	}
	req.jsonParam.val = tmpJson
	req.jsonTag = true
	return req
}

//设置默认参数值
func (req *Request) SetDefault(val string) *Request {
	if req.jsonTag == true {
		if req.jsonParam.val == *new(Json) {
			defJson := fmt.Sprintf(`{"index":"%s"}`, val)
			req.jsonParam.val = *InitJson(defJson).Get("index")
		}
	} else {
		if len(req.params.val) == 0 {
			req.params.val = val
		}
	}
	return req
}

//清除参数缓存
func (req *Request) CleanParams() {
	req.params = new(Param)
	req.jsonParam = new(JsonParam)
}

//清楚错误
func (req *Request) CleanError() {
	req.valid.Clear()
}

//获取参数string类型
func (req *Request) GetString() string {
	return req.getParamValue()
}
func (req *Request) GetInt() int {
	if req.getParamValue() == "" {
		return 0
	}
	if value, err := strconv.Atoi(req.getParamValue()); err != nil {
		req.valid.SetError(req.getParamKey(), "参数不是int类型，参数名称：")
		return -1
	} else {
		return value
	}
}

//获取参数并转化为float类型
func (req *Request) GetFloat() float64 {
	if req.getParamValue() == "" {
		return 0
	}
	if value, err := strconv.ParseFloat(req.getParamValue(), 64); err != nil {
		req.valid.SetError(req.getParamKey(), "参数不是float类型，参数名称：")
		return -1
	} else {
		return value
	}
}

//获取当前请求的错误信息
func (req *Request) GetError() error {
	if req.valid.HasErrors() {
		for _, err := range req.valid.Errors {
			return errors.New(err.Message + err.Key)
		}
	}
	return nil
}

//获取当前参数的值
func (req *Request) getParamValue() string {
	return req.params.val
}

//获取当前参数的键
func (req *Request) getParamKey() string {
	return req.params.key
}

//检查参数最小值
func (req *Request) Min(value int) *Request {
	if checkVal, err := strconv.Atoi(req.getParamValue()); err != nil {
		req.valid.SetError(req.getParamKey(), "参数不是int类型，参数名称：")
	} else {
		req.valid.Min(checkVal, value, req.getParamKey()).Message("参数不能小于%d，参数名称:", value)
	}
	return req
}

//检查参数最大值
func (req *Request) Max(value int) *Request {
	if checkVal, err := strconv.Atoi(req.getParamValue()); err != nil {
		req.valid.SetError(req.getParamKey(), "参数不是int类型，参数名称：")
	} else {
		req.valid.Max(checkVal, value, req.getParamKey()).Message("参数不能大于%d，参数名称:", value)
	}
	return req
}

//检查参数最小长度
func (req *Request) MinLength(length int) *Request {
	req.valid.MinSize(req.getParamValue(), length, req.getParamKey()).Message("参数长度不能小于%d，参数名称:", length)
	return req
}

//检查参数最大长度
func (req *Request) MaxLength(length int) *Request {
	req.valid.MaxSize(req.getParamValue(), length, req.getParamKey()).Message("参数长度不能大于%d，参数名称:", length)
	return req
}

//检查参数是否为手机号或固话
func (req *Request) Phone() *Request {
	req.valid.Phone(req.getParamValue(), req.getParamKey()).Message("号码格式不正确，参数名称：")
	return req
}

//检查参数是否为固话
func (req *Request) Tel() *Request {
	req.valid.Tel(req.getParamValue(), req.getParamKey()).Message("固话号码格式不正确，参数名称：")
	return req
}

//检查参数是否为手机号
func (req *Request) Mobile() *Request {
	req.valid.Mobile(req.getParamValue(), req.getParamKey()).Message("手机号码格式不正确，参数名称：")
	return req
}

//检查参数是否为Email
func (req *Request) Email() *Request {
	req.valid.Email(req.getParamValue(), req.getParamKey()).Message("邮箱地址格式不正确，参数名称：")
	return req
}

//检查参数是否为邮政编码
func (req *Request) ZipCode() *Request {
	req.valid.ZipCode(req.getParamValue(), req.getParamKey()).Message("邮政编码格式不正确，参数名称：")
	return req
}

//检查参数是否为数字
func (req *Request) Numeric() *Request {
	req.valid.Numeric(req.getParamValue(), req.getParamKey()).Message("数字格式不正确，参数名称：")
	return req
}

//检查参数是否为Alpha字符
func (req *Request) Alpha() *Request {
	req.valid.Alpha(req.getParamValue(), req.getParamKey()).Message("Alpha格式不正确，参数名称：")
	return req
}

//检查参数是否为数字或Alpha字符
func (req *Request) AlphaNumeric() *Request {
	req.valid.AlphaNumeric(req.getParamValue(), req.getParamKey()).Message("AlphaNumeric格式不正确，参数名称：")
	return req
}

//检查参数是否为Alpha字符或数字或横杠-_
func (req *Request) AlphaDash() *Request {
	req.valid.AlphaDash(req.getParamValue(), req.getParamKey()).Message("AlphaDash格式不正确，参数名称：")
	return req
}

//检查参数是否为IP地址
func (req *Request) IP() *Request {
	req.valid.IP(req.getParamValue(), req.getParamKey()).Message("IP地址格式不正确，参数名称：")
	return req
}

//检查参数是否匹配正则
func (req *Request) Match(match string) *Request {
	req.valid.Match(req.getParamValue(), regexp.MustCompile(match), req.getParamKey()).Message("正则校验失败，参数名称：")
	return req
}

//检查参数是否不匹配正则
func (req *Request) NoMatch(match string) *Request {
	req.valid.NoMatch(req.getParamValue(), regexp.MustCompile(match), req.getParamKey()).Message("正则校验失败，参数名称：")
	return req
}
