package main

import (
	"flag"
	"log"
	"net/http"

	keycloak "github.com/Bio-core/keycloakgo"
)

//hosting variables defined in flags
var localhost string
var localport string
var keycloakhost string
var keycloakport string
var server string
var keycloakserver string

var goTest bool // true if unit tests are running

func main() {
	flag.StringVar(&localport, "p", "3000", "Specify which port to use")
	flag.StringVar(&localhost, "host", "localhost", "Specify the name of the host")
	flag.Parse()

	server = "http://" + localhost + ":" + localport
	keycloakserver = "https://oa.pmgenomics.ca"

	addKeycloak(keycloakserver, server)

	router := NewRouter()
	//Stats hosting on the constant port
	log.Fatal(http.ListenAndServe(":"+localport, router))
}

func addKeycloak(keycloakserver, server string) {
	keycloak.Init(keycloakserver, "https://www.pmgenomics.ca", "/octane/main", "/octane/logout")
	//keycloak.Init(keycloakserver, "http://127.0.0.1:3000", "/octane/main", "/octane/logout")
}
