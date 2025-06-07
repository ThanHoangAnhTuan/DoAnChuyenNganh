export interface Review {
    id: string;
    accommodation_id: string;
    title: string;
    content: string;
    rating: number;
    author: string;
    avatar: string;
    created_at: string;
}

export interface GetReviewByIdResponse {
    data: Review;
    code: number;
    message: string;
}

export interface GetReviewsByAccommodationIdResponse {
    data: Review[];
    code: number;
    message: string;
}

// create accommodation
export interface CreateReview {
    title: string;
    content: string;
    rating: number;
}

export interface CreateNewReview {
    id: string;
    accommodation_id: string;
    title: string;
    content: string;
    rating: number;
    author: string;
    avatar: string;
    created_at: string;
}

export interface CreateAccommodationResponse {
    data: Review;
    code: number;
    message: string;
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
