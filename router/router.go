package router

import (
	"net/http"
	"strings"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Router struct {
	routes []Route
}

func New() *Router {
	return &Router{
		routes: make([]Route, 0),
	}
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  "GET",
		Path:    path,
		Handler: handler,
	})
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  "POST",
		Path:    path,
		Handler: handler,
	})
}

func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  "PUT",
		Path:    path,
		Handler: handler,
	})
}

func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  "DELETE",
		Path:    path,
		Handler: handler,
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && r.matchPath(route.Path, req.URL.Path) {
			route.Handler(w, req)
			return
		}
	}

	http.NotFound(w, req)
}

func (r *Router) matchPath(routePath, requestPath string) bool {
	routeParts := strings.Split(strings.Trim(routePath, "/"), "/")
	requestParts := strings.Split(strings.Trim(requestPath, "/"), "/")

	if len(routeParts) != len(requestParts) {
		return false
	}

	for i, routePart := range routeParts {
		if strings.HasPrefix(routePart, ":") {
			continue
		}

		if routePart != requestParts[i] {
			return false
		}
	}

	return true
}

func GetParam(req *http.Request, paramName string) string {
	routePath := req.URL.Path
	routeParts := strings.Split(strings.Trim(routePath, "/"), "/")

	for i, part := range routeParts {
		if part == ":"+paramName {
			if i < len(routeParts) {
				return routeParts[i]
			}
		}
	}

	return ""
}
