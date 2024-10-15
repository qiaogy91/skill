package impl_test

import (
	_ "case04/apps"
	"case04/apps/user"
	"context"
	"github.com/qiaogy91/ioc"
	"testing"
)

var (
	ctx = context.Background()
	c   = user.GetSvc()
)

func init() {
	if err := ioc.ConfigIocObject("/Users/qiaogy/GolandProjects/projects/skill/case04/etc/application.yaml"); err != nil {
		panic(err)
	}
}

func TestImpl_CreateTable(t *testing.T) {

	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_GetUser(t *testing.T) {
	u, err := c.GetUser(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", u)
}
