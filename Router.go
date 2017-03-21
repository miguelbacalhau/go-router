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
// including the pattern to match, the action to call and the
// parameters
type Route struct {
	pattern    string
	action     string
	parameters map[string]string
}

// Route action getter
func (route *Route) GetAction() string {
	return route.action
}

// Route parameters getter
func (route *Route) GetParams() map[string]string {
	return route.parameters
}

// Route constructor
func NewRoute(pattern string, action string) *Route {
	return &Route{
		pattern: pattern,
		action:  action,
	}
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
		if matchRoute(route.pattern, url) {
			matchedRoute = route
			foundRoute = true
			route.parameters = extractParamFromUrl(route.pattern, url)
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

// Extracts the parameters from an URL matching a given pattern
// Assumes that the URL has already matched the pattern
func extractParamFromUrl(urlPattern string, urlRecieved string) map[string]string {
	splitUrlPattern := strings.Split(strings.Trim(urlPattern, "/"), "/")
	splitUrlRecieved := strings.Split(strings.Trim(urlRecieved, "/"), "/")

	params := make(map[string]string)

	var paramValue string
	var paramKey string
	for index, urlPatternElement := range splitUrlPattern {
		if urlPatternElement[:1] == ":" {
			paramKey = urlPatternElement[1:len(urlPatternElement)]
			if len(splitUrlRecieved) > index {
				paramValue = splitUrlRecieved[index]
			} else {
				paramValue = ""
			}
			params[paramKey] = paramValue
		}
	}

	return params
}
