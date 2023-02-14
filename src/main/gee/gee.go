package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix string
		middleWares []HandlerFunc
		engine *Engine
	}

	Engine struct {
		*RouterGroup
		router *Router
		groups []*RouterGroup
	}
)

func New() *Engine{
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string)  *RouterGroup{
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRouter(method string, path string, handler HandlerFunc)  {
	pattern := group.prefix + path
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handler)
}

func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRouter("GET", pattern, handler)
}

func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRouter("POST", pattern, handler)
}

func (group *RouterGroup) User(middleWare ...HandlerFunc)  {
	group.middleWares = append(group.middleWares, middleWare...)
}

func (engine *Engine) Run(addr string) (error error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	var middleWares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middleWares = append(middleWares, group.middleWares...)
		}
	}
	context := newContext(w, req)
	context.handlers = middleWares
	engine.router.handle(context)
}



