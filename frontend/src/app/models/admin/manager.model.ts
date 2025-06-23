import { Facility } from "../facility/facility.model";
import { Accommodation } from "../manager/accommodation.model";

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


export interface GetAccommodationsOfManagerByAdmin {
    id: string;
    name: string;
    city: string;
    country: string;
    district: string;
    address: string;
    description: string;
    rating: number;
    facilities: Facility[];
    images: string[];
    google_map: string;
}

export interface GetAccommodationsOfManagerByAdminOutput {
    code: number;
    message: string;
    data: GetAccommodationsOfManagerByAdmin[];
}
