package response

const (
	// global
	ErrCodeSaveDataFailed       = 20001
	ErrCodeUpdateDataFailed     = 20002
	ErrCodeMarshalFailed        = 20003
	ErrCodeCreateJWTTokenFailed = 20004
	ErrCodeParamsInvalid        = 20005

	// register
	ErrCodeInvalidEmailFormat = 30001
	ErrCodeUserAlreadyExists  = 30002
	ErrCodeOTPAlreadyExists   = 30003
	ErrCodeSendEmailFailed    = 30004
	ErrCodeRegisterSuccess    = 30005

	// verify otp
	ErrCodeOTPNotExists           = 40001
	ErrCodeOTPNotMatch            = 40002
	ErrCodeGetInfoOTPFailed       = 40003
	ErrCodeUpdateUserVerifyFailed = 40004
	ErrCodeVerifyOTPSuccess       = 40005

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

	// test
	ErrCodeCreateAccommodationFailed  = 80001
	ErrCodeCreateAccommodationSuccess = 80002
)

var message = map[int]string{
	// global
	ErrCodeSaveDataFailed:       "Save data failed",
	ErrCodeUpdateDataFailed:     "Update data failed",
	ErrCodeMarshalFailed:        "Marshal failed",
	ErrCodeCreateJWTTokenFailed: "Create JWT token failed",
	ErrCodeParamsInvalid:        "Params invalid",

	// register
	ErrCodeInvalidEmailFormat: "Invalid email format",
	ErrCodeUserAlreadyExists:  "User already exists",
	ErrCodeOTPAlreadyExists:   "OTP already exists",
	ErrCodeSendEmailFailed:    "Send email failed",
	ErrCodeRegisterSuccess:    "Register successfully",

	// verify otp
	ErrCodeOTPNotExists:           "OTP not exists",
	ErrCodeOTPNotMatch:            "OTP not match",
	ErrCodeGetInfoOTPFailed:       "Get info OTP failed",
	ErrCodeUpdateUserVerifyFailed: "Update user verify failed",
	ErrCodeVerifyOTPSuccess:       "Verify OTP successfully",

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

	// test
	ErrCodeCreateAccommodationFailed:  "Create accommodation failed",
	ErrCodeCreateAccommodationSuccess: "Create accommodation successfully",
}
