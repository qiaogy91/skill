package impl

import (
	"case04/apps/user"
	"context"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&user.User{})
}

func (i *Impl) GetUser(ctx context.Context) (*user.User, error) {
	u := &user.User{}
	if err := i.db.WithContext(ctx).Model(&user.User{}).First(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
