// server/server.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// set up handler to listen to root path
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("new request")
		fmt.Fprintf(writer, "hello world \n")
	})

	// serve on port 9090 of local host
	server := http.Server{
		Addr:    ":9090",
		Handler: handler,
	}

	/*	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error listening to port: %v", err)
	}*/
	// /home/user/go/src/github.com/sheikh-arman/mtls/server/server.go
	// serve the endpoint with tls encryption
	if err := server.ListenAndServeTLS("/home/user/go/src/github.com/sheikh-arman/mtls/cert/server.crt", "/home/user/go/src/github.com/sheikh-arman/mtls/cert/server.key"); err != nil {
		log.Fatalf("error listening to port: %v", err)
	}

}
