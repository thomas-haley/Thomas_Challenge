package main

import (
	"log"
	"net/http"
)

func redirectFunc(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, r.Host+r.RequestURI, http.StatusMovedPermanently)
}

func tlsRedirect() {
	log.Println("Starting redirect on :80")
	err := http.ListenAndServe(":80", http.HandlerFunc(redirectFunc))
	if err != nil {
		log.Fatal("Redirect server failed:", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	go tlsRedirect()
	// Start the server on port 443
	log.Println("Starting server on :443")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
