package user

import (
	"log"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-rest-api/src/constant"
	models "go-rest-api/src/models/v1"
	"go-rest-api/src/dbentity"
	"go-rest-api/src/http"
	"gorm.io/gorm"
)

type Service struct {
	model models.Repositorier
}

func NewService(
	repositorier models.Repositorier,
) *Service {
	return &Service{
		model: repositorier,
	}
}

type Servicer interface {
	Take(userID int) (users dbentity.User, err error)
	Find(userIDs []int) (users []dbentity.User, err error)
	CheckUsernameExist(email string) (exist bool, err error)
	CheckEmailExist(email string) (exist bool, err error)
	Create(request http.RegisterUser) (err error)
	Update(userID int, request http.UpdateUser) (err error)
	Delete(userID int) (err error)
}

func (svc *Service) Take(userID int) (users dbentity.User, err error) {
	return svc.model.Take(userID)
}

func (svc *Service) Find(userIDs []int) (users []dbentity.User, err error) {
	return svc.model.Find(userIDs)
}

func (svc *Service) CheckEmailExist(email string) (exist bool, err error) {
	exist = false
	_, err = svc.model.TakeUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			exist = false
			err = nil
			return
		} else {
			err = errors.Wrap(err, "check email exist")
		}
	}
	exist = true
	return
}

func (svc *Service) CheckUsernameExist(email string) (exist bool, err error) {
	exist = false
	_, err = svc.model.TakeUserByUsername(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			exist = true
			err = nil
			return
		}
		err = errors.Wrap(err, "check username exist")
	}
	return
}

func (svc *Service) Create(request http.RegisterUser) (err error) {
	exist, err := svc.CheckEmailExist(request.Email)
	if err != nil {
		return
	}
	log.Print(exist)
	log.Print(request)
	
	if !exist {
		newUser := dbentity.User{}
		copier.Copy(&newUser, &request)
		log.Print("newUser : ", newUser)
		newUser.IsVerified = false
		err = svc.model.Create(newUser)
		if err != nil {
			err = errors.Wrap(err, "create")
			return
		}
	}
	return
}

func (svc *Service) Update(userID int, request http.UpdateUser) (err error) {
	_, err = svc.Take(userID)
	if err == gorm.ErrRecordNotFound {
		err = constant.ErrUserNotRegistered
		return
	} else if err != nil {
		err = errors.Wrap(err, "user is not exist")
		return
	} else if err == nil {
		err = svc.model.Update(userID, request)
		if err != nil {
			err = errors.Wrap(err, "delete user")
			return
		}
	}
	return
}

func (svc *Service) Delete(userID int) (err error) {
	_, err = svc.Take(userID)
	if err == gorm.ErrRecordNotFound {
		err = constant.ErrUserNotRegistered
		return
	} else if err != nil {
		err = errors.Wrap(err, "user is not exist")
		return
	} else if err == nil {
		err = svc.model.Delete(userID)
		if err != nil {
			err = errors.Wrap(err, "delete user")
			return
		}
	}
	return
}
