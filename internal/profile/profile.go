package profile

import (
	"net/http"
	"encoding/json"
	// "Deb2Spch/internal/database"
)

type HistoryRequest struct {
    UserID string `json:"user_id"`
}


func  HistoryHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req HistoryRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // requests, err := database.Db.GetRequestsByUser(req.UserID)
    // if err != nil {
    //     http.Error(w, "Server error", http.StatusInternalServerError)
    //     return
    // }

    // w.Header().Set("Content-Type", "application/json")
    // json.NewEncoder(w).Encode(requests)
}