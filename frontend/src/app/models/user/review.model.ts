export interface Review {
    id: string,
    name: string,
    image: string,
    title: string,
    comment: string,
    manager_response: string,
    rating: number,
}

export interface Pagination {
    page: number,
    limit: number,
    total: number,
    total_pages: number,
}

export interface GetReviewByIdResponse {
    data: Review;
    code: number;
    message: string;
}

export interface GetReviewsByAccommodationIdResponse {
    code: number,
    message: string;
    data: Review[];
    pagination: Pagination,
}

// create accommodation
export interface NewReview {
    id: string,
    name: string,
    image: string,
    title: string,
    comment: string,
    rating: number,
}

export interface CreateNewReview {
    accommodation_id: string;
    title: string;
    comment: string;
    rating: number;
    order_id: string;
}

export interface CreateAccommodationResponse {
    code: number;
    message: string;
    data: NewReview;
}

// update accommodation
export interface UpdateReview {
    title: string;
    content: string;
    rating: number;
}

export interface UpdateAccommodationResponse {
    data: Review;
    code: number;
    message: string;
}

// delete accommodation
export interface DeleteAccommodationResponse {
    data: null;
    code: number;
    message: string;
}
