package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Define the endpoint where your server will listen for push notifications
	http.HandleFunc("/push-endpoint", handlePushNotification)

	// Specify the port for your server to listen on
	port := 8080

	// Start the HTTP server
	log.Printf("Server listening on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handlePushNotification(w http.ResponseWriter, r *http.Request) {
	// Read the incoming request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Process the received push notification (add your own logic here)
	log.Printf("Received push notification:\n%s\n", body)

	// Respond to the Pub/Sub acknowledgment
	fmt.Fprint(w, "OK")
}
