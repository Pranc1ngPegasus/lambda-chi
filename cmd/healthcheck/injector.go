// +build wireinject

package main

import(
	"github.com/google/wire"
	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/lambda/handler"
	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/configuration"
	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/controller/router"
)

func initialize() handler.Healthcheck {
	wire.Build(
		configuration.Get,
		router.NewRouter,
		handler.NewHealthcheck,
	)

	return nil
}
