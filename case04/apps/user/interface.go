package user

import (
	"context"
	"github.com/qiaogy91/ioc"
)

const AppName = "user"

func GetSvc() Service { return ioc.Controller().Get(AppName).(Service) }

type Service interface {
	CreateTable(ctx context.Context) error
	GetUser(ctx context.Context) (*User, error)
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
