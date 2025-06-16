export interface User {
    // id: string;
    // email: string;
    phone: string;
    username: string;
    gender: Gender;
    birthday: string;
}

export interface UpdateUser {
    // id: string;
    // email: string;
    phone: string;
    username: string;
    gender: number;
    birthday: string;
}

export interface UserResponse {
    code: number;
    message: string;
    data: User;
}
export enum Gender {
    Male = 'male',
    Female = 'female',
}
