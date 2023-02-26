package user

import (
	"go-rest-api/src/constant"
	models "go-rest-api/src/models/v1"
	pkg "go-rest-api/src/pkg/http"
	"github.com/pkg/errors"
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
	Take(userID int) (users models.User, err error)
	Find(userIDs []int) (users []models.User, err error)
	CheckEmailExist(email string) (exist bool, err error)
	Create(request pkg.RegisterRequestSchema) (err error)
	Update(userID int, request pkg.UpdateRequestSchema) (err error)
	Delete(userID int) (err error)
}

func (svc *Service) Take(userID int) (users models.User, err error) {
	return svc.model.Take(userID)
}

func (svc *Service) Find(userIDs []int) (users []models.User, err error) {
	return svc.model.Find(userIDs)
}

func (svc *Service) CheckEmailExist(email string) (exist bool, err error) {
	exist = false
	_, err = svc.model.TakeByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			exist = true
			err = nil
			return
		}
		err = errors.Wrap(err, "check email exist")
	}
	return
}

func (svc *Service) Create(request pkg.RegisterRequestSchema) (err error) {
	exist, err := svc.CheckEmailExist(request.Email)
	if err != nil {
		return
	}
	
	if !exist {
		err = svc.model.Create(request)
		if err != nil {
			err = errors.Wrap(err, "create")
			return
		}
	}
	return
}

func (svc *Service) Update(userID int, request pkg.UpdateRequestSchema) (err error) {
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
