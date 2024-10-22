package impl

import (
	"case04/apps/user"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/log"
	"gorm.io/gorm"
	"log/slog"
)

type Impl struct {
	ioc.ObjectImpl
	log *slog.Logger
	db  *gorm.DB
}

func (i *Impl) Name() string  { return user.AppName }
func (i *Impl) Priority() int { return 301 }

func (i *Impl) Init() {
	i.db = datasource.DB()
	i.log = log.Sub(user.AppName)
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
