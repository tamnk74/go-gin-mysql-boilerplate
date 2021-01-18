package auth

import (
	"errors"
	"fmt"

	"github.com/tamnk74/todolist-mysql-go/models"
	"github.com/tamnk74/todolist-mysql-go/repository"
	"github.com/tamnk74/todolist-mysql-go/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email string, password string) (token string, err error)
	CreateUser(user models.User) (usr models.User, err error)
}

type Auth struct {
	authRepo repository.AuthRepository
}

func NewAuthService() AuthService {
	authRepo := repository.NewAuthRepository()
	return &Auth{
		authRepo: authRepo,
	}

}

func (a *Auth) Login(email string, password string) (token string, err error) {
	user, err := a.authRepo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("Invalid email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(user.Password, password, err)
		return "", errors.New("Invalid password")
	}

	token = utils.GenerateAccessToken(user)

	return token, nil
}

func (a *Auth) CreateUser(user models.User) (usr models.User, err error) {
	return a.authRepo.Create(user)
}
