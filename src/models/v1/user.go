package models

import (
	"go-rest-api/src/connection"
	"go-rest-api/src/constant"
	"go-rest-api/src/dbentity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"go-rest-api/src/http"
)
 
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
	Take(userID int) (user dbentity.User, err error)
	Find(userIDs []int) (users []dbentity.User, err error)
	TakeUserByEmail(email string) (user dbentity.User, err error)
	TakeUserByUsername(username string) (user dbentity.User, err error)
	Create(user dbentity.User) (err error)
	Update(userID int, request http.UpdateUser) (err error)
	Delete(userID int) (err error)
}

func (repo *Repository) Take(userID int) (user dbentity.User, err error) {
	query := repo.dbMaster.Model(&dbentity.User{}).
		Select("id", "username", "nickname", "email", "used_storage", "status").
		Take(&user, userID)
	err = query.Error
	return
}

func (repo *Repository) Find(userIDs []int) (users []dbentity.User, err error) {
	query := repo.dbMaster.Model(&dbentity.User{}).
		Select("id", "username", "nickname", "email", "used_storage", "status").
		Find(&users, userIDs)
	err = query.Error
	return
}

func (repo *Repository) TakeUserByEmail(email string) (user dbentity.User, err error) {
	query := repo.dbMaster.Model(&dbentity.User{}).
		Select("email").
		Where("email", email).
		Take(&user)
	err = query.Error
	return
}

func (repo *Repository) TakeUserByUsername(username string) (user dbentity.User, err error) {
	query := repo.dbMaster.Model(&dbentity.User{}).
		Select("username").
		Where("username", username).
		Take(&user)
	err = query.Error
	return
}

func (repo *Repository) Create(user dbentity.User) (err error) {
	query := repo.dbMaster.Model(&user).Begin().
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "email"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"deleted_at": nil})}).
		Create(&user)
	err = query.Error
	if err != nil {
		query.Rollback()
		return
	}

	err = query.Commit().Error
	return
}

func (repo *Repository) Update(userID int, request http.UpdateUser) (err error) {
	user := &dbentity.User{}
	query := repo.dbMaster.Model(&user).Begin().
		Where("id IN ?", userID).
		Updates(request)
	err = query.Error
	if err != nil {
		query.Rollback()
		return
	}

	err = query.Commit().Error
	return
}


func (repo *Repository) Delete(userID int) (err error) {
	user := &dbentity.User{}
	query := repo.dbMaster.Model(user).Begin().
		Where("id IN ?", userID).
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
