package user

import (
	"log"
	"strings"
	"net/http"

	"go-rest-api/src/constant"
	pkg "go-rest-api/src/pkg/http"
	service "go-rest-api/src/service/v1"
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

// Register godoc
// @Summary Register User
// @Description Register User
// @Tags Register
// @Param Payload body register.EmailRequest true "Payload"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad Request"
// @Failure 409 {object} string "Resource Conflict"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/auth/register [post]
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
		log.Println("[ERROR] register by email:", err.Error())
		rest.ResponseMessage(ctx, http.StatusInternalServerError)
	} else {
		rest.ResponseMessage(ctx, http.StatusCreated)
	}

	rest.ResponseMessage(ctx, http.StatusOK)
}
