// Package repository save data
package repository

import (
	"github.com/julioc98/ideiasdefuturo/internal/domain"

	"gorm.io/gorm"
)

// CanvasGorm repository.
type CanvasGorm struct {
	db *gorm.DB
}

// NeWCanvasGorm repository factory.
func NeWCanvasGorm(db *gorm.DB) *CanvasGorm {
	return &CanvasGorm{
		db: db,
	}
}

// Store an canvas.
func (g *CanvasGorm) Store(canvas *domain.Canvas) (*domain.Canvas, error) {
	if dbc := g.db.Create(canvas); dbc.Error != nil {
		return nil, dbc.Error
	}

	return canvas, nil
}

// FindOne canvas.
func (g *CanvasGorm) FindOne(query *domain.Canvas, args ...string) (*domain.Canvas, error) {
	e := &domain.Canvas{}
	if dbc := g.db.Where(query, args).First(e); dbc.Error != nil {
		return nil, dbc.Error
	}

	return e, nil
}

// Find many canvas.
func (g *CanvasGorm) Find(query *domain.Canvas, args ...string) ([]domain.Canvas, error) {
	e := []domain.Canvas{}
	if dbc := g.db.Where(query, args).Find(&e); dbc.Error != nil {
		return nil, dbc.Error
	}

	return e, nil
}
