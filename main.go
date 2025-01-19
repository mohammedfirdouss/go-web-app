package main

import (
    "fmt"
    "net/http"
    "time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("15:04:05")
    fmt.Fprintf(w, "Current Server Time: %s", currentTime)
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/time", timeHandler)

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}