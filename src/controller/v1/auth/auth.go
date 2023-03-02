package auth

import (
	"fmt"
	"log"
	"net/http"

	"go-rest-api/src/constant"
	entity "go-rest-api/src/http"
	service "go-rest-api/src/service/v1"
	"go-rest-api/src/pkg/jwt"
	"github.com/forkyid/go-utils/v1/rest"
	"github.com/forkyid/go-utils/v1/validation"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	svc service.Servicer
}

func NewController(
	servicer service.Servicer,
) *Controller {
	return &Controller{
		svc: servicer,
	}
}

// @Summary Get Auth
// @Description Get Auth
// @Tags Auth
// @Produce application/json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} user.Auth
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/auth [get]
func (ctrl *Controller) Login(ctx *gin.Context) {
	request := entity.Auth{}
	err := rest.BindJSON(ctx, &request)
	if err != nil {
		log.Println("bind json:", err, "request:", request)
		rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
			"body": constant.ErrInvalidFormat.Error()})
		return
	}

	if err := validation.Validator.Struct(request); err != nil {
		log.Println("validate struct:", err, "request:", request)
		rest.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}

	exist, _ := ctrl.svc.CheckEmailExist(request.Email)
	if !exist {
		rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
			"users": constant.ErrUserNotRegistered.Error()})
		return
	}

	user, err:= ctrl.svc.TakeUserByEmail(request.Email)
	if err != nil {
		rest.ResponseMessage(ctx, http.StatusInternalServerError)
		return
	}

	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		rest.ResponseMessage(ctx, http.StatusInternalServerError)
		return
	}

	rest.ResponseData(ctx, http.StatusOK, map[string]string{
		"token": fmt.Sprintf("Bearer %v", token),
	})
}
