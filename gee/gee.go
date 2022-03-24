package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	key := request.Method + "_" + request.URL.Path
	handlerFunc, exist := engine.router[key]
	if !exist {
		log.Fatalf("can not find requet of method %s and url %s", request.Method, request.URL.Path)
		return
	}
	handlerFunc(response, request)
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "_" + pattern
	engine.router[key] = handlerFunc
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(address string) error {
	return http.ListenAndServe(address, engine)
}
