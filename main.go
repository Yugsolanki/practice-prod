package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	app_name := os.Getenv("APP_NAME")
	fmt.Println("Starting application:", app_name)

	// Basic Web Server serving html and images from /images folder
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "./index.htm")
		log.Printf("Served index.htm to %s by %s\n", r.RemoteAddr, app_name)
	})

	fs := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
