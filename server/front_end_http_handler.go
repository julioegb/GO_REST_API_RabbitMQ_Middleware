package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type ResponseSlide []Response

func path10seconds(w http.ResponseWriter, r *http.Request) {
	oneResponse := ResponseSlide{
		Response{
			Title:   "10 seconds request",
			Desc:    "Wait 10 seconds for response",
			Content: "This request waited 10 seconds in the Server",
		},
	}
	log.Printf("Current time in the server: %v", time.Unix(time.Now().Unix(), 0))
	log.Printf("Request Received: Start waiting 10 seconds")
	time.Sleep(10 * time.Second)
	log.Printf("Current time in the server: %v", time.Unix(time.Now().Unix(), 0))
	log.Printf("Time finished, sending following response: %v", oneResponse)
	json.NewEncoder(w).Encode(oneResponse)
}

func path5seconds(w http.ResponseWriter, r *http.Request) {
	oneResponse := ResponseSlide{
		Response{
			Title:   "5 seconds request",
			Desc:    "Wait 5 seconds for response",
			Content: "This request waited 5 seconds in the Server",
		},
	}
	log.Printf("Current time in the server: %v", time.Unix(time.Now().Unix(), 0))
	log.Printf("Request Received: Start waiting 5 seconds")
	time.Sleep(5 * time.Second)
	log.Printf("Current time in the server: %v", time.Unix(time.Now().Unix(), 0))
	log.Printf("Time finished, sending following response: %v", oneResponse)
	json.NewEncoder(w).Encode(oneResponse)
}

func path2seconds(w http.ResponseWriter, r *http.Request) {
	oneResponse := ResponseSlide{
		Response{
			Title:   "2 seconds request",
			Desc:    "Wait 2 seconds for response",
			Content: "This request waited 2 seconds in the Server",
		},
	}
	log.Printf("Current time in the server: %v", time.Unix(time.Now().Unix(), 0))
	log.Printf("Request Received: Start waiting 2 seconds")
	time.Sleep(2 * time.Second)
	log.Printf("Current time in the server: %v", time.Unix(time.Now().Unix(), 0))
	log.Printf("Time finished, sending following response: %v", oneResponse)
	json.NewEncoder(w).Encode(oneResponse)
}

func path0seconds(w http.ResponseWriter, r *http.Request) {
	oneResponse := ResponseSlide{
		Response{
			Title:   "Immediate request",
			Desc:    "No wait for response",
			Content: "This request was answered immediately.",
		},
	}
	log.Printf("Current time in the server: %v", time.Unix(time.Now().Unix(), 0))
	log.Printf("Request Received")
	log.Printf("Sending following response: %v", oneResponse)
	json.NewEncoder(w).Encode(oneResponse)
}
