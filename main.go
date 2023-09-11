package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	port := os.Getenv("EASY_MASM_IDE_PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if EASY_MASM_IDE_PORT is not set
	}

	// Serve the HTML page with a button at the root path ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			log.Printf("Request received from IP %s for /", r.RemoteAddr)
			http.ServeFile(w, r, "static/index.html")
		}
	})

	// Handle the "/execute" route for executing commands
	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			log.Printf("Request received from IP %s for /execute", r.RemoteAddr)

			// Execute the "echo hello && echo world" command
			cmd := exec.Command("sh", "-c", "echo hello && echo world")
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Error executing command: %v", err)
				http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
				return
			}
			// Send the command output as the response
			log.Printf("Command executed successfully")
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		}
	})

	// Serve static files from the "static" directory under the "/static/" URL path
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	listenAddr := ":" + port
	log.Printf("Server is listening on %s", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
