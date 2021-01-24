package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tamnk74/todolist-mysql-go/models"
	Email "github.com/tamnk74/todolist-mysql-go/utils/email"
)

//login contorller interface
type AuthController interface {
	login(ctx *gin.Context)
	register(ctx *gin.Context)
}

type authController struct {
	authService AuthService
}

func NewAuthController() AuthController {
	authService := NewAuthService()
	return &authController{
		authService: authService,
	}
}

func (a *authController) login(c *gin.Context) {
	var form Login
	if err := c.ShouldBindJSON(&form); err != nil {
		c.Error(err)
		return
	}
	token, err := a.authService.Login(form.Email, form.Password)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{
		"data": gin.H{
			"access_token": token,
			"token_type":   "bearer",
		},
	})
}

func (a *authController) register(c *gin.Context) {
	var form Register
	if err := c.ShouldBindJSON(&form); err != nil {
		c.Error(err)
		return
	}

	user := models.User{Name: form.Name, Email: form.Email, Password: form.Password}
	newItem, err := a.authService.CreateUser(user)
	if err != nil {
		c.Error(err)
		return
	}
	Email.Send([]string{"khac.tam.94@gmail.com"})
	c.JSON(200, gin.H{"data": newItem})
}
