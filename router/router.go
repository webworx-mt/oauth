package router

import (
	"net/http"
	"strings"
)

// Route represents a single route
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Router handles routing for the application
type Router struct {
	routes []Route
}

// New creates a new router
func New() *Router {
	return &Router{
		routes: make([]Route, 0),
	}
}

// GET adds a GET route
func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  "GET",
		Path:    path,
		Handler: handler,
	})
}

// POST adds a POST route
func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  "POST",
		Path:    path,
		Handler: handler,
	})
}

// PUT adds a PUT route
func (r *Router) PUT(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  "PUT",
		Path:    path,
		Handler: handler,
	})
}

// DELETE adds a DELETE route
func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  "DELETE",
		Path:    path,
		Handler: handler,
	})
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && r.matchPath(route.Path, req.URL.Path) {
			route.Handler(w, req)
			return
		}
	}

	// No route found
	http.NotFound(w, req)
}

// matchPath checks if the request path matches the route path
func (r *Router) matchPath(routePath, requestPath string) bool {
	routeParts := strings.Split(strings.Trim(routePath, "/"), "/")
	requestParts := strings.Split(strings.Trim(requestPath, "/"), "/")

	if len(routeParts) != len(requestParts) {
		return false
	}

	for i, routePart := range routeParts {
		// Check if it's a parameter (starts with :)
		if strings.HasPrefix(routePart, ":") {
			continue // Parameter matches any value
		}

		// Exact match required for non-parameters
		if routePart != requestParts[i] {
			return false
		}
	}

	return true
}

// GetParam extracts a parameter from the URL path
func GetParam(req *http.Request, paramName string) string {
	routePath := req.URL.Path
	routeParts := strings.Split(strings.Trim(routePath, "/"), "/")

	// Find the parameter index
	for i, part := range routeParts {
		if part == ":"+paramName {
			if i < len(routeParts) {
				return routeParts[i]
			}
		}
	}

	return ""
}
