// main.go
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
    // Serve static files from the "static" directory
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Handle the root path to serve the index.html
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })

    // Existing handlers
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/time", timeHandler)

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}