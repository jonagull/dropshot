package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// MaxUploadSize limits the size of uploaded files (in bytes)
const MaxUploadSize = 10 * 1024 * 1024 // 10MB

// UploadDirectory is where the images will be stored
const UploadDirectory = "./uploads"

func main() {
	// Create the upload directory if it doesn't exist
	if _, err := os.Stat(UploadDirectory); os.IsNotExist(err) {
		os.Mkdir(UploadDirectory, 0755)
	}

	// Set up routes and start the server
	http.HandleFunc("/upload", uploadFileHandler)
	http.HandleFunc("/uploads", listUploadsHandler)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(UploadDirectory))))

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// uploadFileHandler handles image uploads
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*") // Change "*" to specific domain for restricted access
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Parse the multipart form with a size limit
	r.ParseMultipartForm(MaxUploadSize)

	// Get the uploaded file
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Check if the file is an image by extension (optional)
	ext := filepath.Ext(handler.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		http.Error(w, "Only JPEG and PNG images are allowed", http.StatusBadRequest)
		return
	}

	// Create a new file in the upload directory
	dst, err := os.Create(filepath.Join(UploadDirectory, handler.Filename))
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Respond to the client with success
	w.Write([]byte(fmt.Sprintf("File uploaded successfully: %s", handler.Filename)))
}

// listUploadsHandler handles listing all uploaded images
func listUploadsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	files, err := os.ReadDir(UploadDirectory)
	if err != nil {
		http.Error(w, "Unable to read uploads directory", http.StatusInternalServerError)
		return
	}

	// Create a slice to store the filenames
	fileNames := make([]string, 0)

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	// Structure the JSON response with "images" as the key
	response := map[string]interface{}{
		"images": fileNames,
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Convert the map to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all domains, or specify your frontend URL

	// Write the JSON response
	w.Write(jsonData)
}
