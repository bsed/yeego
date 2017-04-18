/**
 * Created by WillkYang on 2017/3/10.
 */

package yeego

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func mockResponse(url string) *Response {
	e := echo.New()
	req, _ := http.NewRequest(echo.GET, url, strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return NewResponse(c)
}

func TestResponse_Success(t *testing.T) {
	resp := mockResponse("/")
	Equal(resp.Success("成功"), nil)
}

func TestResponse_SetStatus(t *testing.T) {
	resp := mockResponse("/")
	resp.SetStatus(500)
	Equal(resp.Success("成功"), nil)
	Equal(resp.Context.Response().Status, 500)
	resp = mockResponse("/")
	Equal(resp.Success("成功"), nil)
	Equal(resp.Context.Response().Status, 200)
}
