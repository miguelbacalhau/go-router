Golang router
------
A very simple implementation of a router for matching URL's.

### Features
- [x] Full match
- [x] Match with parameters
- [ ] Options
- [ ] Complex Pattern match?

### Usage
To match parameters use a `:` and and identifier to identify the parameter. Example:
`/example/:parameter`

#### Code example

Create a route and then fetch the matched route with the parameter
``` go
route := router.NewRoute("/example/:parameter", "example")

router := &router.Router{
	Routes: config.GetRoutes(actions),
}

matchedRoute: = router.GetRoute(url)
```
