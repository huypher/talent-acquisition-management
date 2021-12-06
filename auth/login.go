package auth

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/pghuy/talent-acquistion-management/domain"
	"github.com/sirupsen/logrus"
)

type authUsecase struct {
	userUsecase domain.TalentUsecase
}

func NewAuthUsecase(userUsecase domain.TalentUsecase) *authUsecase {
	return &authUsecase{
		userUsecase: userUsecase,
	}
}

func (u *authUsecase) Login(ctx context.Context, username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("Invalid username or password")
	}

	user, err := u.userUsecase.GetByUserName(ctx, username)
	if err != nil {
		return "", err
	}
	if user == (domain.Talent{}) {
		return "", NewInvalidUserNameError("invalid username")
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passErr != nil {
		return "", NewWrongPasswordError("wrong password")
	}

	token, err := genJWT(&user, jwtTokenKey)
	if err != nil {
		logrus.WithError(err).Errorf("gen JWT failed")
		return "", errors.New("internal error")
	}

	return token, nil
}

func genJWT(acc *domain.Talent, secretJWT string) (string, error) {
	if acc == nil {
		return "", errors.New("talent empty")
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = acc.ID
	claims["username"] = acc.UserName
	claims["name"] = acc.Name
	claims["expired_time"] = time.Now().Add(time.Hour * 24).Unix()
	claims["token_type"] = "Bearer"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretJWT))
}

// use to hash talent password when talent create account
func genHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logrus.WithError(err).Errorf("Gen hash from pwd failed")
		return "", err
	}
	return string(hash), nil
}
