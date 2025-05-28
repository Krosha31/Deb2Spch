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
	"archive/zip"
	"strings"
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
	req.FileName = "/app/uploads/" + req.FileName
	fmt.Println(req)
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


type DownloadRequest struct {
    SeparatedPaths   []string `json:"separated_paths"`
    User             string   `json:"user"`
    OriginalFilename string   `json:"original_filename"`
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    var req DownloadRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Not decoded" + err.Error())
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
	fmt.Println("DownloadHandler second")
    if len(req.SeparatedPaths) == 0 {
        http.Error(w, "No files to download", http.StatusBadRequest)
        return
    }
	fmt.Println("Файлы До", req.SeparatedPaths)
    archivePath, err := createZipArchive(req.SeparatedPaths, req.OriginalFilename)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to create archive: %v", err), http.StatusInternalServerError)
        return
    }
    defer os.Remove(archivePath) 
    file, err := os.Open(archivePath)
    if err != nil {
        http.Error(w, "Failed to open archive", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    fileInfo, err := file.Stat()
    if err != nil {
        http.Error(w, "Failed to get file info", http.StatusInternalServerError)
        return
    }

    downloadName := getDownloadFilename(req.OriginalFilename)

    w.Header().Set("Content-Type", "application/zip")
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", downloadName))
    w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

    _, err = io.Copy(w, file)
    if err != nil {
        fmt.Printf("Error sending file: %v\n", err)
    }
}

func createZipArchive(filePaths []string, originalFilename string) (string, error) {
	fmt.Println("Файлы для архивации:", filePaths)
    tempFile, err := os.CreateTemp("", "separated_audio_*.zip")
    if err != nil {
        return "", fmt.Errorf("failed to create temp file: %v", err)
    }
    defer tempFile.Close()

    zipWriter := zip.NewWriter(tempFile)
    defer zipWriter.Close()

    for _, filePath := range filePaths {
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            fmt.Printf("Warning: file not found: %s\n", filePath)
            continue
        }

        file, err := os.Open(filePath)
        if err != nil {
            fmt.Printf("Warning: failed to open file %s: %v\n", filePath, err)
            continue
        }

        fileName := filepath.Base(filePath)
        
        zipFile, err := zipWriter.Create(fileName)
        if err != nil {
            file.Close()
            return "", fmt.Errorf("failed to create file in archive: %v", err)
        }

        _, err = io.Copy(zipFile, file)
        file.Close()
        
        if err != nil {
            return "", fmt.Errorf("failed to copy file to archive: %v", err)
        }
    }

    return tempFile.Name(), nil
}

func getDownloadFilename(originalFilename string) string {
    if originalFilename == "" {
        return "separated_audio.zip"
    }

    nameWithoutExt := strings.TrimSuffix(originalFilename, filepath.Ext(originalFilename))
    return fmt.Sprintf("%s_separated.zip", nameWithoutExt)
}


