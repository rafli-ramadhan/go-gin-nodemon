package http

type Auth struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type Token struct {
	Token string `json:"token"`
}

type GetUser struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type RegisterUser struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Gender    string `json:"gender"`
	Country   string `gorm:"column:country;type:varchar(50)"`
}

type UpdateUser struct {
	Username    string  `json:"username" validate:"required"`
	Email       string  `json:"email" validate:"required"`
	Password    string  `json:"password" validate:"required"`
	Country     string `gorm:"column:country;type:varchar(50)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20)"`
	Description string `gorm:"column:desctipytion;type:varchar(80)"`
	Gender      string  `gorm:"column:gender"`
	DOBString   string  `json:"date_of_birth" validate:"required" example:"dd/mm/yyyy"`
}
