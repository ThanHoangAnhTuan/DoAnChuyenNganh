package routers

import (
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/accommodation"
	accommodationDetail "github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/accommodation_detail"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/admin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/image"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User                user.UserRouterGroup
	Admin               admin.AdminRouterGroup
	Manager             manager.ManagerRouterGroup
	Accommodation       accommodation.AccommodationRouterGroup
	AccommodationDetail accommodationDetail.AccommodationDetailRouterGroup
	Image               image.ImageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
