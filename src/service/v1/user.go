package user

import (
	// "log"
	"github.com/forkyid/go-utils/v1/aes"
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
	Take(userID int) (users http.GetUser, err error)
	TakeUserByEmail(email string) (user dbentity.User, err error)
	Find(userIDs []int) (users []http.GetUser, err error)
	CheckEmailExist(email string) (exist bool, err error)
	Create(request http.RegisterUser) (err error)
	Update(userID int, request http.UpdateUser) (err error)
	Delete(userID int) (err error)
}

func (svc *Service) Take(userID int) (user http.GetUser, err error) {
	takeUser, err := svc.model.Take(userID)
	if err == gorm.ErrRecordNotFound {
		err = constant.ErrUserNotRegistered
		return
	} else if err != nil {
		err = errors.Wrap(err, "take user")
		return
	}

	user = http.GetUser{}
	copier.Copy(&user, &takeUser)
	user.ID = aes.Encrypt(int(takeUser.ID))
	return
}

func (svc *Service) TakeUserByEmail(email string) (users dbentity.User, err error) {
	return svc.model.TakeUserByEmail(email)
}

func (svc *Service) Find(userIDs []int) (users []http.GetUser, err error) {
	findUsers, err := svc.model.Find(userIDs)
	if err != nil {
		err = errors.Wrap(err, "find users")
		return
	}
	for i := range findUsers {
		user := http.GetUser{}
		copier.Copy(&user, &findUsers[i])
		user.ID = aes.Encrypt(int(findUsers[i].ID))
		users = append(users, user)
	} 
	return
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

func (svc *Service) Create(request http.RegisterUser) (err error) {
	exist, err := svc.CheckEmailExist(request.Email)
	if err != nil {
		return
	}
	
	if !exist {
		newUser := dbentity.User{}
		copier.Copy(&newUser, &request)
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
	if err != nil {
		err = errors.Wrap(err, "delete user: user is not exist")
		return
	}

	err = svc.model.Delete(userID)
	if err != nil {
		err = errors.Wrap(err, "delete user")
		return
	}
	return
}
