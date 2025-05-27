package upload

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
	"encoding/json"
	"time"
	"bytes"
)



func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving the file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    username := r.FormValue("user")
    fileName := r.FormValue("name")
    fmt.Printf("File %s  upload by user: %s\n", fileName,  username)

    outFile, err := os.Create(filepath.Join("/app/uploads", fileName))
    if err != nil {
        http.Error(w, "Error creating the file", http.StatusInternalServerError)
        return
    }
    defer outFile.Close()

    if _, err := io.Copy(outFile, file); err != nil {
        http.Error(w, "Error saving the file", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("File uploaded succesfuly"))
}

type SplitRequest struct {
	FileName string `json:"path"`
}

type SplitResponse struct {
	Segments []string `json:"separated_paths"`
}

func SplitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SplitRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	req.FileName = "/app/uploads" + req.FileName

	bodyBytes, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to serialize request", http.StatusInternalServerError)
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	resp, err := client.Post("http://mossformer:5000/separate", "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		http.Error(w, "Model service unreachable", http.StatusBadGateway)
		fmt.Println("POST to MossFormer failed:", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Model service error: " + string(respBody), http.StatusBadGateway)
		return
	}

	var modelResp SplitResponse
	if err := json.Unmarshal(respBody, &modelResp); err != nil {
		http.Error(w, "Failed to parse model response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modelResp)
}
