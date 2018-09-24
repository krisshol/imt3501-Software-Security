package main

import (
	"fmt"
	"net/http"

	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/app"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
)

func main() {

	config.Init()
	fmt.Print("Starting server listening on " + config.Address + " with port " + config.Port + "\n")
	http.HandleFunc("/", app.DefaultHandler)
	http.ListenAndServe(config.Address+":"+config.Port, nil) // Start serving incomming requests. Will continue to serve forever.
}