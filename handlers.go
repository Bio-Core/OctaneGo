package main

import (
	"fmt"
	"html/template"
	"net/http"

	keycloak "github.com/Bio-core/keycloakgo"
	"github.com/Joker/jade"
)

//Global vairable definitions
var err error

//Index returns when the main page is called and returns HTML indicating the availale paths
var indexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	tplfile, err := jade.ParseFile("./views/index.jade")
	//tplstring, _ := jade.Parse(tplfile, "doctype 5: html: body: p Hello world!")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	tpl := template.New("index").Delims("<<", ">>")
	temp, _ := tpl.Parse(tplfile)
	temp.Execute(w, nil)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return
})

var uploadHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	return
})

//handleLogin is the login function, redirects to the loginCallback function
func handleLogin(w http.ResponseWriter, r *http.Request) {
	keycloak.HandleLogin(w, r)
	return
}

//handleLoginCallback is a fuction that verifies login success and forwards to index
func handleLoginCallback(w http.ResponseWriter, r *http.Request) {
	keycloak.HandleLoginCallback(w, r)
	return
}

//authMiddleware is a middlefuntion that verifies authentication before each redirect
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	//return keycloak.AuthMiddleware(next)
	return next
}

//logout logs the user out
func logout(w http.ResponseWriter, r *http.Request) {
	keycloak.Logout(w, r)
}
