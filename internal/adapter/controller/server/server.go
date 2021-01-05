package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/configuration"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func NewServer(handler http.Handler, config configuration.Config) *http.Server {
	port := config.ListenPort
	log.Printf("Listen at :%d", port)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	return server
}
