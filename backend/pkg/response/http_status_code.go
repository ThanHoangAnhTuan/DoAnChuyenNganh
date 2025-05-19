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
	ErrCodeConvertISOToUnixFailed = 200011
	ErrCodeConvertUnixToISOFailed = 200012

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
	ErrCodeGetUserInfoFailed = 60001
	ErrCodePasswordNotMatch  = 60002
	ErrCodeLoginSuccess      = 60003

	// jwt
	ErrCodeMissAuthorizationHeader    = 70001
	ErrCodeInvalidAuthorizationFormat = 70002
	ErrCodeInvalidToken               = 70003

	// accommodation
	ErrCodeCreateAccommodationFailed  = 80001
	ErrCodeCreateAccommodationSuccess = 80002
	ErrCodeManagerNotFound            = 80003
	ErrCodeGetAccommodationsFailed    = 80004
	ErrCodeGetAccommodationSuccess    = 80005
	ErrCodeGetAccommodationFailed     = 80006
	ErrCodeAccommodationNotFound      = 80007
	ErrCodeUpdateAccommodationSuccess = 80008
	ErrCodeUpdateAccommodationFailed  = 80009
	ErrCodeDeleteAccommodationFailed  = 80010
	ErrCodeDeleteAccommodationSuccess = 80011

	// accommodation image
	ErrCodeGetAccommodationImagesFailed    = 800012
	ErrCodeGetAccommodationImagesSuccess   = 800013
	ErrCodeDeleteAccommodationImagesFailed = 800014
	ErrCodeSaveAccommodationImagesFailed   = 800015

	// accommodation detail
	ErrCodeCreateAccommodationDetailFailed  = 90001
	ErrCodeCreateAccommodationDetailSuccess = 90002
	ErrCodeGetAccommodationDetailsFailed    = 90004
	ErrCodeGetAccommodationDetailsSuccess   = 90005
	ErrCodeGetAccommodationDetailFailed     = 90006
	ErrCodeAccommodationDetailNotFound      = 90007
	ErrCodeUpdateAccommodationDetailSuccess = 90008
	ErrCodeUpdateAccommodationDetailFailed  = 90009
	ErrCodeDeleteAccommodationDetailFailed  = 90010
	ErrCodeDeleteAccommodationDetailSuccess = 90011

	// accommodation detail image
	ErrCodeGetAccommodationDetailImagesFailed    = 900012
	ErrCodeGetAccommodationDetailImagesSuccess   = 900013
	ErrCodeDeleteAccommodationDetailImagesFailed = 900014
	ErrCodeSaveAccommodationDetailImagesFailed   = 900015

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
	ErrCodeUpdateManagerFailed = 110001
	ErrCodeGetManagerFailed    = 110002

	// admin
	ErrCodeGetAdminFailed = 120001
	ErrCodeUserNotAdmin   = 120002

	// facility
	ErrCodeCreateFacilitySuccess = 130001
	ErrCodeCreateFacilityFailed  = 130002
	ErrCodeGetFacilityFailed     = 130003

	// user base
	ErrCodeGetUserBaseFailed = 140001
	ErrCodeUserBaseNotFound  = 140002

	// transaction
	ErrCodeBeginTransactionFailed  = 150001
	ErrCodeCommitTransactionFailed = 150002

	// order
	ErrCodeCreateOrderDetailFailed = 160001
	ErrCodeCreateOrderFailed       = 160002
	ErrCodeCreatePaymentFailed     = 160003
	ErrCodeCreateOrderSuccess      = 160004
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
	ErrCodeGetUserInfoFailed: "Get user infor failed",
	ErrCodePasswordNotMatch:  "Password not match",
	ErrCodeLoginSuccess:      "Login successfully",

	// jwt
	ErrCodeMissAuthorizationHeader:    "Missing authorization header",
	ErrCodeInvalidAuthorizationFormat: "Invalid authorization format",
	ErrCodeInvalidToken:               "Invalid token",

	// accommodation
	ErrCodeCreateAccommodationFailed:  "Create accommodation failed",
	ErrCodeCreateAccommodationSuccess: "Create accommodation successfully",
	ErrCodeGetAccommodationsFailed:    "Get accommodations failed",
	ErrCodeManagerNotFound:            "Manager not found",
	ErrCodeGetAccommodationSuccess:    "Get accommodations successfully",
	ErrCodeGetAccommodationFailed:     "Get accommodation failed",
	ErrCodeAccommodationNotFound:      "Accommodation not found",
	ErrCodeUpdateAccommodationSuccess: "Update accommodation successfully",
	ErrCodeUpdateAccommodationFailed:  "Update accommodation failed",
	ErrCodeDeleteAccommodationFailed:  "Delete accommodation failed",
	ErrCodeDeleteAccommodationSuccess: "Delete accommodation successfully",

	// accommodation image
	ErrCodeGetAccommodationImagesFailed:    "Get images of accommodation failed",
	ErrCodeDeleteAccommodationImagesFailed: "Delete images of accommodation failed",
	ErrCodeGetAccommodationImagesSuccess:   "Get images of accommodation success",
	ErrCodeSaveAccommodationImagesFailed:   "Save images of accommodation failed",

	// accommodation detail
	ErrCodeCreateAccommodationDetailFailed:  "Create accommodation details failed",
	ErrCodeCreateAccommodationDetailSuccess: "Create accommodation details successfully",
	ErrCodeGetAccommodationDetailsFailed:    "Get accommodation details failed",
	ErrCodeGetAccommodationDetailsSuccess:   "Get accommodation details successfully",
	ErrCodeGetAccommodationDetailFailed:     "Get accommodation details failed",
	ErrCodeAccommodationDetailNotFound:      "Accommodation details not found",
	ErrCodeUpdateAccommodationDetailSuccess: "Update accommodation detail successfully",
	ErrCodeUpdateAccommodationDetailFailed:  "Update accommodation detail failed",
	ErrCodeDeleteAccommodationDetailFailed:  "Delete accommodation detail failed",
	ErrCodeDeleteAccommodationDetailSuccess: "Delete accommodation detail successfully",

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
	ErrCodeUpdateManagerFailed: "Update manager failed",
	ErrCodeGetManagerFailed:    "Get manager failed",

	// admin
	ErrCodeGetAdminFailed: "Get admin failed",

	// facility
	ErrCodeCreateFacilitySuccess: "Create facility successfully",
	ErrCodeCreateFacilityFailed:  "Create facility failed",
	ErrCodeGetFacilityFailed:     "Get facility failed",

	// user base
	ErrCodeGetUserBaseFailed: "Get user base failed",
	ErrCodeUserBaseNotFound:  "User base not found",

	// transaction
	ErrCodeBeginTransactionFailed:  "Start transaction failed",
	ErrCodeCommitTransactionFailed: "Commit transaction failed",

	// order vs order detail
	ErrCodeCreateOrderDetailFailed: "Create order detail failed",
	ErrCodeCreateOrderFailed:       "Create order failed",
	ErrCodeCreatePaymentFailed:     "Create payment failed",
	ErrCodeCreateOrderSuccess:      "Create order successfully",
}
