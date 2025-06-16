export interface Manager {
    account: string;
    user_name: string;
    login_time: string;
    logout_time: string;
    is_deleted: string;
    created_at: string;
    updated_at: string;
}

export interface CreateManager {
    account: string;
    password: string;
}

export interface CreateManagerOutput {
    code: number;
    message: string;
    data: null;
}