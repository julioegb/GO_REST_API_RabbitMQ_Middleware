package main

import (
	"front-end-httpserver/config"
	"front-end-httpserver/server"
	"log"
	"net/http"
)

func main() {
	// Crete the New Server
	server := server.NewFrontEndServer()
	// Close connection & channel at the end of execution
	defer server.CloseRabbitmqConnection()
	go func() {
		server.StartConsume()
	}()
	log.Printf("Front End http server serving on %s port %s", config.DefaultIP, config.DefaultPort)
	http.ListenAndServe(config.DefaultIP+":"+config.DefaultPort, server)
}
