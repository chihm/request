package request

import (
	"io"
	"net/http"
	"net/url"
)

type HttpRequest struct {
	headers  map[string]string
	cookies  []http.Cookie
	username string
	password string
	request  *http.Request
	response *http.Response
}

func (r *HttpRequest) setOptions(options []OptionFunc) {
	for _, option := range options {
		option(r)
	}
}

func (r *HttpRequest) buildUrl(_url string, params url.Values) (string, error) {
	if params != nil {
		url2, err := url.Parse(_url)
		if err != nil {
			return "", err
		}
		query := params.Encode()

		if url2.RawQuery != "" {
			url2.RawQuery = url2.RawQuery + "&" + query
		} else {
			url2.RawQuery = query
		}

		_url = url2.String()

		return _url, nil
	}

	return "", nil
}

func (r *HttpRequest) Request(method, _url string, body io.Reader) (*HttpRequest, error) {
	var err error

	r.request, err = http.NewRequest(method, _url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.headers {
		r.request.Header.Set(k, v)
	}

	for _, cookie := range r.cookies {
		r.request.AddCookie(&cookie)
	}

	r.response, err = http.DefaultClient.Do(r.request)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *HttpRequest) Get(_url string, params url.Values) (*HttpRequest, error) {
	newUrl, err := r.buildUrl(_url, params)
	if err != nil {
		return nil, err
	}
	return r.Request("GET", newUrl, nil)
}

func (r *HttpRequest) Post(_url string, body io.Reader) (*HttpRequest, error) {
	return r.Request("POST", _url, body)
}

func (r *HttpRequest) Options(_url string, params url.Values) (*HttpRequest, error) {
	newUrl, err := r.buildUrl(_url, params)
	if err != nil {
		return nil, err
	}
	return r.Request("OPTIONS", newUrl, nil)
}

func (r *HttpRequest) Put(_url string, body io.Reader) (*HttpRequest, error) {
	return r.Request("PUT", _url, body)
}

func (r *HttpRequest) Patch(_url string, body io.Reader) (*HttpRequest, error) {
	return r.Request("PATCH", _url, body)
}

func (r *HttpRequest) Delete(_url string, params url.Values) (*HttpRequest, error) {
	newUrl, err := r.buildUrl(_url, params)
	if err != nil {
		return nil, err
	}
	return r.Request("DELETE", newUrl, nil)
}
