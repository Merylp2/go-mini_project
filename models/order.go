package models

type Order struct {
	ID       uint  `gorm:"primary_key" json:"id"`
	UserID   uint  `json:"user_id"`
	MovieID  uint  `json:"movie_id"`
	Quantity int   `json:"quantity"`
	Movie    Movie `json:"-"`
}
