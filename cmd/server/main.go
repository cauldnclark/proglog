package main

import (
	"log"

	"github.com/cauldnclark/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8888")
	log.Fatal(srv.ListenAndServe())
}
