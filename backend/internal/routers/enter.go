package routers

import (
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/accommodation"
	accommodationDetail "github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/accommodation_detail"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/admin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/facility"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/order"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/payment"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/upload"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User                user.UserRouterGroup
	Admin               admin.AdminRouterGroup
	Manager             manager.ManagerRouterGroup
	Accommodation       accommodation.AccommodationRouterGroup
	AccommodationDetail accommodationDetail.AccommodationDetailRouterGroup
	Upload              upload.UploadRouterGroup
	Order               order.OrderRouterGroup
	Facility            facility.FacilityRouterGroup
	Payment             payment.PaymentRouterGroup
}

var RouterGroupApp = new(RouterGroup)
