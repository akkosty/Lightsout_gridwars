package bot

import (
	"fmt"
	"net/http"
)

// RunServer starts a simple HTTP server for health checks
func RunServer(port string) error {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	return http.ListenAndServe(":"+port, nil)
}
