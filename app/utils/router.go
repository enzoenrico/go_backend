package utils

import (
	"net"
	"net/http"

	"github.com/codecrafters-io/http-server-starter-go/app/responses"
)

// needs a connection and a path for a pattern string
type HandlerFunc func(net.Conn, string)

//single route properties
type RouteEntry struct {
	Path string
	Method string 
	Handler http.HandlerFunc 
}

//stores all of the paths with their perspective methods and handlers
type Router struct{
	routes []RouteEntry
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc){
	e:= RouteEntry{
		Method: method,
		Path: path,
		Handler: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
}


func (re *RouteEntry) Match(r *http.Request) bool {
	if r.Method != re.Method || r.URL.Path != re.Path{
		return false
	}
	return true
}


func (rtr *Router) ServeHTTP(r *http.Request, conn net.Conn) {
	for _, e := range rtr.routes{
		match := e.Match(r)
		if !match{
			continue
		}
        e.Handler.ServeHTTP()
	}

	responses.NotFound(conn)
}


