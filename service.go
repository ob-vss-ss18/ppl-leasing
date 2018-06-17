package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/ob-vss-ss18/ppl-leasing/leasing"
)

//http://localhost:8080/graphql?query={user(id:"1"){name}}
//http://localhost:8080/graphql?query=mutation{createKunde(name: "hans", vorname: "dieter" ){kunden_id}}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, Dev Heroku!") // send data to client side
}

func main() {
	leasing.ConnectDB()

	http.HandleFunc("/", sayhelloName)                     // set router

	leasing.Init()

	http.Handle("/graphql" , leasing.ApiHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
