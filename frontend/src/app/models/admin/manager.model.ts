export interface Manager {
    id: string;
    account: string;
    username: string;
    is_deleted: boolean;
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

export interface GetManagerOutput {
    code: number;
    message: string;
    data: Manager[];
}
