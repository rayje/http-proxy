package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Endpoint struct {
	Name    string            `json:"name"`
	Status  int               `json:"status"`
	Delay   int               `json:"delay"`
	Headers map[string]string `json: "headers"`
	Body    string            `json:"body"`
}

type Config struct {
	Endpoints []Endpoint `json:"endpoints"`
}

func main() {
	configJson := flag.String("config", "", "The location fo the JSON config file")
	flag.Parse()

	file, err := ioutil.ReadFile(*configJson)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	endpoints := config.Endpoints
	for _, endpoint := range endpoints {
		http.HandleFunc(endpoint.Name, func(w http.ResponseWriter, r *http.Request) {
			for k, v := range endpoint.Headers {
				w.Header().Set(k, v)
			}
			w.WriteHeader(endpoint.Status)
			w.Write([]byte(endpoint.Body))
		})
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
