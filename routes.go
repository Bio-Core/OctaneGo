package main

import "net/http"

//Route object creates to keep track of routes for router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is an array of Route objects
type Routes []Route

var routes = Routes{
	//Home page
	///Authenticated
	Route{
		"Index",
		"GET",
		"/octane/main",
		authMiddleware(indexHandler),
	},
	//Favicon
	///Unauthenticated
	Route{
		"Index",
		"GET",
		"/octane/favicon.ico",
		faviconHandler,
	},
	//Login Page
	///unAuthenticated
	Route{
		"Login Page",
		"GET",
		"/octane/",
		loginHandler,
	},
	//HTML list of people
	///Authenticated
	Route{
		"Upload",
		"POST",
		"/octane/upload",
		authMiddleware(uploadHandler),
	},
	//Login page
	///Unauthenticated
	Route{
		"handleLogin",
		"GET",
		"/octane/login",
		authMiddleware(handleLogin),
	},
	//Logout, redirects to login
	///Unauthenticatec
	Route{
		"logout",
		"GET",
		"/octane/logout",
		logout,
	},
}
