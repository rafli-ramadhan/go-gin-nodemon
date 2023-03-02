package http

type Auth struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type Token struct {
	Token string `json:"token"`
}

type GetUser struct {
	ID        string `json:"id" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
}

type RegisterUser struct {
	Username  string  `json:"username" validate:"required"`
	Email     string  `json:"email" validate:"required"`
	Password  string  `json:"password" validate:"required"`
	Gender    string  `json:"gender"`
	Country   *string `json:"country"`
}

type UpdateUser struct {
	Username    string  `json:"username" validate:"required"`
	Email       string  `json:"email" validate:"required"`
	Password    string  `json:"password" validate:"required"`
	Country     string  `json:"country"`
	PhoneNumber string  `json:"phone_number"`
	Description string  `json:"description"`
	Gender      string  `json:"gender"`
	DOBString   string  `json:"date_of_birth" validate:"required" example:"dd/mm/yyyy"`
}
