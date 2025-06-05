export interface RegisterModel {
    verify_key: string;
    verify_type: number;
    verify_purpose: string;
}

export interface OTP {
    verify_key: string;
    verify_code: string;
}

export interface CreateAccount {
    verify_key: string;
}

export interface UpdatePassword {
    token: string;
    password: string;
}

export interface RegisterResponse {
    data: null;
    code: number;
    message: string;
}

export interface OTPResponse {
    data: { token: string };
    code: number;
    message: string;
}

export interface UpdateResponse {
    data: null;
    code: number;
    message: string;
}
