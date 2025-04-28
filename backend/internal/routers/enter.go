package routers

import (
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/accommodation"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/admin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User          user.UserRouterGroup
	Admin         admin.AdminRouterGroup
	Manager       manager.ManagerRouterGroup
	Accommodation accommodation.AccommodationRouterGroup
}

var RouterGroupApp = new(RouterGroup)
