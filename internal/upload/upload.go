package upload

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"
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

    outFile, err := os.Create(filepath.Join("uploads", fileName))
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

func SplitHandler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(15 * time.Second)
    w.WriteHeader(http.StatusOK)
}
