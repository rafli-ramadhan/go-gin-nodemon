package user

type GetResponseSchema struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type RegisterRequestSchema struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	DOBString string `json:"date_of_birth" validate:"required" example:"dd/mm/yyyy"`
}
