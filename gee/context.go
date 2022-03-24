package gee

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
}

func (c *Context) Json(status int, obj ...interface{}) {
	c.SetHeader("content-Type", "application/json")
	c.SetStatus(status)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) SetHeader(headerName string, headerValue string) {
	c.Writer.Header().Set(headerName, headerValue)
}

func (c *Context) SetStatus(status int) {
	c.StatusCode = status
	c.Writer.WriteHeader(status)
}

func NewContext(w http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: request,
		Path:    request.URL.Path,
		Method:  request.Method,
	}
}
