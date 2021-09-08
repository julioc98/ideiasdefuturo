// Package handler manage flow
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julioc98/ideiasdefuturo/internal/domain"
	"github.com/julioc98/ideiasdefuturo/internal/infra/gateway"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type canvasServicer interface {
	Create(user *domain.Canvas) (*domain.Canvas, error)
}

// CanvasRestHandler http handler.
type CanvasRestHandler struct {
	usecase canvasServicer
	guard   *gateway.Guardian
}

// NewCanvasRestHandler factory.
func NewCanvasRestHandler(uc canvasServicer, guard *gateway.Guardian) *CanvasRestHandler {
	return &CanvasRestHandler{
		usecase: uc,
		guard:   guard,
	}
}

// Create endpoint.
func (uh *CanvasRestHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := &domain.Canvas{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	req.UserID = uh.guard.GetUserID(r)

	canvas, err := uh.usecase.Create(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	res, err := json.Marshal(canvas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

// SetCanvasRoutes mux configuration.
func (uh *CanvasRestHandler) SetCanvasRoutes(r *mux.Router, n negroni.Negroni) {
	r.Handle("", n.With(
		negroni.WrapFunc(uh.Create),
	)).Methods(http.MethodPost, http.MethodOptions).Name("create")
}
