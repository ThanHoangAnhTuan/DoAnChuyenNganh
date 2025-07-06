package response

const (
	// global
	ErrCodeInternalServerError = 100001
	ErrCodeUnauthorized        = 100002
	ErrCodeValidator           = 100003
	ErrCodeForbidden           = 100004
	ErrCodeSuccessfully        = 100005

	// login
	ErrCodeLoginSuccess    = 200001
	ErrCodeLoginFailed     = 200002
	ErrCodeRegisterSuccess = 200003

	// otp
	ErrCodeOTPNotMatch        = 300001
	ErrCodeVerifyOTPSuccess   = 300002
	ErrCodeOTPAlreadyVerified = 300003
	ErrCodeOTPNotVerified     = 300004
	ErrCodeOTPAlreadyExists   = 300005

	// user manager
	ErrCodeManagerNotFound   = 400001
	ErrCodeGetManagerSuccess = 400002

	// user admin
	ErrCodeGetAdminFailed        = 500001
	ErrCodeAdminPasswordNotMatch = 500002
	ErrCodeAdminLoginSuccess     = 500003
	ErrCodeAccountAlreadyExists  = 500004

	// user base
	ErrCodeUserNotFound      = 600001
	ErrCodeGetUserSuccess    = 600002
	ErrCodeUpdateUserSuccess = 600003

	// accommodation
	ErrCodeCreateAccommodationSuccess = 700001
	ErrCodeGetAccommodationSuccess    = 700002
	ErrCodeAccommodationNotFound      = 700003
	ErrCodeUpdateAccommodationSuccess = 700004
	ErrCodeDeleteAccommodationSuccess = 700005

	// accommodation type
	ErrCodeCreateAccommodationTypeSuccess = 800001
	ErrCodeAccommodationTypeNotFound      = 800002
	ErrCodeDeleteAccommodationTypeSuccess = 800003
	ErrCodeGetAccommodationTypeSuccess    = 800004
	ErrCodeUpdateAccommodationTypeSuccess = 800005

	// accommodation room
	ErrCodeCreateAccommodationRoomSuccess = 900001
	ErrCodeGetAccommodationRoomSuccess    = 900002
	ErrCodeUpdateAccommodationRoomSuccess = 900003
	ErrCodeDeleteAccommodationRoomSuccess = 900004
	ErrCodeAccommodationRoomNotFound      = 900005

	// facility
	ErrCodeCreateFacilitySuccess = 1000001
	ErrCodeGetFacilitySuccess    = 1000002
	ErrCodeUpdateFacilitySuccess = 1000003
	ErrCodeDeleteFacilitySuccess = 1000004

	// facility detail
	ErrCodeCreateFacilityDetailSuccess = 1100001
	ErrCodeGetFacilityDetailSuccess    = 1100002
	ErrCodeDeleteFacilityDetailSuccess = 1100003
	ErrCodeUpdateFacilityDetailSuccess = 1100004

	// order
	ErrCodeOrderNotFound      = 1200001
	ErrCodeGetOrderSuccess    = 1200002
	ErrCodeUpdateOrderSuccess = 1200003

	// review
	ErrCodeCreateReviewSuccess      = 1300001
	ErrCodeUserNotBookAccommodation = 1300002
	ErrCodeGetReviewsSuccess        = 1300003

	// stats
	ErrCodeStatsSuccess = 1400001

	// export
	ErrCodeExportSuccess = 1500001

	// image
	ErrCodeGetFileSuccess    = 1600001
	ErrCodeUploadFileSuccess = 1600002
)

var message = map[int]string{
	// global
	ErrCodeInternalServerError: "Đã xảy ra lỗi hệ thống. Vui lòng thử lại sau.",
	ErrCodeValidator:           "Dữ liệu không hợp lệ. Vui lòng kiểm tra lại các trường thông tin.",
	ErrCodeUnauthorized:        "Bạn chưa đăng nhập. Vui lòng đăng nhập",
	ErrCodeForbidden:           "Bạn không có quyền truy cập tài nguyên này.",
	ErrCodeSuccessfully:        "Thành công",

	// login
	ErrCodeLoginSuccess:    "Đăng nhập thành công",
	ErrCodeRegisterSuccess: "Đăng ký thành công",
	ErrCodeLoginFailed:     "Tài khoản hoặc mật khẩu không đúng",

	// accommodation
	ErrCodeAccommodationNotFound:      "Không tìm thấy khách sạn",
	ErrCodeGetAccommodationSuccess:    "Tải dữ liệu khách sạn thành công",
	ErrCodeDeleteAccommodationSuccess: "Xoá khách sạn thành công",
	ErrCodeUpdateAccommodationSuccess: "Cập nhập thông tin khách sạn thành công",
	ErrCodeCreateAccommodationSuccess: "Tạo khách sạn thành công",

	// accommodation type
	ErrCodeCreateAccommodationTypeSuccess: "Tạo loại phòng khách sạn thành công",
	ErrCodeAccommodationTypeNotFound:      "Không tìm thấy loại phòng khách sạn",
	ErrCodeDeleteAccommodationTypeSuccess: "Xoá loại phòng khách sạn thành công",
	ErrCodeGetAccommodationTypeSuccess:    "Tải dữ liệu loại phòng khách sạn thành công",
	ErrCodeUpdateAccommodationTypeSuccess: "Cập nhập dữ liệu loại phòng khách sạn thành công",

	// accommodation room
	ErrCodeDeleteAccommodationRoomSuccess: "Xoá phòng thành công",
	ErrCodeUpdateAccommodationRoomSuccess: "Cập nhập thông tin phòng thành công",
	ErrCodeGetAccommodationRoomSuccess:    "Tải dữ liệu phòng khách sạn thành công",
	ErrCodeCreateAccommodationRoomSuccess: "Tạo phòng khách sạn thành công",
	ErrCodeAccommodationRoomNotFound:      "Không tìm thấy phòng khách sạn",

	// facility
	ErrCodeDeleteFacilitySuccess: "Xoá tiện nghi khách sạn thành công",
	ErrCodeUpdateFacilitySuccess: "Cập nhập tiện nghi khách sạn thành công",
	ErrCodeCreateFacilitySuccess: "Tạo tiện nghi khách sạn thành công",
	ErrCodeGetFacilitySuccess:    "Tải dữ liệu tiện nghi khách sạn thành công",

	// facility detail
	ErrCodeDeleteFacilityDetailSuccess: "Xoá tiện nghi phòng thành công",
	ErrCodeUpdateFacilityDetailSuccess: "Cập nhập tiện nghi phòng thành công",
	ErrCodeCreateFacilityDetailSuccess: "Tạo tiện nghi phòng thành công",
	ErrCodeGetFacilityDetailSuccess:    "Tải dữ liệu tiện nghi phòng thành công",

	// order
	ErrCodeOrderNotFound:      "Đơn hàng không tồn tại",
	ErrCodeUpdateOrderSuccess: "Cập nhật đơn hàng thành công",
	ErrCodeGetOrderSuccess:    "Tải thông tin đơn hàng thành công",

	// review
	ErrCodeCreateReviewSuccess:      "Bình luận thành công",
	ErrCodeUserNotBookAccommodation: "Bạn chưa đặt phòng, nên không có quyền bình luận",
	ErrCodeGetReviewsSuccess:        "Tải danh sách bình luận thành công",

	// stats
	ErrCodeStatsSuccess: "Thống kê thành công",

	// export
	ErrCodeExportSuccess: "Tải file thành công",

	// image
	ErrCodeGetFileSuccess:    "Lấy file thành công",
	ErrCodeUploadFileSuccess: "Tải file thành công",

	// user manager
	ErrCodeManagerNotFound:   "Bạn không có quyền truy cập tài nguyên này.",
	ErrCodeGetManagerSuccess: "Tải danh sách quản lý thành công",

	// user admin
	ErrCodeGetAdminFailed:        "Đã xảy ra lỗi trong quá trình đăng nhập. Vui lòng thử lại sau",
	ErrCodeAdminPasswordNotMatch: "Tải khoản hoặc mật khẩu không đúng",
	ErrCodeAdminLoginSuccess:     "Đăng nhập thành công",
	ErrCodeAccountAlreadyExists:  "Tài khoản đã tồn tại. Vui lòng chọn email khác.",

	// user base
	ErrCodeUserNotFound:      "Người dùng không tồn tại",
	ErrCodeGetUserSuccess:    "Lấy thông tin người dùng thành công",
	ErrCodeUpdateUserSuccess: "Cập nhật thông tin người dùng thành công",

	// otp
	ErrCodeOTPAlreadyVerified: "OTP không hợp lệ",
	ErrCodeOTPNotMatch:        "OTP không đúng",
	ErrCodeVerifyOTPSuccess:   "Xác thực OTP thành công",
	ErrCodeOTPNotVerified:     "Bạn chưa xác thực OTP",
	ErrCodeOTPAlreadyExists:   "OTP đã tồn tại",
}
