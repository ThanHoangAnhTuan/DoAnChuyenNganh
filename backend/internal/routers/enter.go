package routers

import (
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/accommodation"
	accommodationDetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/accommodation_detail"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/admin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/facility"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/manager"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/order"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/payment"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/upload"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/user"
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
