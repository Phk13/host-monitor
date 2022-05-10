package utils

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func OAuthRedirectionServer(codeChannel chan string) {
	log.Debug("Starting OAuth callback server...")
	handler := http.NewServeMux()
	server := http.Server{Addr: ":9090", Handler: handler}
	handler.HandleFunc("/callback-gl", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		codeChannel <- code
		fmt.Fprintf(w, "OAuth configured. You can close this window")
	})
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Debug("Stopping OAuth callback server")
}
