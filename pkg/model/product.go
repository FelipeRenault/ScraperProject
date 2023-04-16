package model

import "time"

type Product struct {
	ID          int       `json:"id" db:"id"`
	URL         string    `json:"url" db:"url"`
	Title       string    `json:"title" db:"title"`
	Image       string    `json:"image" db:"image"`
	Price       int       `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

var maxTimeThreshold = time.Hour

func (p Product) IsWithinTimeThreshold() bool {
	oldestAllowed := time.Now().Add(-maxTimeThreshold)
	return p.UpdatedAt.After(oldestAllowed)
}
