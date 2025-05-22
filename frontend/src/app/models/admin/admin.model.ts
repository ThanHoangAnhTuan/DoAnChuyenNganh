export interface AdminLoginInput {
    account: string;
    password: string;
}

export interface AdminLoginOutput {
    code: number;
    message: string;
    data: {
        token: string;
        account: string;
        userName: string;
    };
}