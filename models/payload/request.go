package payload

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" form:"name" validate:"required"`
	NoHP     string `json:"no_hp" form:"no_hp" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type OrderRequest struct {
	UserID   uint `json:"user_id" form:"user_id" validate:"required"`
	MovieID  uint `json:"movie_id" form:"movie_id" validate:"required"`
	Quantity uint `json:"quantity" form:"quantity" validate:"required"`
}

type TopupRequest struct {
	Saldo uint `json:"saldo" form:"saldo"`
}
