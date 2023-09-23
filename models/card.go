package models

type Card struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Word        string `json:"word"`
	Example     string `json:"example"`
	Description string `json:"description"`
	DeckId      int    `json:"deck_id"`
	Deck        Deck   `json:"-" gorm:"foreignKey:DeckId"`
}
