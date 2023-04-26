package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	fh "pacpayz-full-hosting/form_handling"
)

var serverPort = 8080

func main() {
	// handle form submission
	http.HandleFunc("/basic-submit", fh.HandleBasicForm)
	http.HandleFunc("/select-payment-submit", fh.HandleSelectPaymentForm)

	http.HandleFunc("/card-submit", fh.HandleCardForm)
	http.HandleFunc("/crypto-submit", fh.HandleCryptoForm)
	http.HandleFunc("/other-submit", fh.HandleOtherForm)

	http.HandleFunc("/processing", fh.HandleProcessingForm)

	// serve static assets
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// serve HTML files
	http.Handle("/", http.FileServer(http.Dir("web")))

	// start server
	log.Println("Server listening on port ", strconv.Itoa(serverPort))

	// handle SIGINT (CTRL+C) signal
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigint
		log.Println("Stopping server...")
		os.Exit(0)
	}()

	err := http.ListenAndServe(":"+strconv.Itoa(serverPort), nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
