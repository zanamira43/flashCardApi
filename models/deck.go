package models

import (
	"time"
)

type Deck struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Bucket         int       `json:"bucket"`
	LastReviewDate time.Time `json:"last_review_date"`
	NextReviewDate time.Time `json:"next_review_date"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoCreateTime"`
	Card           []Card    `json:"cards"`
}
