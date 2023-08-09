package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/stream", streamHandler)
	fmt.Println("Servidor de streaming iniciado en http://localhost:8080/stream")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Listen and serve ERR:", err)
	}

}

func streamHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// encoder
	encoder := json.NewEncoder(w)

	// Envía datos en un bucle cada segundo
	for {
		data := Data{
			Message: "streaming!",
			Time:    time.Now().Format(time.RFC3339),
			Ton:     rand.Intn(100),
		}

		if err := encoder.Encode(data); err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		time.Sleep(1 * time.Second)
	}
}

type Data struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	Ton     int    `json:"ton"`
}
