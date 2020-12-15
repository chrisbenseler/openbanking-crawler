package adapters

import (
	"net/http"
	"openbankingcrawler/services"

	"github.com/gin-gonic/gin"
)

//AuthenticateController authenticate controller
type AuthenticateController interface {
	SignIn(c *gin.Context)
}

type authenticateController struct {
	service services.Auth
}

//NewAuthenticateController create a new controller for authentication
func NewAuthenticateController(service services.Auth) AuthenticateController {
	return &authenticateController{
		service: service,
	}
}

//SignIn sign in controller
func (ctrl *authenticateController) SignIn(c *gin.Context) {

	var payload AuthenticatePayload

	c.BindJSON(&payload)

	token, err := ctrl.service.CreateAccessToken(payload.Email, payload.Password)

	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
