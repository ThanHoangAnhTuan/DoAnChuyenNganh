package response

const (
	// global
	ErrCodeSaveDataFailed         = 20001
	ErrCodeUpdateDataFailed       = 20002
	ErrCodeMarshalFailed          = 20003
	ErrCodeUnMarshalFailed        = 20004
	ErrCodeCreateJWTTokenFailed   = 20005
	ErrCodeParamsInvalid          = 20006
	ErrCodeUnauthorized           = 20007
	ErrCodeForbidden              = 20008
	ErrCodeValidatorNotFound      = 20009
	ErrCodeValidator              = 20010
	ErrCodeConvertISOToUnixFailed = 20011
	ErrCodeConvertUnixToISOFailed = 20012
	ErrCodeSuccessfully           = 20013
	ErrCodeParseTimeFailed        = 20014

	// register
	ErrCodeInvalidEmailFormat = 30001
	ErrCodeUserAlreadyExists  = 30002
	ErrCodeOTPAlreadyExists   = 30003
	ErrCodeSendEmailFailed    = 30004
	ErrCodeRegisterSuccess    = 30005
	ErrCodeRegisterFailed     = 30006

	// verify otp
	ErrCodeOTPNotExists           = 40001
	ErrCodeOTPNotMatch            = 40002
	ErrCodeGetInfoOTPFailed       = 40003
	ErrCodeUpdateUserVerifyFailed = 40004
	ErrCodeVerifyOTPSuccess       = 40005
	ErrCodeOTPAlreadyVerified     = 40006

	// update password register
	ErrCodeOTPNotVerified                = 50001
	ErrCodeHashPasswordFailed            = 50002
	ErrCodeUpdatePasswordRegisterSuccess = 50003

	// login
	ErrCodeGetUserInfoFailed     = 60001
	ErrCodePasswordNotMatch      = 60002
	ErrCodeLoginSuccess          = 60003
	ErrCodeGetUserInfoSuccess    = 60004
	ErrCodeGetUserInfoNotFound   = 60005
	ErrCodeUpdateUserInfoFailed  = 60006
	ErrCodeUpdateUserInfoSuccess = 60007

	// jwt
	ErrCodeMissAuthorizationHeader    = 70001
	ErrCodeInvalidAuthorizationFormat = 70002
	ErrCodeInvalidToken               = 70003

	// accommodation
	ErrCodeCreateAccommodationFailed   = 80001
	ErrCodeCreateAccommodationSuccess  = 80002
	ErrCodeManagerNotFound             = 80003
	ErrCodeGetAccommodationsFailed     = 80004
	ErrCodeGetAccommodationSuccess     = 80005
	ErrCodeGetAccommodationFailed      = 80006
	ErrCodeAccommodationNotFound       = 80007
	ErrCodeUpdateAccommodationSuccess  = 80008
	ErrCodeUpdateAccommodationFailed   = 80009
	ErrCodeDeleteAccommodationFailed   = 80010
	ErrCodeDeleteAccommodationSuccess  = 80011
	ErrCodeGetCountAccommodationFailed = 80012

	// accommodation image
	ErrCodeGetAccommodationImagesFailed    = 80013
	ErrCodeGetAccommodationImagesSuccess   = 80014
	ErrCodeDeleteAccommodationImagesFailed = 80015
	ErrCodeSaveAccommodationImagesFailed   = 80016

	// accommodation detail
	ErrCodeCreateAccommodationDetailFailed   = 90001
	ErrCodeCreateAccommodationDetailSuccess  = 90002
	ErrCodeGetAccommodationDetailsFailed     = 90004
	ErrCodeGetAccommodationDetailsSuccess    = 90005
	ErrCodeGetAccommodationDetailFailed      = 90006
	ErrCodeAccommodationDetailNotFound       = 90007
	ErrCodeUpdateAccommodationDetailSuccess  = 90008
	ErrCodeUpdateAccommodationDetailFailed   = 90009
	ErrCodeDeleteAccommodationDetailFailed   = 90010
	ErrCodeDeleteAccommodationDetailSuccess  = 90011
	ErrCodeGetCountAccommodationDetailFailed = 90012
	ErrCodeNumberOfAvailableRoomsNotEnough   = 90013

	// accommodation detail image
	ErrCodeGetAccommodationDetailImagesFailed    = 90014
	ErrCodeGetAccommodationDetailImagesSuccess   = 90015
	ErrCodeDeleteAccommodationDetailImagesFailed = 90016
	ErrCodeSaveAccommodationDetailImagesFailed   = 90017

	// file
	ErrCodeOpenFileFailed     = 100001
	ErrCodeReadFileFailed     = 100002
	ErrCodeInvalidFileType    = 100003
	ErrCodeGetFilesFailed     = 100006
	ErrCodeDeleteFileFailed   = 100007
	ErrCodeCreateFolderFailed = 100008
	ErrCodeUploadFileSuccess  = 100009
	ErrCodeUploadFileFailed   = 100010

	// manager
	ErrCodeUpdateManagerFailed        = 110001
	ErrCodeGetManagerFailed           = 110002
	ErrCodeGetManagerSuccess          = 110003
	ErrCodeCountNumberOfManagerFailed = 110004

	// admin
	ErrCodeGetAdminFailed = 120001
	ErrCodeUserNotAdmin   = 120002

	// facility
	ErrCodeCreateFacilitySuccess       = 130001
	ErrCodeCreateFacilityFailed        = 130002
	ErrCodeGetFacilityFailed           = 130003
	ErrCodeGetFacilitySuccess          = 130004
	ErrCodeUpdateFacilityFailed        = 130005
	ErrCodeUpdateFacilitySuccess       = 130006
	ErrCodeDeleteFacilityImageFailed   = 130007
	ErrCodeDeleteFacilityFailed        = 130008
	ErrCodeDeleteFacilitySuccess       = 130009
	ErrCodeUpdateFacilityDetailFailed  = 130010
	ErrCodeUpdateFacilityDetailSuccess = 130011
	ErrCodeDeleteFacilityDetailFailed  = 130012
	ErrCodeDeleteFacilityDetailSuccess = 130013

	// user base
	ErrCodeGetUserBaseFailed = 140001
	ErrCodeUserBaseNotFound  = 140002

	// user admin
	ErrCodeGetUserAdminFailed = 140003
	ErrCodeUserAdminNotFound  = 140004

	// transaction
	ErrCodeBeginTransactionFailed  = 150001
	ErrCodeCommitTransactionFailed = 150002

	// order
	ErrCodeCreateOrderDetailFailed  = 160001
	ErrCodeCreateOrderFailed        = 160002
	ErrCodeCreateOrderSuccess       = 160003
	ErrCodeGetOrderFailed           = 160004
	ErrCodeGetOrderByUserIDNotFound = 160005
	ErrCodeUpdateOrderStatusFailed  = 160006
	ErrCodeUpdateOrderStatusSuccess = 160007
	ErrCodeGetOrderSuccess          = 160008
	ErrCodeOrderNotFound            = 160009

	// payment
	ErrCodeCreatePaymentURLSuccess = 170001
	ErrCodeCreatePaymentFailed     = 170002
	ErrCodeGetPaymentFailed        = 170003

	// review
	ErrCodeCreateReviewFailed             = 180001
	ErrCodeCreateReviewSuccess            = 180002
	ErrCodeGetReviewByAccommodationFailed = 180003
	ErrCodeGetReviewsSuccess              = 180004

	// stats
	ErrCodeGetMonthlyEarningSuccess = 190001
	ErrCodeGetMonthlyEarningFailed  = 190002

	// Decimal
	ErrCodeInvalidPriceFormat  = 200001
	ErrCodePriceMustBePositive = 200002

	// accommodation room
	ErrCodeCheckAccommodationTypeBelongsToManagerFailed = 210001
	ErrCodeCheckAccommodationTypeNotBelongsToManager    = 210002
	ErrCodeCreateAccommodationRoomFailed                = 210003
	ErrCodeCreateAccommodationRoomSuccess               = 210004
	ErrCodeGetAccommodationRoomFailed                   = 210005
	ErrCodeGetAccommodationRoomSuccess                  = 210006
	ErrCodeCheckAccommodationRoomBelongsToManagerFailed = 210007
	ErrCodeCheckAccommodationRoomNotBelongsToManager    = 210008
	ErrCodeUpdateAccommodationRoomFailed                = 210009
	ErrCodeUpdateAccommodationRoomSuccess               = 210010
	ErrCodeDeleteAccommodationRoomFailed                = 210011
	ErrCodeDeleteAccommodationRoomSuccess               = 210012
	ErrCodeAccommodationRoomNotFound                    = 210013
)

var message = map[int]string{
	// global
	ErrCodeSaveDataFailed:         "Save data failed",
	ErrCodeUpdateDataFailed:       "Update data failed",
	ErrCodeMarshalFailed:          "Marshal failed",
	ErrCodeUnMarshalFailed:        "Unmarshal failed",
	ErrCodeCreateJWTTokenFailed:   "Create JWT token failed",
	ErrCodeParamsInvalid:          "Params invalid",
	ErrCodeUnauthorized:           "Unauthorized",
	ErrCodeForbidden:              "You do not have permission to access this resource.",
	ErrCodeValidatorNotFound:      "Validator not found",
	ErrCodeValidator:              "Validator error",
	ErrCodeConvertISOToUnixFailed: "Convert ISO to Unix failed",
	ErrCodeConvertUnixToISOFailed: "Convert Unix to ISO failed",
	ErrCodeSuccessfully:           "Success",
	ErrCodeParseTimeFailed:        "Parse time failed",

	// register
	ErrCodeInvalidEmailFormat: "Invalid email format",
	ErrCodeUserAlreadyExists:  "User already exists",
	ErrCodeOTPAlreadyExists:   "OTP already exists",
	ErrCodeSendEmailFailed:    "Send email failed",
	ErrCodeRegisterSuccess:    "Register successfully",
	ErrCodeRegisterFailed:     "Register failed",

	// verify otp
	ErrCodeOTPNotExists:           "OTP not exists",
	ErrCodeOTPNotMatch:            "OTP not match",
	ErrCodeGetInfoOTPFailed:       "Get info OTP failed",
	ErrCodeUpdateUserVerifyFailed: "Update user verify failed",
	ErrCodeVerifyOTPSuccess:       "Verify OTP successfully",
	ErrCodeOTPAlreadyVerified:     "OTP already verified",

	// update password register
	ErrCodeOTPNotVerified:                "OTP not verified",
	ErrCodeHashPasswordFailed:            "Hash password failed",
	ErrCodeUpdatePasswordRegisterSuccess: "Update password register successfully",

	// login
	ErrCodeGetUserInfoFailed:     "Get user infor failed",
	ErrCodePasswordNotMatch:      "Password not match",
	ErrCodeLoginSuccess:          "Login successfully",
	ErrCodeGetUserInfoSuccess:    "Get user infor successfully",
	ErrCodeGetUserInfoNotFound:   "User info not found",
	ErrCodeUpdateUserInfoFailed:  "Update user info failed",
	ErrCodeUpdateUserInfoSuccess: "Update user info successfully",

	// jwt
	ErrCodeMissAuthorizationHeader:    "Missing authorization header",
	ErrCodeInvalidAuthorizationFormat: "Invalid authorization format",
	ErrCodeInvalidToken:               "Invalid token",

	// accommodation
	ErrCodeCreateAccommodationFailed:   "Create accommodation failed",
	ErrCodeCreateAccommodationSuccess:  "Create accommodation successfully",
	ErrCodeGetAccommodationsFailed:     "Get accommodations failed",
	ErrCodeManagerNotFound:             "Manager not found",
	ErrCodeGetAccommodationSuccess:     "Get accommodations successfully",
	ErrCodeGetAccommodationFailed:      "Get accommodation failed",
	ErrCodeAccommodationNotFound:       "Accommodation not found",
	ErrCodeUpdateAccommodationSuccess:  "Update accommodation successfully",
	ErrCodeUpdateAccommodationFailed:   "Update accommodation failed",
	ErrCodeDeleteAccommodationFailed:   "Delete accommodation failed",
	ErrCodeDeleteAccommodationSuccess:  "Delete accommodation successfully",
	ErrCodeGetCountAccommodationFailed: "Get count accommodation failed",

	// accommodation image
	ErrCodeGetAccommodationImagesFailed:    "Get images of accommodation failed",
	ErrCodeDeleteAccommodationImagesFailed: "Delete images of accommodation failed",
	ErrCodeGetAccommodationImagesSuccess:   "Get images of accommodation success",
	ErrCodeSaveAccommodationImagesFailed:   "Save images of accommodation failed",

	// accommodation detail
	ErrCodeCreateAccommodationDetailFailed:   "Create accommodation details failed",
	ErrCodeCreateAccommodationDetailSuccess:  "Create accommodation details successfully",
	ErrCodeGetAccommodationDetailsFailed:     "Get accommodation details failed",
	ErrCodeGetAccommodationDetailsSuccess:    "Get accommodation details successfully",
	ErrCodeGetAccommodationDetailFailed:      "Get accommodation details failed",
	ErrCodeAccommodationDetailNotFound:       "Accommodation details not found",
	ErrCodeUpdateAccommodationDetailSuccess:  "Update accommodation detail successfully",
	ErrCodeUpdateAccommodationDetailFailed:   "Update accommodation detail failed",
	ErrCodeDeleteAccommodationDetailFailed:   "Delete accommodation detail failed",
	ErrCodeDeleteAccommodationDetailSuccess:  "Delete accommodation detail successfully",
	ErrCodeGetCountAccommodationDetailFailed: "Get count accommodation detail failed",
	ErrCodeNumberOfAvailableRoomsNotEnough:   "Number of available rooms not enough",

	// accommodation detail image
	ErrCodeGetAccommodationDetailImagesFailed:    "Get images of accommodation detail failed",
	ErrCodeDeleteAccommodationDetailImagesFailed: "Delete images of accommodation detail failed",
	ErrCodeGetAccommodationDetailImagesSuccess:   "Get images of accommodation detail success",
	ErrCodeSaveAccommodationDetailImagesFailed:   "Save images of accommodation detail failed",

	// file
	ErrCodeOpenFileFailed:     "Open file failed",
	ErrCodeReadFileFailed:     "Read file failed",
	ErrCodeInvalidFileType:    "Invalid file type",
	ErrCodeGetFilesFailed:     "Get files failed",
	ErrCodeDeleteFileFailed:   "Delete file failed",
	ErrCodeCreateFolderFailed: "Create folder failed",
	ErrCodeUploadFileSuccess:  "Upload files success",
	ErrCodeUploadFileFailed:   "Upload files failed",

	// manager
	ErrCodeUpdateManagerFailed:        "Update manager failed",
	ErrCodeGetManagerFailed:           "Get manager failed",
	ErrCodeGetManagerSuccess:          "Get manager successfully",
	ErrCodeCountNumberOfManagerFailed: "Count number of managers failed",

	// admin
	ErrCodeGetAdminFailed: "Get admin failed",

	// facility
	ErrCodeCreateFacilitySuccess:       "Create facility successfully",
	ErrCodeCreateFacilityFailed:        "Create facility failed",
	ErrCodeGetFacilityFailed:           "Get facility failed",
	ErrCodeGetFacilitySuccess:          "Get facility successfully",
	ErrCodeUpdateFacilityFailed:        "Update facility failed",
	ErrCodeUpdateFacilitySuccess:       "Update facility successfully",
	ErrCodeDeleteFacilityImageFailed:   "Detele facitliy image failed",
	ErrCodeDeleteFacilityFailed:        "Detele facitliy failed",
	ErrCodeDeleteFacilitySuccess:       "Detele facitliy successfully",
	ErrCodeUpdateFacilityDetailFailed:  "Update facitliy detail failed",
	ErrCodeUpdateFacilityDetailSuccess: "Update facitliy detail successfully",
	ErrCodeDeleteFacilityDetailFailed:  "Detele facitliy detail failed",
	ErrCodeDeleteFacilityDetailSuccess: "Detele facitliy detail successfully",

	// user base
	ErrCodeGetUserBaseFailed: "Get user base failed",
	ErrCodeUserBaseNotFound:  "User base not found",

	// user admin
	ErrCodeGetUserAdminFailed: "Get user admin failed",
	ErrCodeUserAdminNotFound:  "User admin not found",

	// transaction
	ErrCodeBeginTransactionFailed:  "Start transaction failed",
	ErrCodeCommitTransactionFailed: "Commit transaction failed",

	// order vs order detail
	ErrCodeCreateOrderDetailFailed:  "Create order detail failed",
	ErrCodeCreateOrderFailed:        "Create order failed",
	ErrCodeCreateOrderSuccess:       "Create order successfully",
	ErrCodeGetOrderFailed:           "Get order failed",
	ErrCodeGetOrderSuccess:          "Get order successfully",
	ErrCodeGetOrderByUserIDNotFound: "Get order by user id not found",
	ErrCodeUpdateOrderStatusFailed:  "Update order status failed",
	ErrCodeUpdateOrderStatusSuccess: "Update order status successfully",
	ErrCodeOrderNotFound:            "Order not found",

	// payment
	ErrCodeCreatePaymentURLSuccess: "Create payment url successfully",
	ErrCodeCreatePaymentFailed:     "Create payment failed",
	ErrCodeGetPaymentFailed:        "Get payment failed",

	// review
	ErrCodeCreateReviewFailed:             "Create review failed",
	ErrCodeCreateReviewSuccess:            "Create review successfully",
	ErrCodeGetReviewByAccommodationFailed: "Get reviews by accommodation failed",
	ErrCodeGetReviewsSuccess:              "Get reviews successfully",

	// stats
	ErrCodeGetMonthlyEarningSuccess: "Get monthly earning successfully",
	ErrCodeGetMonthlyEarningFailed:  "Get monthly earning failed",

	// Decimal
	ErrCodeInvalidPriceFormat: "Invalid price format",

	// accommodation type
	ErrCodeCheckAccommodationTypeBelongsToManagerFailed: "Check accommodation type belongs to manager failed",
	ErrCodeCheckAccommodationTypeNotBelongsToManager:    "Accommodation type not belongs to manager",
	ErrCodeCreateAccommodationRoomFailed:                "Create accommodation room failed",
	ErrCodeCreateAccommodationRoomSuccess:               "Create accommodation room successfully",
	ErrCodeGetAccommodationRoomFailed:                   "Get accommodation room failed",
	ErrCodeGetAccommodationRoomSuccess:                  "Get accommodation room successfully",
	ErrCodeCheckAccommodationRoomBelongsToManagerFailed: "Check accommodation room belongs to manager failed",
	ErrCodeCheckAccommodationRoomNotBelongsToManager:    "Accommodation room not belongs to manager",
	ErrCodeUpdateAccommodationRoomFailed:                "Update accommodation room failed",
	ErrCodeUpdateAccommodationRoomSuccess:               "Update accommodation room successfully",
	ErrCodeDeleteAccommodationRoomFailed:                "Delete accommodation room failed",
	ErrCodeDeleteAccommodationRoomSuccess:               "Delete accommodation room successfully",
	ErrCodeAccommodationRoomNotFound:                    "Accommodation room not found",
}
