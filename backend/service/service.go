package service

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionService interface {
	SaveUserID(ctx *gin.Context, userID int) error
	GetUserID(ctx *gin.Context) (int, error)
	DeleteUserID(ctx *gin.Context) error
}

type sessionService struct{}

func NewSessionService() SessionService {
	return &sessionService{}
}

func (s *sessionService) SaveUserID(ctx *gin.Context, userID int) error {
	session := sessions.Default(ctx)
	session.Set("user_id", userID)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (s *sessionService) GetUserID(ctx *gin.Context) (int, error) {
	session := sessions.Default(ctx)
	id := session.Get("user_id")
	if id == nil {
		return 0, errors.New("user_id not found in session")
	}

	userID, ok := id.(int)
	if !ok {
		return 0, errors.New("invalid user_id type")
	}

	return userID, nil
}

func (s *sessionService) DeleteUserID(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Delete("user_id")
	return session.Save()
}
