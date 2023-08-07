package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/stream", streamHandler)
	fmt.Println("Servidor de streaming iniciado en http://localhost:8080/stream")
	http.ListenAndServe(":8080", nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Envía datos en un bucle cada segundo
	for {
		// Cambia el mensaje y el evento como desees
		message := "Mensaje de streaming: " + time.Now().Format("2006-01-02 15:04:05")
		event := "actualizacion"

		fmt.Fprintf(w, "event: %s\n", event)
		fmt.Fprintf(w, "data: %s\n\n", message)

		w.(http.Flusher).Flush()

		time.Sleep(1 * time.Second)
	}
}
