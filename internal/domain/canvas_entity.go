// Package domain entities
package domain

import "time"

// Canvas ....
type Canvas struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title" validate:"required"`
	UserID       string    `json:"userId" validate:"required"`
	Description  string    `json:"description"`
	Description2 string    `json:"description2"`
	Description3 string    `json:"description3"`
	Description4 string    `json:"description4"`
	Description5 string    `json:"description5"`
	Description6 string    `json:"description6"`
	Description7 string    `json:"description7"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
}
