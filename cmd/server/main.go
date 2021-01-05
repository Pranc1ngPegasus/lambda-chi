package main

import (
	"fmt"
	"net/http"

	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/configuration"
)

func main() {
	configuration.Load()

	server := initialize()
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("server did not boot. err:%v", err)
		return
	}
}
