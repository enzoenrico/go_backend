package router

import "net"

type Router struct {
	Connection net.Conn
	Routes map[string]func(net.Conn) 
}

func NewRouter(conn net.Conn) *Router {
	return &Router{
		Connection: conn,
		Routes:     make(map[string]func(net.Conn)),
	}
}

func (r *Router) Handle(method string, handler func(net.Conn)){
	r.Routes[method] = handler
}