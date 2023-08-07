package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const fileName = "peso.txt"
const endPoint = "/weight-scale/api"

func main() {
	if errEnv := os.Setenv("PORT", "9999"); errEnv != nil {
		log.Fatalf("ENV: %s", errEnv)
	}

	http.HandleFunc(endPoint, getWeight)
	port := os.Getenv("PORT")
	fmt.Printf("Servidor iniciado \nhttp://localhost:%s%s", port, endPoint)
	createError := http.ListenAndServe(":"+port, nil)
	if createError != nil {
		log.Fatalf("Listen serve error: %s", createError)
	}

}

func getWeight(w http.ResponseWriter, _ *http.Request) {
	filePath := os.Getenv("HOME") + "/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		if s, err := strconv.ParseFloat(line, 64); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			weightScale := WeightScale{
				Date:   getDate(),
				Weight: s,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(weightScale); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	} else if err := scanner.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("CLOSE FILE: %s", err)
		}
	}(file)
}

func getDate() string {
	t := time.Now()
	date := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return date
}

type WeightScale struct {
	Date   string  `json:"Date"`
	Weight float64 `json:"weight"`
}
