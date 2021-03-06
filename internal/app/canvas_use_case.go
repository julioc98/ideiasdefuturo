// Package app use cases
package app

import (
	"github.com/julioc98/ideiasdefuturo/internal/domain"
)

type canvasStorager interface {
	Store(user *domain.Canvas) (*domain.Canvas, error)
	FindOne(query *domain.Canvas, args ...string) (*domain.Canvas, error)
	Remove(query *domain.Canvas, userID string) error
	Find(query *domain.Canvas, args ...string) ([]domain.Canvas, error)
	Update(query *domain.Canvas) error
}

// CanvasUseCase Canvas auth uses case.
type CanvasUseCase struct {
	repository canvasStorager
	validate   checker
}

// NewCanvasUseCase factory.
func NewCanvasUseCase(s canvasStorager, v checker) *CanvasUseCase {
	return &CanvasUseCase{
		repository: s,
		validate:   v,
	}
}

// Create a new Canvas.
func (u *CanvasUseCase) Create(canvas *domain.Canvas) (*domain.Canvas, error) {
	if err := u.validate.Struct(canvas); err != nil {
		return nil, ErrInvalid
	}

	newCanvas, err := u.repository.Store(canvas)
	if err != nil {
		return nil, ErrOnSave
	}

	return newCanvas, nil
}

// GetByUserID Get canvas By UserID.
func (u *CanvasUseCase) GetByUserID(userID string) ([]domain.Canvas, error) {
	e := &domain.Canvas{
		UserID: userID,
	}

	canvas, err := u.repository.Find(e, "user_id")
	if err != nil {
		return nil, ErrInvalid
	}

	return canvas, nil
}

// Get canvas.
func (u *CanvasUseCase) Get(id uint, userID string) (*domain.Canvas, error) {
	e := &domain.Canvas{
		ID:     id,
		UserID: userID,
	}

	canvas, err := u.repository.FindOne(e, "id", "user_id")
	if err != nil {
		return nil, ErrInvalid
	}

	return canvas, nil
}

// Delete canvas.
func (u *CanvasUseCase) Delete(id uint, userID string) error {
	e := &domain.Canvas{
		ID:     id,
		UserID: userID,
	}

	err := u.repository.Remove(e, userID)
	if err != nil {
		return ErrInvalid
	}

	return nil
}

// Update Canvas.
func (u *CanvasUseCase) Update(canvas *domain.Canvas) error {
	if err := u.validate.Struct(canvas); err != nil {
		return ErrInvalid
	}

	err := u.repository.Update(canvas)
	if err != nil {
		return ErrOnSave
	}

	return nil
}
