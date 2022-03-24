package gee

import (
	"net/http"
)

type HandlerFunc func(context *Context)

type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	newContext := NewContext(response, request)
	engine.router.handler(newContext)
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	engine.router.addRoute(method, pattern, handlerFunc)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) Run(address string) error {
	return http.ListenAndServe(address, engine)
}
