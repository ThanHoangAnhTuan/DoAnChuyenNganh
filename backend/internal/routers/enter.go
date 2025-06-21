package routers

import (
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/accommodation"
	accommodationDetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/accommodation_detail"
	accommodationRoom "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/accommodation_room"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/admin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/facility"
	facilityDetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/facility_detail"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/manager"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/order"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/payment"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/review"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers/stats"
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
	FacilityDetail      facilityDetail.FacilityDetailRouterGroup
	Payment             payment.PaymentRouterGroup
	Review              review.ReviewRouterGroup
	Stats               stats.StatsRouterGroup
	Room                accommodationRoom.AccommodationRoomRouterGroup
}

var RouterGroupApp = new(RouterGroup)
