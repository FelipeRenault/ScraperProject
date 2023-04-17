package model

import "time"

type Product struct {
	ID          int       `db:"id"              json:"-"`
	URL         string    `db:"url"             json:"url"`
	Title       string    `db:"title"           json:"title"`
	Image       string    `db:"image"           json:"image"`
	Price       int       `db:"price"           json:"price"`
	Description string    `db:"description"     json:"description"`
	CreatedAt   time.Time `db:"created_at"      json:"-"`
	UpdatedAt   time.Time `db:"updated_at"      json:"-"`
}

var maxTimeThreshold = time.Minute

func (p Product) IsWithinTimeThreshold() bool {
	oldestAllowed := time.Now().Add(-maxTimeThreshold)
	return p.UpdatedAt.After(oldestAllowed)
}
