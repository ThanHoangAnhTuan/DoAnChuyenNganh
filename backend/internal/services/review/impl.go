package review

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{
		sqlc: sqlc,
	}
}

func (r *serviceImpl) CreateReview(ctx *gin.Context, in *vo.CreateReviewInput) (codeStatus int, out *vo.CreateReviewOutput, err error) {
	out = &vo.CreateReviewOutput{}

	// TODO: get id in gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	account, err := r.sqlc.GetUserBaseByIdAndReturnAccount(ctx, userID)
	if err != nil {
		return response.ErrCodeGetUserBaseFailed, nil, fmt.Errorf("get user base failed: %s", err)
	}

	if account == "" {
		return response.ErrCodeUserBaseNotFound, nil, fmt.Errorf("user base not found")
	}

	// TODO: check accommodation exists
	accommodationExists, err := r.sqlc.CheckAccommodationExists(ctx, in.AccommodationID)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("get acommodation failed: %s", err)
	}

	if !accommodationExists {
		return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
	}

	// TODO: Check if the user has booked a room before
	booked, err := r.sqlc.CheckUserBookedOrder(ctx, database.CheckUserBookedOrderParams{
		UserID:          userID,
		OrderIDExternal: in.OrderIDExternal,
	})
	if err != nil {
		return response.ErrCodeGetOrderFailed, nil, fmt.Errorf("get order failed: %s", err)
	}

	if !booked {
		return response.ErrCodeGetOrderByUserIDNotFound, nil, fmt.Errorf("user not booked accommodation")
	}

	// TODO: Create review
	ID := uuid.NewString()
	now := utiltime.GetTimeNow()
	err = r.sqlc.CreateReview(ctx, database.CreateReviewParams{
		ID:              ID,
		UserID:          userID,
		AccommodationID: in.AccommodationID,
		Title:           in.Title,
		Comment:         in.Comment,
		Rating:          in.Rating,
		CreatedAt:       now,
		UpdatedAt:       now,
	})

	if err != nil {
		return response.ErrCodeCreateReviewFailed, nil, fmt.Errorf("create review failed: %s", err)
	}

	// TODO: get user info
	userInfo, err := r.sqlc.GetNameAndImageUserInfo(ctx, account)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, fmt.Errorf("get user info failed: %s", err)
	}
	out.Name = userInfo.UserName
	out.Image = userInfo.Image

	out.ID = ID
	out.Comment = in.Comment
	out.Rating = in.Rating
	out.Title = in.Title

	return response.ErrCodeCreateReviewSuccess, out, nil
}

func (r *serviceImpl) GetReviews(ctx *gin.Context, in *vo.GetReviewsInput) (codeStatus int, out []*vo.GetReviewOutput, pagination *vo.BasePaginationOutput, err error) {
	out = []*vo.GetReviewOutput{}

	page := in.GetPage()
	limit := in.GetLimit()

	// TODO: check accommodation exists
	accommodationExists, err := r.sqlc.CheckAccommodationExists(ctx, in.AccommodationID)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, nil, fmt.Errorf("get acommodation failed: %s", err)
	}

	if !accommodationExists {
		return response.ErrCodeAccommodationNotFound, nil, nil, fmt.Errorf("accommodation not found")
	}

	totalReviews, err := r.sqlc.CountReviewsByAccommodation(ctx, in.AccommodationID)
	if err != nil {
		return response.ErrCodeGetReviewByAccommodationFailed, nil, nil, fmt.Errorf("count reviews failed: %s", err)
	}

	offset := (page - 1) * limit

	// TODO: get reviews
	reviews, err := r.sqlc.GetReviewsWithPagination(ctx, database.GetReviewsWithPaginationParams{
		AccommodationID: in.AccommodationID,
		Limit:           limit,
		Offset:          offset,
	})
	if err != nil {
		return response.ErrCodeGetReviewByAccommodationFailed, nil, nil, fmt.Errorf("get reviews by accommodation failed: %s", err)
	}

	for _, review := range reviews {
		ID := uuid.NewString()

		// TODO: check user exists
		account, err := r.sqlc.GetUserBaseByIdAndReturnAccount(ctx, review.UserID)
		if err != nil {
			return response.ErrCodeGetUserBaseFailed, nil, nil, fmt.Errorf("get user base failed: %s", err)
		}

		if account == "" {
			return response.ErrCodeUserBaseNotFound, nil, nil, fmt.Errorf("user base not found")
		}

		// TODO: get info user
		userInfo, err := r.sqlc.GetNameAndImageUserInfo(ctx, account)
		if err != nil {
			return response.ErrCodeGetUserInfoFailed, nil, nil, fmt.Errorf("get user info failed: %s", err)
		}

		out = append(out, &vo.GetReviewOutput{
			ID:              ID,
			Name:            userInfo.UserName,
			Image:           userInfo.Image,
			Title:           review.Title,
			Comment:         review.Comment,
			ManagerResponse: review.ManagerResponse.String,
			Rating:          review.Rating,
		})
	}

	totalPages := (totalReviews + int64(limit) - 1) / int64(limit)
	pagination = &vo.BasePaginationOutput{
		Page:       page,
		Limit:      limit,
		Total:      totalReviews,
		TotalPages: totalPages,
	}

	return response.ErrCodeGetReviewsSuccess, out, pagination, nil
}
