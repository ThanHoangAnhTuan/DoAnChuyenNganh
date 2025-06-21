package controllers

import (
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/accommodation"
	accommodationDetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/accommodation_detail"
	accommodationRoom "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/accommodation_room"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/admin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/facility"
	facilityDetail "github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/facility_detail"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/manager"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/order"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/payment"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/review"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/stats"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/upload"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers/user"
)

var AdminManager = admin.AdminManager
var AdminLogin = admin.AdminLogin
var AccommodationDetail = accommodationDetail.Handler
var AccommodationRoom = accommodationRoom.Handler
var Accommodation = accommodation.Handler
var Facility = facility.Handler
var FacilityDetail = facilityDetail.Handler
var UserLogin = user.UserLogin
var UserInfo = user.UserInfo
var ManagerLogin = manager.Handler
var Review = review.Handler
var Order = order.Handler
var Payment = payment.Handler
var Stats = stats.Handler
var UploadImage = upload.Handler
