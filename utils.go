package main

import (
	"archive/tar"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ulikunitz/xz"
)

// HTTP Handler Functions

// homeHandler serves the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>BuildConfig Git LFS Testing</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .endpoint { background: #f4f4f4; padding: 10px; margin: 10px 0; border-radius: 5px; }
        .method { color: #007bff; font-weight: bold; }
    </style>
</head>
<body>
    <h1>BuildConfig Git LFS Testing Web Server</h1>
    <p>Welcome to the web server version of the BuildConfig testing application!</p>
    
    <h2>Available Endpoints:</h2>
    <div class="endpoint">
        <span class="method">GET</span> <a href="/api/time">/api/time</a> - Current time
    </div>
    <div class="endpoint">
        <span class="method">GET</span> <a href="/api/health">/api/health</a> - Health check
    </div>
    <div class="endpoint">
        <span class="method">GET</span> <a href="/api/unzip">/api/unzip</a> - Unzip and list crc-linux-amd64.tar.xz contents
    </div>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

// timeHandler returns the current time
func timeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"time":     time.Now().Format("2006-01-02 15:04:05"),
		"timezone": time.Now().Format("MST"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// healthHandler returns the health status
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"service":   "BuildConfig Git LFS Testing",
		"uptime":    "running",
		"timestamp": time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// unzipHandler extracts and lists contents of crc-linux-amd64.tar.xz
func unzipHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"file":      "crc-linux-amd64.tar.xz",
		"status":    "processing",
		"timestamp": time.Now().Unix(),
		"contents":  []string{},
		"error":     "",
	}

	// Check if the file exists
	if _, err := os.Stat("crc-linux-amd64.tar.xz"); os.IsNotExist(err) {
		response["error"] = "File crc-linux-amd64.tar.xz not found"
		response["status"] = "error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Open the .tar.xz file
	file, err := os.Open("crc-linux-amd64.tar.xz")
	if err != nil {
		response["error"] = fmt.Sprintf("Failed to open file: %v", err)
		response["status"] = "error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()

	// Create xz reader
	xzReader, err := xz.NewReader(file)
	if err != nil {
		response["error"] = fmt.Sprintf("Failed to create xz reader: %v", err)
		response["status"] = "error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create tar reader
	tarReader := tar.NewReader(xzReader)

	var contents []string
	var totalSize int64
	var fileCount int

	// Read through the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			response["error"] = fmt.Sprintf("Failed to read tar header: %v", err)
			response["status"] = "error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		// Add file information to contents
		fileInfo := fmt.Sprintf("%s (Size: %d bytes, Type: %c)",
			header.Name,
			header.Size,
			header.Typeflag)
		contents = append(contents, fileInfo)

		totalSize += header.Size
		fileCount++

		// Skip file content (we only want the listing)
		if header.Typeflag == tar.TypeReg {
			_, err = io.Copy(io.Discard, tarReader)
			if err != nil {
				response["error"] = fmt.Sprintf("Failed to skip file content: %v", err)
				response["status"] = "error"
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
				return
			}
		}
	}

	// Update response with successful results
	response["status"] = "success"
	response["contents"] = contents
	response["file_count"] = fileCount
	response["total_size"] = totalSize
	response["archive_type"] = "tar.xz"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
