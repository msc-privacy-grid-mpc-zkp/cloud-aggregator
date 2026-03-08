package api

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(port string) {
	http.HandleFunc("/api/proofs", HandleProof)

	fmt.Printf("[SERVER] Listening on http://localhost%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("[FATAL] Server crashed: %v", err)
	}
}
