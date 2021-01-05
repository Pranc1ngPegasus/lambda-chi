// +build wireinject

package main

import (
	"net/http"

	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/configuration"
	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/controller/router"
	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/controller/server"
	"github.com/google/wire"
)

func initialize() *http.Server {
	wire.Build(
		configuration.Get,
		router.NewRouter,
		server.NewServer,
	)

	return nil
}
