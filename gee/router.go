package gee

import (
	"fmt"
	"net/http"
)

type router struct {
	router map[string]HandlerFunc
}

func newRouter() *router {
	return &router{router: make(map[string]HandlerFunc)}
}

func (router *router) handler(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := router.router[key]; ok {
		handler(c)
	} else {
		errorMsg := fmt.Sprintf("can not find :method = %s, path = %s", c.Method, c.Path)
		c.Json(http.StatusNotFound, errorMsg)
	}
}

func (router *router) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "_" + pattern
	router.router[key] = handlerFunc
}
