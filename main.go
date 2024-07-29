package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response структура для JSON ответа
type Response struct {
	Message string `json:"message"`
}


// handler функция для обработки запросов
func handler(w http.ResponseWriter, r *http.Request) {
	// Логируем информацию о запросе
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

	if r.Method == http.MethodOptions {
		// Обработка OPTIONS запросов
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Создаем и отправляем ответ для других типов запросов
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :81")
	err := http.ListenAndServe(":81", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}