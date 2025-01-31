// server.go

package server

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func RegisterHandlers() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handlers.ServeTemplate)
	http.HandleFunc("/ascii-art", handlers.HandleAsciiArt)
}

func StartServer() {
	RegisterHandlers()
	
	log.Print("Listening on :http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
