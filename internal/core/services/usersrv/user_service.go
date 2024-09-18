package usersrv

import (
	"context"
	"departments-organigram/internal/core/domain"
	"departments-organigram/internal/core/ports"
	"departments-organigram/internal/init/config"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type usersSrv struct {
	usersStore ports.UsersStore
	conf       config.Config
}

func NewUsersSrv(usersStore ports.UsersStore, conf config.Config) *usersSrv {
	return &usersSrv{
		usersStore: usersStore,
		conf:       conf,
	}
}

func (s *usersSrv) Register(ctx context.Context, username, password string) error {
	passwordHashed, err := hashPassword(password)
	if err != nil {
		return err
	}
	return s.usersStore.CreateUser(context.Background(), domain.User{
		Username: username,
		Password: passwordHashed,
	})
}

func (s *usersSrv) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.usersStore.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if !isPasswordValid(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.conf.JWTSecret))
}

func hashPassword(plainTextPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func isPasswordValid(password, paswordStore string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(paswordStore), []byte(password))

	return err == nil
}
