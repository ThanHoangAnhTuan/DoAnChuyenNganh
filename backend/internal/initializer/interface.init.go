package initializer

import (
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services/impl"
)

func InitInterface() {
	db := global.Mysql
	queries := database.New(db)
	services.InitUserLogin(impl.NewUserLoginImpl(queries))
	services.InitAccommodation(impl.NewAccommodationImpl(queries))
	services.InitAccommodationDetail(impl.NewAccommodationDetailImpl(queries))
	services.InitManagerLogin(impl.NewManagerLoginImpl(queries))
	services.InitAdminLogin(impl.NewAdminLoginImpl(queries))
	services.InitUpload(impl.NewUploadImpl(queries))
	services.InitOrder(impl.NewOrderImpl(queries, db))
	services.InitFacility(impl.NewFacilityImpl(queries))
	services.InitPayment(impl.NewPaymentImpl(queries, db))
}
