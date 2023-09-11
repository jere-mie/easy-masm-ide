package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	// Serve the HTML page with a button only at the root path ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "static/index.html")
		}
	})

	// Handle the "/execute" route for executing commands
	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Execute the "echo hello && echo world" command
			cmd := exec.Command("sh", "-c", "echo hello && echo world")
			output, err := cmd.CombinedOutput()
			if err != nil {
				http.Error(w, fmt.Sprintf("Error executing command: %v", err), http.StatusInternalServerError)
				return
			}
			// Send the command output as the response
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write(output)
		}
	})

	// Serve static files from the "static" directory under the "/static/" URL path
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
