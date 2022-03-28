package gee

import (
	"net/http"
)

type HandlerFunc func(context *Context)

type RouteGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouteGroup
	engine      *Engine
}

type Engine struct {
	*RouteGroup
	router *router
	groups []*RouteGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (group *RouteGroup) Group(prefix string) *RouteGroup {
	engine := group.engine
	newGroup := &RouteGroup{
		engine: engine,
		prefix: group.prefix + prefix,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (engine *Engine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	newContext := NewContext(response, request)
	engine.router.handler(newContext)
}

func (group *RouteGroup) addRoute(method string, comp string, handlerFunc HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handlerFunc)
}

func (group *RouteGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouteGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(address string) error {
	return http.ListenAndServe(address, engine)
}
