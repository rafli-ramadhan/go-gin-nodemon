package dbentity

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/forkyid/go-utils/v1/aes"
	"go-rest-api/src/constant"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username          string    `gorm:"column:username;type:varchar(50)"`
	FullName          *string   `gorm:"column:full_name;type:varchar(150)"`
	Email             string    `gorm:"column:email;type:varchar(150)"`
	Password          string    `gorm:"column:password;type:varchar(64)"`
	Country           *string   `gorm:"column:country;type:varchar(50)"`
	PhoneNumber       *string   `gorm:"column:phone_number;type:varchar(20)"`
	Description       *string   `gorm:"column:description;type:varchar(80)"`
	Gender            string    `gorm:"column:gender"`
	DateOfBirth       time.Time `gorm:"column:date_of_birth;type:date"`
	SignupMethod      string   `gorm:"column:signup_method;type:varchar(5)"`
	IsVerified        bool      `gorm:"column:is_verified;type:bool"`
}

func (User) TableName() string {
	return "users"
}

func (m User) DOBString() string {
	return m.DateOfBirth.Format(constant.DOBLayout)
}

func (m User) EncID() string {
	return aes.Encrypt(int(m.ID))
}

func (m *User) GeneratePasswordHarsh() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), 14)
	m.Password = string(bytes)
	return err
}
func (m *User) CheckPasswordHarsh(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}