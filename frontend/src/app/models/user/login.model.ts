export interface LoginModel {
    account: string;
    password: string;
}

export interface LoginResponse {
    data: {
        token: string;
    };
    code: number;
    message: string;
}
