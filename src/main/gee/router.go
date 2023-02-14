package gee

import (
	"log"
	"net/http"
	"strings"
)

type Router struct {
	roots map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
		roots:  make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	split := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _,s := range split {
		if s!= "" {
			parts = append(parts, s)
			if s[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)

	_,ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	parts := parsePattern(pattern)
	r.roots[method].insert(pattern, parts, 0)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) getRouter(method string, pattern string)  (*node, map[string]string){
	rootNode,ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	parts := parsePattern(pattern)
	target := rootNode.search(parts, 0)
	params := make(map[string]string)
	if target != nil {
		targetParts := parsePattern(target.pattern)
		for index, part := range targetParts {
			if part[0] == ':' {
				params[part[1:]] = parts[index]
			}
			if part[0] == '*' && len(part) > 1{
				params[part[1:]] = strings.Join(parts[index:], "/")
				break
			}
		}
		return target, params
	}
	return nil, nil
}

func (r *Router) handle(c *Context) {
	node, params := r.getRouter(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	}else{
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
	//key := c.Req.Method + "-" + c.Req.URL.Path
	//if handler, ok := r.handlers[key]; ok{
	//	handler(c)
	//}else{
	//	fmt.Fprintf(c.Writer,"404 not found %q", c.Req)
	//}
}


