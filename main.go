package request

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Middleware interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type OptionFunc func(req *HttpRequest)

var globalRequest = &HttpRequest{}

func New(options ...OptionFunc) *HttpRequest {
	req := &HttpRequest{
		headers:  globalRequest.headers,
		cookies:  globalRequest.cookies,
		username: globalRequest.username,
		password: globalRequest.password,
	}

	for _, option := range options {
		option(req)
	}

	return req
}

func SetGlobalOption(options ...OptionFunc) {
	for _, option := range options {
		option(globalRequest)
	}
}

func WithHeader(headers map[string]string) OptionFunc {
	return func(req *HttpRequest) {
		for k, v := range headers {
			req.headers[k] = v
		}
	}
}

func WithCookie(cookie ...http.Cookie) OptionFunc {
	return func(req *HttpRequest) {
		for _, c := range cookie {
			req.cookies = append(req.cookies, c)
		}
	}
}

func WithBasicAuth(username, password string) OptionFunc {
	return func(req *HttpRequest) {
		req.username = username
		req.password = password
	}
}

func WithAuthz(authorization string) OptionFunc {
	return func(req *HttpRequest) {
		req.headers["Authorization"] = authorization
	}
}

func Json(data map[string]interface{}) io.Reader {
	contents, _ := json.Marshal(data)
	return strings.NewReader(string(contents))
}

func Get(url string, params url.Values, option ...OptionFunc) (*HttpRequest, error) {
	req := New(option...)
	return req.Get(url, params)
}

func Post(url string, body io.Reader, option ...OptionFunc) (*HttpRequest, error) {
	req := New(option...)
	return req.Post(url, body)
}

func Put(url string, body io.Reader, option ...OptionFunc) (*HttpRequest, error) {
	req := New(option...)
	return req.Put(url, body)
}

func Patch(url string, body io.Reader, option ...OptionFunc) (*HttpRequest, error) {
	req := New(option...)
	return req.Patch(url, body)
}

func Options(url string, params url.Values, option ...OptionFunc) (*HttpRequest, error) {
	req := New(option...)
	return req.Options(url, params)
}

func Delete(url string, params url.Values, option ...OptionFunc) (*HttpRequest, error) {
	req := New(option...)
	return req.Delete(url, params)
}
