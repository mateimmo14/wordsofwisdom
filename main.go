package main

import (
	"log"
	"net/http"
)

func main() {
	servmux := http.NewServeMux()
	servmux.HandleFunc("GET /", handleRoot)
	servmux.HandleFunc("POST /", handleSubmit)
	err = http.ListenAndServe(":8080", servmux)
	if err != nil {
		log.Println(err)
		return
	}
}
