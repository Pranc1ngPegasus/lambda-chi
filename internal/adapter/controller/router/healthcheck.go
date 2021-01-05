package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var _ Healthcheck = (http.Handler)(nil)

type(
	Healthcheck http.Handler
)

func NewHealthcheck() Healthcheck {
	r := chi.NewRouter()

	r.Get("/", healthcheck)

	return r
}

type HealthcheckResponse struct {
	Message string `json:"message"`
}

func healthcheck(w http.ResponseWriter, req *http.Request) {
	render.Status(req, http.StatusOK)
	render.JSON(w, req, HealthcheckResponse{
		Message: "ok",
	})
}
