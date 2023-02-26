package user

import (
	"github.com/forkyid/go-utils/v1/aes"
	"go-rest-api/src/connection"
	"go-rest-api/src/constant"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	pkg "go-rest-api/src/pkg/http"
	"time"
)

type User struct {
	gorm.Model
	Username          string     `gorm:"column:username;type:varchar(50)"`
	FullName          *string    `gorm:"column:full_name;type:varchar(150)"`
	Email             *string    `gorm:"column:email;type:varchar(150)"`
	Password          *string    `gorm:"column:password;type:varchar(64)"`
	Country           *string    `gorm:"column:country;type:varchar(50)"`
	PhoneNumber       *string    `gorm:"column:phone_number;type:varchar(20)"`
	Description       *string    `gorm:"column:desctipytion;type:varchar(80)"`
	Gender            string     `gorm:"column:gender"`
	DateOfBirth       time.Time  `gorm:"column:date_of_birth;type:date"`
	SignupMethod      string     `gorm:"column:signup_method;type:varchar(5)"`
	StatusBanned      string     `gorm:"column:status_banned"`
	IsVerified        bool       `gorm:"column:is_verified;type:bool"`
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

type DB struct {
	Master *gorm.DB
}

type Repository struct {
	dbMaster *gorm.DB
}

func NewRepository(
	db connection.DB,
) *Repository {
	return &Repository{
		dbMaster: db.Master,
	}
}

type Repositorier interface {
	Find(userIDs []int) (users []pkg.GetResponseSchema, err error)
	Exist(email string) (exist bool, err error)
	Create(request pkg.RegisterRequestSchema) (err error)
	Delete(id string) (err error)
}

func (repo *Repository) Find(userIDs []int) (users []pkg.GetResponseSchema, err error) {
	query := repo.dbMaster.Model(&users).
		Select("id", "username", "nickname", "email", "used_storage", "status").
		Find(&users, userIDs)
	err = query.Error
	return
}

func (repo *Repository) Exist(email string) (exist bool, err error) {
	user := &User{}
	query := repo.dbMaster.Model(user).
		Where("email IN ?", user).
		Take(&user)
	err = query.Error
	return
}

func (repo *Repository) Create(request pkg.RegisterRequestSchema) (err error) {
	user := &User{}
	query := repo.dbMaster.Model(user).Begin().
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "account_id"}, {Name: "account_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"deleted_at": nil})}).
		Create(user)
	err = query.Error
	if err != nil {
		query.Rollback()
		return
	}

	err = query.Commit().Error
	return
}

func (repo *Repository) Delete(id string) (err error) {
	user := &User{}
	query := repo.dbMaster.Model(user).Begin().
		Where("id IN ?", id).
		Delete(user)
	err = query.Error
	if err != nil {
		query.Rollback()
		return
	}
	if query.RowsAffected != 1 {
		query.Rollback()
		err = constant.ErrInvalidID
		return
	}

	err = query.Commit().Error
	return
}
