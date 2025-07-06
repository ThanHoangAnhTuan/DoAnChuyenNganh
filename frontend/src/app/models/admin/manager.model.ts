import { Facility } from "../facility/facility.model";

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
    username: string;
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

export interface Rules {
    check_in: string,
    check_out: string,
    cancellation: string,
    refundable_damage_deposit: number,
    age_restriction: boolean,
    pet: boolean
}

export interface GetAccommodationsOfManagerByAdmin {
    id: string,
    name: string,
    city: string,
    country: string,
    district: string,
    address: string,
    description: string,
    rating: number,
    facilities: Facility[],
    images: string[],
    google_map: string,
    rules: Rules,
    is_verified: boolean,
    is_deleted: boolean
}

export interface GetAccommodationsOfManagerByAdminOutput {
    code: number,
    message: string,
    data: GetAccommodationsOfManagerByAdmin[]
}

export interface VerifyAccommodationInput {
    accommodation_id: string;
    status: boolean;
}

export interface SetDeletedAccommodationInput {
    accommodation_id: string;
    status: boolean;
}

export interface VerifyAccommodationOutput {
    code: number,
    message: string,
    data: null
}

export interface SetDeletedAccommodationOutput {
    code: number,
    message: string,
    data: null
}
