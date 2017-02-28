// Package router provides a simple router to match URL's
// with Controllers/Services/etc. Support for parameters
// and options.
package router

import (
	"reflect"
)

// Controller Interface
// TODO improve name?
type ControllerInterface interface {
}

// Route struct contains all the information for the route
// including the Controller and the Method to call
type Route struct {
	Name       string
	Controller ControllerInterface
	Method     string
}

// Router struct stores all the routes configured using
// a simple map of the url to the respective Route struct
// TODO improve mapping
type Router struct {
	Routes map[string]*Route
}

// Router Handler method, recieves a http resquest and calls
// the matching Controller method
func (router *Router) Handle(url string) string {

	if route, hasRoute := router.Routes[url]; hasRoute {
		// TODO improve calling, error prone
		response := reflect.
			ValueOf(route.Controller).
			MethodByName(route.Method).
			Call([]reflect.Value{})

		// TODO improve return value
		return response[0].String()
	}

	return ""
}
