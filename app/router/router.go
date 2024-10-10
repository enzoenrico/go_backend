package router

import (
	"fmt"
	"net"
)

type HandleFunc func(net.Conn)

type MiddleWareFunc func(HandleFunc) HandleFunc

type Route struct {
	Method  string
	Pattern string
	Handler HandleFunc
}

type Router struct {
	routes     []Route
	middleware []MiddleWareFunc
}

func NewRouter() *Router {
	return &Router{
		routes:     []Route{},
		middleware: []MiddleWareFunc{},
	}
}

func (r *Router) Use(m MiddleWareFunc) {
	r.middleware = append(r.middleware, m)
}

func (r *Router) Handle(method, pattern string, handler HandleFunc) {

	for _, m := range r.middleware {
		handler = m(handler)
	}
	r.routes = append(r.routes, Route{Method: method, Pattern: pattern, Handler: handler})

}

func (r *Router) ServeHTTP(conn net.Conn, method, path string) {
	for _, route := range r.routes {
		fmt.Println(route)
		if route.Method == method && route.Pattern == path {
			route.Handler(conn)
			return
		}
	}
}
