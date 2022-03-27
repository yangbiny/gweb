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
	return &router{router: make(map[string]HandlerFunc), roots: make(map[string]*node)}
}

func (r *router) handler(c *Context) {
	route, params := r.getRoute(c.Method, c.Path)
	if route != nil {
		c.Prams = params
		key := c.Method + "_" + route.pattern
		r.router[key](c)
	} else {
		errorMsg := fmt.Sprintf("can not find :method = %s, path = %s", c.Method, c.Path)
		c.Json(http.StatusNotFound, errorMsg)
	}
}

func (r *router) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "_" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parsePattern(pattern), 0)
	r.router[key] = handlerFunc
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	pattern := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	search := root.search(pattern, 0)
	if search != nil {
		parts := parsePattern(search.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = pattern[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(pattern[index:], "/")
				break
			}
		}
		return search, params
	}
	return nil, nil
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
