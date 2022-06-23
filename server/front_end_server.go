package server

import (
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

type FrontEndServer struct {
	*mux.Router
	amqpConnection *amqp.Connection
	amqpChannel    *amqp.Channel
}

func NewFrontEndServer() *FrontEndServer {

	// Initialize RabbitMQ Connection & Channel
	conn, _ := InitRabbitMQConnection()
	ch, _ := InitRabbitMQChannel(conn)

	frontEndServer := &FrontEndServer{
		// Initialize REST API Server
		Router: mux.NewRouter().StrictSlash(true),
		// RabbitMQ connection and channel
		amqpConnection: conn,
		amqpChannel:    ch,
	}
	frontEndServer.routes()

	return frontEndServer
}

func (s *FrontEndServer) routes() {
	s.HandleFunc("/path10seconds", path10seconds).Methods("GET")
	s.HandleFunc("/path5seconds", path5seconds).Methods("GET")
	s.HandleFunc("/path2seconds", path2seconds).Methods("GET")
	s.HandleFunc("/path0seconds", path0seconds).Methods("GET")

	//s.HandleFunc("/path10seconds", path10seconds)
	//s.HandleFunc("/path5seconds", path5seconds).Methods("GET")
	//s.HandleFunc("/path2seconds", path2seconds).Methods("POST")
	//s.HandleFunc("/path0seconds", path0seconds).Methods("POST")
}
