package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

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

func main() {
	http.HandleFunc("/weight-scale/api", getWeight)
	fmt.Println("Servidor iniciado en http://localhost:9999")
	error := http.ListenAndServe(":9999", nil)
	fmt.Println("error: ", error)
}

func getWeight(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/home/m/peso.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

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
}
