package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func main() {
	port := os.Getenv("EASY_MASM_IDE_PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if EASY_MASM_IDE_PORT is not set
	}

	// Serve the HTML page with a button at the root path ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Printf("Request received from IP %s for /", r.RemoteAddr)
			http.ServeFile(w, r, "static/index.html")
		}
	})

	// Handle the "/execute" route for executing commands
	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			log.Printf("Request received from IP %s for /execute", r.RemoteAddr)

			// Parse the JSON request body to extract the "code" variable
			var requestBody map[string]string
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&requestBody); err != nil {
				http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
				return
			}
			code, found := requestBody["code"]
			if !found {
				http.Error(w, "JSON request missing 'code' field", http.StatusBadRequest)
				return
			}

			// Generate a random hex code for the filename
			randomHex := strconv.FormatInt(rand.Int63(), 16)
			fileName := randomHex + ".asm"
			filePath := filepath.Join("easy-masm/src", fileName)

			// Ensure that the directory exists
			if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				log.Printf("Error creating directory: %v", err)
				http.Error(w, fmt.Sprintf("Error creating directory: %v", err), http.StatusInternalServerError)
				return
			}

			// Save the "code" variable to the generated filename
			file, err := os.Create(filePath)
			if err != nil {
				log.Printf("Error creating %s: %v", filePath, err)
				http.Error(w, fmt.Sprintf("Error creating %s: %v", filePath, err), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			_, err = file.WriteString(code)
			if err != nil {
				log.Printf("Error writing code to %s: %v", filePath, err)
				http.Error(w, fmt.Sprintf("Error writing code to %s: %v", filePath, err), http.StatusInternalServerError)
				return
			}

			// Execute the "echo hello && echo world" command
			cmd := exec.Command("sh", "-c", "ls -lah /")
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

			// Delete the file after command execution
			if err := os.Remove(filePath); err != nil {
				log.Printf("Error deleting %s: %v", filePath, err)
			}
			log.Printf("Successfully deleted %s", filePath)
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
