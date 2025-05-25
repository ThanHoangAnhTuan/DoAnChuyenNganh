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
