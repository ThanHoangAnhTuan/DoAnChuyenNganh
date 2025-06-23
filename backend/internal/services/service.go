package services

import (
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/accommodation"
	accommodationDetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/accommodation_detail"
	accommodationRoom "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/accommodation_room"
	adminLogin "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/admin/login"
	adminManager "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/admin/manager"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/facility"
	facilityDetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/facility_detail"
	managerLogin "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/manager/login"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/order"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/payment"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/review"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/stats"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/upload"
	userInfo "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/user/info"
	userLogin "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/user/login"
)

var (
	Accommodation       = accommodation.Use
	AccommodationDetail = accommodationDetail.Use
	AccommodationRoom   = accommodationRoom.Use
	AdminLogin          = adminLogin.Use
	AdminManager        = adminManager.Use
	Facility            = facility.Use
	FacilityDetail      = facilityDetail.Use
	ManagerLogin        = managerLogin.Use
	Order               = order.Use
	Payment             = payment.Use
	Review              = review.Use
	Stats               = stats.Use
	Upload              = upload.Use
	UserLogin           = userLogin.Use
	UserInfo            = userInfo.Use
)
