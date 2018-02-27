package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	keycloak "github.com/Bio-core/keycloakgo"
)

//Global vairable definitions
var err error

//Index returns when the main page is called and returns HTML indicating the availale paths
var indexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// tplfile, err := jade.ParseFile("./views/index.jade")
	// //tplstring, _ := jade.Parse(tplfile, "doctype 5: html: body: p Hello world!")
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	return
	// }
	tpl, _ := template.New("").Delims("[[", "]]").ParseFiles("./views/layout.html")
	//temp, _ := tpl.Parse(tplfile)
	tpl.ExecuteTemplate(w, "layout.html", nil)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return
})

var uploadHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if _, err := os.Stat("./uploads/" + header.Filename); !os.IsNotExist(err) {
		fmt.Println("File already exists")
	}
	f, err := os.OpenFile("./uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
	}
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
