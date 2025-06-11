export interface Facility {
    id: string;
    name: string;
    image: string;
}

export interface GetFacilitiesOutput {
    code: number;
    message: string;
    data: Facility[];
}

export interface FacilityDetail {
    id: string;
    name: string;
}

export interface GetFacilitiesDetailOutput {
    code: number;
    message: string;
    data: FacilityDetail[];
}
