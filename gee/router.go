package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	router map[string]HandlerFunc
	roots  map[string]*node
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
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(pattern, parsePattern(pattern), 0)
	router.router[key] = handlerFunc
}

func parsePattern(pattern string) []string {
	if len(pattern) == 0 {
		return nil
	}
	sp := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, value := range sp {
		if value == "" {
			continue
		}
		parts = append(parts, value)
		if value[0] == '*' {
			break
		}
	}
	return parts
}
