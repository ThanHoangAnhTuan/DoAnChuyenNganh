package initializer

import (
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/accommodation"
	accommodationDetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/accommodation_detail"
	accommodationroom "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/accommodation_room"
	adminLogin "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/admin/login"
	adminManager "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/admin/manager"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/facility"
	facilitydetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/facility_detail"
	managerLogin "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/manager/login"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/order"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/payment"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/review"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/stats"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/upload"
	userInfo "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/user/info"
	userLogin "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/user/login"
)

func InitInterface() {
	db := global.Mysql
	queries := database.New(db)

	userLogin.Init(userLogin.New(queries))
	userInfo.Init(userInfo.New(queries))
	accommodation.Init(accommodation.New(queries))
	accommodationDetail.Init(accommodationDetail.New(queries))
	accommodationroom.Init(accommodationroom.New(queries))
	upload.Init(upload.New(queries))
	order.Init(order.New(queries, db))
	facility.Init(facility.New(queries))
	facilitydetail.Init(facilitydetail.New(queries))
	review.Init(review.New(queries))
	stats.Init(stats.New(queries))
	payment.Init(payment.New(queries, db))
	managerLogin.Init(managerLogin.New(queries))
	adminLogin.Init(adminLogin.New(queries))
	adminManager.Init(adminManager.New(queries))
}
