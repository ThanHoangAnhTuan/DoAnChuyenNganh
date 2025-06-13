export interface Facility {
    id: string;
    name: string;
    image: string;
}
export interface FacilityDetail {
    id: string;
    name: string;
}

export interface UpdateFacility {
    id: string;
    name: string;
    image: string;
}

export interface CreateFacilityInput {
    name: string;
    image: string;
}
export interface CreateFacilityDetailInput {
    name: string;
    image: string;
}
export interface CreateFacilityOutput {
    code: number;
    message: string;
    data: Facility[];
}

export interface CreateFacilityDetailOutput {
    code: number;
    message: string;
    data: FacilityDetail[];
}

export interface GetFacilitiesOutput {
    code: number;
    message: string;
    data: Facility[];
}

export interface GetFacilitiesDetailOutput {
    code: number;
    message: string;
    data: FacilityDetail[];
}
export interface UpdateFacilityResponse {
    code: number;
    message: string;
    data: Facility[];
}
export interface DeleteFacilityResponse {
    data: null;
    code: number;
    message: string;
}
