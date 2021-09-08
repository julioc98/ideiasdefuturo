// Package handler manage flow
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julioc98/ideiasdefuturo/internal/domain"
	"github.com/julioc98/ideiasdefuturo/internal/infra/gateway"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type canvasServicer interface {
	Create(user *domain.Canvas) (*domain.Canvas, error)
	GetByUserID(userID string) ([]domain.Canvas, error)
	Get(id uint, userID string) (*domain.Canvas, error)
	Delete(id uint, userID string) error
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

// FindOne endpoint.
func (uh *CanvasRestHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	canvas, err := uh.usecase.Get(uint(id), uh.guard.GetUserID(r))
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

// Find endpoint.
func (uh *CanvasRestHandler) Find(w http.ResponseWriter, r *http.Request) {
	canvas, err := uh.usecase.GetByUserID(uh.guard.GetUserID(r))
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

// Delete endpoint.
func (uh *CanvasRestHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = uh.usecase.Delete(uint(id), uh.guard.GetUserID(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// SetCanvasRoutes mux configuration.
func (uh *CanvasRestHandler) SetCanvasRoutes(r *mux.Router, n negroni.Negroni) {
	r.Handle("", n.With(
		negroni.WrapFunc(uh.Create),
	)).Methods(http.MethodPost, http.MethodOptions).Name("create")
	r.Handle("/{id:[0-9]+}", n.With(
		negroni.WrapFunc(uh.FindOne),
	)).Methods(http.MethodGet, http.MethodOptions).Name("findOne")
	r.Handle("", n.With(
		negroni.WrapFunc(uh.Find),
	)).Methods(http.MethodGet, http.MethodOptions).Name("find")
	r.Handle("/{id:[0-9]+}", n.With(
		negroni.WrapFunc(uh.Delete),
	)).Methods(http.MethodDelete, http.MethodOptions).Name("delete")
}
