// Package router provides a simple router to match URL's
// with Actions and support for parameters and options.
// TODO options
package router

import (
	"errors"
	"strings"
)

// Router interface is a generic specification of the router
type RouterInterface interface {
	GetRoute(url string) (*Route, error)
}

// Route struct contains all the information for the route
// including the Controller and the Method to call
type Route struct {
	Pattern    string
	Action     string
	Method     string
}

// Router struct stores all the routes configured using
// a simple map of the url to the respective Route struct
// TODO improve mapping
type Router struct {
	Routes map[string]*Route
}

// Router Handler method, receives a url and return the
// associated route
func (router *Router) GetRoute(url string) (*Route, error) {

	var matchedRoute *Route
	foundRoute := false
	for _, route := range router.Routes {
		if matchRoute(route.Pattern, url) {
			matchedRoute = route
			foundRoute = true
		}
	}
	var err error
	if !foundRoute {
		err = errors.New("Route not found:" + url)
	}

	return matchedRoute, err
}

// Verifies if a URL matches a given pattern
// Only supports static routes and simple parameters declared with an ':'
// before the name
func matchRoute(urlPattern string, urlRecieved string) bool {
	splitUrlPattern := strings.Split(strings.Trim(urlPattern, "/"), "/")
	splitUrlRecieved := strings.Split(strings.Trim(urlRecieved, "/"), "/")

	if len(splitUrlRecieved) > len(splitUrlPattern) {
		return false
	}

	for index, urlPatternElement := range splitUrlPattern {
		if urlPatternElement[:1] == ":" {
			continue
		}
		if len(splitUrlRecieved) > index {
			if urlPatternElement != splitUrlRecieved[index] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
}
