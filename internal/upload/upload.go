package upload

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
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

    username := r.FormValue("username")
    fmt.Printf("File upload by user: %s\n", username)

    outFile, err := os.Create(filepath.Join("uploads", "uploaded_file")) // Замените на желаемое имя файла
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