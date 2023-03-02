package user

import (
	"log"
	"strings"
	"net/http"

	"go-rest-api/src/constant"
	"go-rest-api/src/pkg/jwt"
	entity "go-rest-api/src/http"
	service "go-rest-api/src/service/v1"
	"github.com/forkyid/go-utils/v1/aes"
	"github.com/forkyid/go-utils/v1/rest"
	"github.com/forkyid/go-utils/v1/validation"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

// Get user godoc
// @Summary Get User
// @Description Get id, username, fullname, and email
// @Tags Users
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param user_ids query string false "user_ids separated by comma"
// @Success 200 {object} http.GetUser
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/users [get]
func (ctrl *Controller) Get(ctx *gin.Context) {
	err := jwt.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		rest.ResponseMessage(ctx, http.StatusUnauthorized)
		return
	}

	// example query : a5V7Wz49nlPgbKEqVJkMQLpm3dyO6Zq2,3ZmaNeMwQ2X4oAE1RJ6ObLvkrVPxB9G5 (without spacing)
	log.Print(aes.Encrypt(1))
	log.Print(aes.Encrypt(9))
	userIDs := ctx.Query("user_ids")
	if userIDs != "" {
		userIDs, err := aes.DecryptBulk(strings.Split(userIDs, ","))
		if err != nil {
			rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
				"user_ids": constant.ErrInvalidID.Error()})
			return
		}
		response, err := ctrl.svc.Find(userIDs)
		if err != nil {
			rest.ResponseMessage(ctx, http.StatusInternalServerError)
			log.Println("get user by id:", err)
			return
		}
		rest.ResponseData(ctx, http.StatusOK, response)
		return
	}
}

// Register godoc
// @Summary Register User
// @Description Register User
// @Tags Users
// @Param Payload body http.RegisterUser true "Payload"
// @Success 201 {object} string "Created"
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "Resource Conflict"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/users/register [post]
func (ctrl *Controller) Register(ctx *gin.Context) {
	req := entity.RegisterUser{}
	if err := rest.BindJSON(ctx, &req); err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
			"body": constant.ErrInvalidFormat.Error()})
		return
	}

	req.Username = strings.ToLower(req.Username)
	err := ctrl.svc.Create(req)
	if errors.Is(err, constant.ErrUserExist) {
		rest.ResponseMessage(ctx, http.StatusConflict, errors.Cause(err).Error())
	} else if err != nil {
		log.Println("register:", err.Error())
		rest.ResponseMessage(ctx, http.StatusInternalServerError)
	} else {
		rest.ResponseMessage(ctx, http.StatusCreated)
	}
}

// Register godoc
// @Summary Register User
// @Description Register User
// @Tags Users
// @Param Payload body http.UpdateUser true "Payload"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/users [patch]
func (ctrl *Controller) Update(ctx *gin.Context) {
	request := entity.UpdateUser{}
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

	err = jwt.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		rest.ResponseMessage(ctx, http.StatusUnauthorized)
		return
	}

	userID := ctx.Query("user_ids")
	if userID == "" {
		rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
			"users": constant.ErrUserNotRegistered.Error()})
		return
	} else {
		err = ctrl.svc.Update(aes.Decrypt(userID), request)
		if err != nil {
			if errors.Is(err, constant.ErrUserNotRegistered) {
				rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
					"users": constant.ErrUserNotRegistered.Error()})
				return
			}
			rest.ResponseMessage(ctx, http.StatusInternalServerError)
			log.Println("update user: ", err.Error())
			return
		}

		rest.ResponseMessage(ctx, http.StatusOK)
	}
}

// Delete godoc
// @Summary Delete User
// @Description Delete User
// @Tags Users
// @Param Authorization header string true "Bearer Token"
// @Param user_ids query string false "user_ids separated by comma"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "Resource Conflict"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/users [delete]
func (ctrl *Controller) Delete(ctx *gin.Context) {
	err := jwt.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		rest.ResponseMessage(ctx, http.StatusUnauthorized)
		return
	}

	userID := ctx.Query("user_ids")
	if userID == "" {
		rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
			"users": constant.ErrUserNotRegistered.Error()})
		return
	} else {
		err := ctrl.svc.Delete(aes.Decrypt(userID))
		if err != nil {
			if errors.Is(err, constant.ErrUserNotRegistered) {
				rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
					"users": constant.ErrUserNotRegistered.Error()})
				return
			}
			rest.ResponseMessage(ctx, http.StatusInternalServerError)
			log.Println("delete user: ", err.Error())
			return
		}
		rest.ResponseMessage(ctx, http.StatusOK)
	}
}
