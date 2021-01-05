package handler

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/configuration"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

var _ Healthcheck = (*healthcheck)(nil)

type (
	Healthcheck interface {
		Get(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	}

	healthcheck struct {
		config  configuration.Config
		router  http.Handler
		adapter *httpadapter.HandlerAdapter
	}
)

func NewHealthcheck(
	config configuration.Config,
	router http.Handler,
) Healthcheck {
	return &healthcheck{
		config: configuration.Get(),
		router: router,
		adapter: httpadapter.New(router),
	}
}

func (h *healthcheck) Get(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.adapter.ProxyWithContext(ctx, req)
}
