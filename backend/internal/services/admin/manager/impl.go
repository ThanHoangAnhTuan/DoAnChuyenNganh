package manager

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
	"go.uber.org/zap"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{sqlc: sqlc}
}

func (s *serviceImpl) GetManagers(ctx *gin.Context, in *vo.GetManagerInput) (codeStatus int, out []*vo.GetManagerOutput, pagination *vo.BasePaginationOutput, err error) {
	out = []*vo.GetManagerOutput{}

	// TODO: check user is admin
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	exists, err := s.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("get user admin failed: %s", err)
	}

	if !exists {
		return response.ErrCodeForbidden, nil, nil, fmt.Errorf("user admin not found")
	}

	// TODO: ger all manager have pagination

	page := in.GetPage()
	limit := in.GetLimit()

	totalManagers, err := s.sqlc.CountNumberOfManagers(ctx)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("count number of manager failed: %s", err)
	}

	offset := (page - 1) * limit
	managers, err := s.sqlc.GetManagers(ctx, database.GetManagersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("get managers failed: %s", err)
	}

	for _, manager := range managers {
		createdAt, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(manager.CreatedAt))
		if err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("convert ISO to Unix failed: %s", err)
		}

		updatedAt, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(manager.UpdatedAt))
		if err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("convert ISO to Unix failed: %s", err)
		}

		out = append(out, &vo.GetManagerOutput{
			ID:        manager.ID,
			Account:   manager.Account,
			Username:  manager.UserName,
			IsDeleted: manager.IsDeleted == 1,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	totalPages := (totalManagers + int64(limit) - 1) / int64(limit)
	pagination = &vo.BasePaginationOutput{
		Page:       page,
		Limit:      limit,
		Total:      totalManagers,
		TotalPages: totalPages,
	}
	return response.ErrCodeGetManagerSuccess, out, pagination, nil
}

func (s *serviceImpl) GetAccommodationsOfManager(ctx *gin.Context, in *vo.GetAccommodationsOfManagerInput) (codeStatus int, out []*vo.GetAccommodationsOfManagerOutput, pagination *vo.BasePaginationOutput, err error) {
	out = []*vo.GetAccommodationsOfManagerOutput{}

	page := in.GetPage()
	limit := in.GetLimit()

	var totalAccommodation int64
	var accommodationData []vo.AccommodationData

	// TODO: get accommodations
	totalAccommodation, err = s.sqlc.CountAccommodationOfManager(ctx, in.ManagerID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("count reviews failed: %s", err)
	}

	offset := (page - 1) * limit

	accommodations, err := s.sqlc.GetAccommodationsOfManagerWithPagination(ctx, database.GetAccommodationsOfManagerWithPaginationParams{
		Limit:     limit,
		Offset:    offset,
		ManagerID: in.ManagerID,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("error for get accommodations: %s", err)
	}

	for _, acc := range accommodations {
		accommodationData = append(accommodationData, vo.AccommodationData{
			ID:          acc.ID,
			ManagerID:   acc.ManagerID,
			Name:        acc.Name,
			Country:     acc.Country,
			City:        acc.City,
			District:    acc.District,
			Address:     acc.Address,
			Description: acc.Description,
			Rating:      acc.Rating,
			GgMap:       acc.GgMap,
			Facilities:  acc.Facilities,
			Rules:       acc.Rules,
			IsVerified:  acc.IsVerified == 1,
			IsDeleted:   acc.IsDeleted == 1,
		})
	}

	for _, accommodation := range accommodationData {
		// TODO: get facility
		var facilityIDs []string
		if err := json.Unmarshal(accommodation.Facilities, &facilityIDs); err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		facilities := []vo.FacilitiesOutput{}

		for _, facilityID := range facilityIDs {
			facility, err := s.sqlc.GetAccommodationFacilityById(ctx, facilityID)
			if err != nil {
				// TODO: Nếu không tìm thấy facility thì bỏ qua luôn thay vì báo lỗi
				fmt.Printf("Cannot found facility: %s", err.Error())
				global.Logger.Error("Cannot found facility: ", zap.Error(err))
				break
			}

			facilities = append(facilities, vo.FacilitiesOutput{
				ID:    facility.ID,
				Name:  facility.Name,
				Image: facility.Image,
			})
		}

		rules := vo.Rule{}
		if err := json.Unmarshal(accommodation.Rules, &rules); err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("error unmarshaling property surroundings: %s", err)
		}

		// TODO: get images of accommodation
		accommodationImages, err := s.sqlc.GetAccommodationImages(ctx, accommodation.ID)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		var imagePaths []string
		for _, i := range accommodationImages {
			imagePaths = append(imagePaths, i.Image)
		}

		out = append(out, &vo.GetAccommodationsOfManagerOutput{
			ID:          accommodation.ID,
			Name:        accommodation.Name,
			Country:     accommodation.Country,
			City:        accommodation.City,
			District:    accommodation.District,
			Address:     accommodation.Address,
			Description: accommodation.Description,
			Rating:      accommodation.Rating,
			GoogleMap:   accommodation.GgMap,
			Facilities:  facilities,
			Rules:       rules,
			Images:      imagePaths,
			IsVerified:  accommodation.IsVerified,
			IsDeleted:   accommodation.IsDeleted,
		})
	}

	totalPages := (totalAccommodation + int64(limit) - 1) / int64(limit)
	pagination = &vo.BasePaginationOutput{
		Page:       page,
		Limit:      limit,
		Total:      totalAccommodation,
		TotalPages: totalPages,
	}

	return response.ErrCodeGetAccommodationSuccess, out, pagination, nil
}

func (s *serviceImpl) VerifyAccommodation(ctx *gin.Context, in *vo.VerifyAccommodationInput) (codeStatus int, err error) {
	// TODO: check user is admin
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	exists, err := s.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get user admin failed: %s", err)
	}

	if !exists {
		return response.ErrCodeForbidden, fmt.Errorf("user admin not found")
	}
	// TODO: check accommodation exists
	accommodationExists, err := s.sqlc.CheckAccommodationExists(ctx, in.AccommodationID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get acommodation failed: %s", err)
	}

	if !accommodationExists {
		return response.ErrCodeAccommodationNotFound, fmt.Errorf("accommodation not found")
	}
	// TODO: update status of accommodation
	err = s.sqlc.UpdateStatusAccommodation(ctx, database.UpdateStatusAccommodationParams{
		IsVerified: in.Status,
		ID:         in.AccommodationID,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("update status acommodation failed: %s", err)
	}

	return response.ErrCodeUpdateAccommodationSuccess, nil
}
