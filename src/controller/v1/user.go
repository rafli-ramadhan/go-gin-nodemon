package user

import (
	"log"
	"strings"
	"net/http"

	"go-rest-api/src/constant"
	pkg "go-rest-api/src/pkg/http"
	service "go-rest-api/src/service/v1"
	"github.com/forkyid/go-utils/v1/aes"
	"github.com/forkyid/go-utils/v1/jwt"
	"github.com/forkyid/go-utils/v1/rest"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Controller struct {
	svc service.Servicer
}

func NewController(
	svc service.Servicer,
) *Controller {
	return &Controller{
		svc: svc,
	}
}

// Get user godoc
// @Summary Get User
// @Description Get id, username, fullname, and email
// @Tags Users
// @Produce application/json
// @Param Authorization header string true "Bearer Token"
// @Param user_ids query string false "user_ids separated by comma"
// @Success 200 {object} user.GetResponseSchema
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/users [get]
func (ctrl *Controller) Get(ctx *gin.Context) {
	userID, err := jwt.ExtractClient(ctx.GetHeader("Authorization"))
	if err != nil || aes.Decrypt(userID.ID) < 0 {
		rest.ResponseMessage(ctx, http.StatusUnauthorized)
		return
	}

	users := ctx.Query("user_ids")
	if users != "" {
		result := []pkg.GetResponseSchema{}
		userIDs, err := aes.DecryptBulk(strings.Split(users, ","))
		if err != nil {
			rest.ResponseError(ctx, http.StatusBadRequest, map[string]string{
				"user_ids": constant.ErrInvalidID.Error()})
			return
		}

		usersData, err := ctrl.svc.Find(userIDs)
		if err != nil {
			rest.ResponseMessage(ctx, http.StatusInternalServerError)
			log.Println("get user by id:", err)
			return
		}

		for i := range usersData {
			result = append(result, usersData[i])
		}
		rest.ResponseData(ctx, http.StatusOK, result)
		return
	}
}

// Register godoc
// @Summary Register User
// @Description Register User
// @Tags Users
// @Param Payload body user.RegisterRequestSchema true "Payload"
// @Success 201 {object} string "Created"
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "Resource Conflict"
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/users/register [post]
func (ctrl *Controller) Register(ctx *gin.Context) {
	req := pkg.RegisterRequestSchema{}
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
		log.Println("register by email:", err.Error())
		rest.ResponseMessage(ctx, http.StatusInternalServerError)
	} else {
		rest.ResponseMessage(ctx, http.StatusCreated)
	}
}
