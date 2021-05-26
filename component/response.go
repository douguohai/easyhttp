package component

import (
	"io/ioutil"
	"net/http"
)

type BuildResponse func(resp *http.Response, err error) IResponse

//使用client发去请求后,返回一个实现了这个接口的对象
//只要实现这个接口,就能作为返回值
//在 `BuildResponse` 函数中构造出返回的对象
//默认提供了 `HttpResponse` 实现了这个接口
//可以根据自己的需求自己重新实现这个接口
type IResponse interface {
	//返回这个请求的错误
	Error() error

	//返回这个请求的http状态码
	StatusCode() int

	//返回HTTP请求的header信息
	Header() http.Header

	//返回HTTP内容长度
	ContentLength() int64

	//返回HTTP的内容
	Content() []byte

	//返回HTTP包中的 response信息
	Resp() *http.Response

	//返回这次请求的request信息
	Request() *http.Request

	//根据name 返回response的cookie
	Cookie(name string) *http.Cookie
}

type HttpResponse struct {
	err             error
	ResponseContent []byte
	HttpResp        *http.Response
}

func (h *HttpResponse) Error() error {
	return h.err
}

func (h *HttpResponse) StatusCode() int {
	if h.HttpResp == nil {
		return 0
	}
	return h.HttpResp.StatusCode
}

func (h *HttpResponse) Header() http.Header {
	if h.HttpResp == nil {
		return nil
	}
	return h.HttpResp.Header
}

func (h *HttpResponse) ContentLength() int64 {
	if h.HttpResp == nil {
		return 0
	}
	return h.HttpResp.ContentLength
}

func (h *HttpResponse) Content() []byte {
	return h.ResponseContent
}

func (h *HttpResponse) Resp() *http.Response {
	return h.HttpResp
}

func (h *HttpResponse) Request() *http.Request {
	if h.HttpResp == nil {
		return nil
	}
	return h.HttpResp.Request
}

func (h *HttpResponse) Cookie(name string) *http.Cookie {
	for _, cookie := range h.HttpResp.Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

//默认构造HTTP response的函数
func EasyBuildResponse(resp *http.Response, err error) IResponse {
	response := new(HttpResponse)
	if err != nil {
		response.err = err
		return response
	}
	response.HttpResp = resp

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.err = err
		return response
	}
	response.ResponseContent = all
	_ = resp.Body.Close()
	return response
}
