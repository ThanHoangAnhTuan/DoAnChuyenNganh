package impl

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type FacilityDetailImpl struct {
	sqlc *database.Queries
}

func (f *FacilityDetailImpl) CreateFacilityDetail(ctx *gin.Context, in *vo.CreateFacilityDetailInput) (codeStatus int, out *vo.CreateFacilityDetailOutput, err error) {
	out = &vo.CreateFacilityDetailOutput{}

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

	// TODO: save facility
	id := uuid.New().String()
	now := utiltime.GetTimeNow()
	err = f.sqlc.CreateAccommodationFacilityDetail(ctx, database.CreateAccommodationFacilityDetailParams{
		ID:        id,
		Name:      in.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return response.ErrCodeCreateFacilityFailed, nil, fmt.Errorf("create facility failed: %s", err)
	}

	// TODO: get facility
	facility, err := f.sqlc.GetAccommodationFacilityDetailById(ctx, id)
	if err != nil {
		return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
	}

	out.Id = facility.ID
	out.Name = facility.Name
	return response.ErrCodeCreateFacilitySuccess, out, nil
}

func (f *FacilityDetailImpl) GetFacilityDetail(ctx context.Context) (codeStatus int, out []*vo.GetFacilityDetailOutput, err error) {
	// TODO: get facilities
	facilities, err := f.sqlc.GetAccommodationFacilityDetail(ctx)
	if err != nil {
		return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
	}

	for _, facility := range facilities {
		out = append(out, &vo.GetFacilityDetailOutput{
			Id:   facility.ID,
			Name: facility.Name,
		})
	}

	return response.ErrCodeGetFacilitySuccess, out, nil
}

func NewFacilityDetailImpl(sqlc *database.Queries) *FacilityDetailImpl {
	return &FacilityDetailImpl{
		sqlc: sqlc,
	}
}
