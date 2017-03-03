// Package router provides a simple router to match URL's
// with Controllers/Services/etc. Support for parameters
// and options.
package router

import (
	"errors"
)

type RouterInterface interface {
	GetRoute(url string) (*Route, error)
}

// Route struct contains all the information for the route
// including the Controller and the Method to call
type Route struct {
	Pattern string
	Action  string
	Method  string
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

	// @TODO pattern matching
	route, found := router.Routes[url]
	var err error
	if !found {
		err = errors.New("Route not found:" + url)
	}

	return route, err
}
