package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"testing"
)

// To use Wait Groups for Goroutines
var wg sync.WaitGroup

// To store the number of responses
var responses10Seconds int
var responses5Seconds int
var responses2Seconds int
var responses0Seconds int

// This function is a http client that sends GET request to test the server responses
// path: is the REST API route of the server
// countResponse: is a pointer that points to the variable who stores the number of responses
func httpClient(path string, countResponse *int) {
	log.Printf("New Gorutine. Goroutine number: %v", runtime.NumGoroutine())
	resp, err := http.Get("http://localhost:8089/" + path)
	if err != nil {
		log.Fatalln(err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		stringResponse := string(body)
		log.Printf("Response Body: %s", stringResponse)
		*countResponse++
	}
	// Is important to close the client connection at the end
	defer resp.Body.Close()
}

// This test is going to send 400 requests to the front-end-httpserver web server
// 100 request that each request is going to wait 10 seconds
// 100 request that each request is going to wait 5 seconds
// 100 request that each request is going to wait 2 seconds
// 100 request for immediate response
func TestFrontEndHTTPServer(t *testing.T) {
	// Find print system information
	log.Printf("*** Server Information ***")
	log.Printf("OS: %v", runtime.GOOS)
	log.Printf("Architecture: %v", runtime.GOARCH)
	log.Printf("CPUs: %v", runtime.NumCPU())
	log.Printf("Goroutines: %v", runtime.NumGoroutine())
	for i := 0; i < 400; i++ {
		// wg.Add(1): Before create a new Goroutine increase the value in the Wait Group (wg)
		// wg.Done(): We decrease one in the Wait Group saying we finish the GO routine
		if i < 100 {
			wg.Add(1)
			go func() {
				httpClient("/path10seconds", &responses10Seconds)
				wg.Done()
			}()
		} else if i < 200 {
			wg.Add(1)
			go func() {
				httpClient("/path5seconds", &responses5Seconds)
				wg.Done()
			}()
		} else if i < 300 {
			wg.Add(1)
			go func() {
				httpClient("/path2seconds", &responses2Seconds)
				wg.Done()
			}()
		} else if i < 400 {
			wg.Add(1)
			go func() {
				httpClient("/path0seconds", &responses0Seconds)
				wg.Done()
			}()
		}
	}
	// Wait for all Goroutines to finish
	wg.Wait()
	log.Printf("Program is going to finish. Number of Goroutines: %v", runtime.NumGoroutine())
	// Evaluate the quantity of responses to ensure all responses were received
	assert.Equal(t, 100, responses0Seconds, "Some immediate responses MISSING")
	assert.Equal(t, 100, responses2Seconds, "Some 2 seconds responses MISSING")
	assert.Equal(t, 100, responses5Seconds, "Some 5 seconds responses MISSING")
	assert.Equal(t, 100, responses10Seconds, "Some 10 seconds responses MISSING")
}
