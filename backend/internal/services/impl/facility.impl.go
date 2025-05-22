package impl

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

type FacilityImpl struct {
	sqlc *database.Queries
}

func (f *FacilityImpl) CreateFacility(ctx *gin.Context, in *vo.CreateFacilityInput) (codeStatus int, out *vo.CreateFacilityOutput, err error) {
	out = &vo.CreateFacilityOutput{}

	// TODO: get user id in gin.Context
	val, exists := ctx.Get("userId")
	if !exists {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("unauthorized")
	}

	userID, ok := val.(string)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("invalid user id format")
	}

	// TODO: check user is admin
	isExists, err := f.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeGetAdminFailed, nil, fmt.Errorf("get admin failed")
	}

	if !isExists {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user not admin")
	}

	// TODO: save image to disk
	codeStatus, imagesFileName, err := saveImageToDisk(ctx, []*multipart.FileHeader{in.Image})
	if err != nil {
		return codeStatus, nil, err
	}

	// TODO: save facility
	id := uuid.New().String()
	now := utiltime.GetTimeNow()
	err = f.sqlc.CreateAccommodationFacility(ctx, database.CreateAccommodationFacilityParams{
		ID:        id,
		Image:     imagesFileName[0],
		Name:      in.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return response.ErrCodeCreateFacilityFailed, nil, fmt.Errorf("create facility failed: %s", err)
	}

	// TODO: get facility
	facility, err := f.sqlc.GetAccommodationFacilityById(ctx, id)
	if err != nil {
		return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
	}

	out.Id = facility.ID
	out.Image = facility.Image
	out.Name = facility.Name

	return response.ErrCodeCreateFacilitySuccess, out, nil
}

func (f *FacilityImpl) GetFacilities(ctx context.Context) (codeStatus int, out []*vo.GetFacilitiesOutput, err error) {
	// TODO: get user id in gin.Context
	val := ctx.Value("userId")
	if val == nil {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("unauthorized")
	}
	userID, ok := val.(string)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("invalid user id format")
	}

	// TODO: check user is admin
	isExists, err := f.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeGetAdminFailed, nil, fmt.Errorf("get admin failed")
	}

	if !isExists {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user not admin")
	}

	// TODO: get facilities
	facilities, err := f.sqlc.GetAccommodationFacilityNames(ctx)
	if err != nil {
		return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
	}

	for _, facility := range facilities {
		out = append(out, &vo.GetFacilitiesOutput{
			Id:    facility.ID,
			Name:  facility.Name,
			Image: facility.Image,
		})
	}

	return response.ErrCodeGetFacilitySuccess, out, nil
}

func NewFacilityImpl(sqlc *database.Queries) *FacilityImpl {
	return &FacilityImpl{
		sqlc: sqlc,
	}
}
