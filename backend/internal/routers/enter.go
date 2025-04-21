package routers

import (
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/admin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Admin   admin.AdminRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
