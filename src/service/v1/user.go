package user

import (
	models "go-rest-api/src/models/v1"
	pkg "go-rest-api/src/pkg/http"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Service struct {
	repo models.Repositorier
}

func NewService(
	repository models.Repositorier,
) *Service {
	return &Service{
		repo: repository,
	}
}

type Servicer interface {
	CheckEmailExist(email string) (exist bool, err error)
	Create(request pkg.RegisterRequestSchema) (err error)
}

func (svc *Service) CheckEmailExist(email string) (exist bool, err error) {
	exist = false
	_, err = svc.repo.Exist(email)
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
		err = errors.Wrap(err, "create")
		return
	}
	
	if !exist {
		err = svc.repo.Create(request)
		if err != nil {
			err = errors.Wrap(err, "create")
			return
		}
	}
	return
}

/*
func (svc *Service) Delete(ctx *gin.Context, request DeleteRequest) (err error) {
	host, err := svc.host.TakeID(request.MemberID)
	if err != nil {
		err = errors.Wrap(err, "take host")
		return
	}

	if request.userID == host.userID {
		err = constant.ErrInvalidID
		return
	}

	userID := aes.Decrypt(request.userID)
	userData, err := svc.repo.TakeIDByHostID(userID, aes.Decrypt(host.HostID))
	if err != nil {
		err = errors.Wrap(err, "db: take admin by user id and host id")
		return
	}

	user := dbuser.user{
		userID: userID,
		HostID:    aes.Decrypt(host.HostID),
	}
	err = svc.repo.Delete(user)
	if err != nil {
		err = errors.Wrap(err, "db: delete user")
		return
	}

	users, err := svc.user.GetuserByIDs(ctx, []int{userID})
	if err != nil {
		err = errors.Wrap(err, "get user by ids")
		return
	}
	return
}
*/