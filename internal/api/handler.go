package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ProofPayload struct {
	MeterID   string `json:"meter_id"`
	Timestamp int64  `json:"timestamp"`
	Proof     []byte `json:"proof"`
}

func HandleProof(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[DEBUG] Incoming request: Method=%s, URL=%s\n", r.Method, r.URL.Path)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload ProofPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("[ERROR] Failed to decode JSON from request: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("[WARNING] Failed to close request body: %v\n", err)
		}
	}()

	fmt.Printf("[API] Received ZKP proof from %s | Size: %d bytes | Time: %d\n",
		payload.MeterID, len(payload.Proof), payload.Timestamp)

	w.WriteHeader(http.StatusOK)
}
