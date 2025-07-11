package accommodation

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (t *serviceImpl) GetAccommodation(ctx *gin.Context, in *vo.GetAccommodationInput) (codeStatus int, out *vo.GetAccommodationOutput, err error) {
	out = &vo.GetAccommodationOutput{}

	// TODO: get accommodation by id
	accommodation, err := t.sqlc.GetAccommodationById(ctx, in.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get accommodation by id: %s", err)
	}

	// TODO: get facility
	var facilityIDs []string
	if err := json.Unmarshal(accommodation.Facilities, &facilityIDs); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("unmarshaling facilities: %s", err)
	}

	facilities := []vo.FacilitiesOutput{}

	for _, facilityID := range facilityIDs {
		facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
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
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("unmarshaling rules: %s", err)
	}

	// TODO: get images of accommodation
	images, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get images of accommodation by id failed: %s", err)
	}

	var imagesName []string
	for _, img := range images {
		imagesName = append(imagesName, img.Image)
	}

	out = &vo.GetAccommodationOutput{
		ID:          accommodation.ID,
		ManagerID:   accommodation.ManagerID,
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
		Images:      imagesName,
	}

	return response.ErrCodeGetAccommodationSuccess, out, nil
}

func (t *serviceImpl) GetAccommodationsByManager(ctx *gin.Context, in *vo.GetAccommodationsInput) (codeStatus int, out []*vo.GetAccommodationsOutput, pagination *vo.BasePaginationOutput, err error) {
	out = []*vo.GetAccommodationsOutput{}

	// TODO: get managerID from context
	managerID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, nil, fmt.Errorf("userID not found in context")
	}

	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, managerID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, nil, nil, fmt.Errorf("manager not found")
	}

	page := in.GetPage()
	limit := in.GetLimit()

	var totalAccommodation int64
	var accommodationData []vo.AccommodationData

	// TODO: get accommodations
	totalAccommodation, err = t.sqlc.CountAccommodationByManager(ctx, managerID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("count accommodations failed: %s", err)
	}

	offset := (page - 1) * limit

	accommodations, err := t.sqlc.GetAccommodationsByManagerWithPagination(ctx, database.GetAccommodationsByManagerWithPaginationParams{
		ManagerID: managerID,
		Limit:     limit,
		Offset:    offset,
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
			facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
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
		accommodationImages, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		var imagePaths []string
		for _, i := range accommodationImages {
			imagePaths = append(imagePaths, i.Image)
		}

		out = append(out, &vo.GetAccommodationsOutput{
			ID:          accommodation.ID,
			ManagerID:   accommodation.ManagerID,
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

func (t *serviceImpl) DeleteAccommodation(ctx *gin.Context, in *vo.DeleteAccommodationInput) (codeResult int, err error) {
	// TODO: get userId from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check manager exists in database
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, fmt.Errorf("manager not found")
	}

	// TODO: check accommodation exists in database
	accommodation, err := t.sqlc.GetAccommodationByIdNoVerify(ctx, in.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeInternalServerError, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userID {
		return response.ErrCodeForbidden, fmt.Errorf("user is not the owner of the accommodation")
	}

	// TODO: delete accommodation
	err = t.sqlc.DeleteAccommodation(ctx, database.DeleteAccommodationParams{
		ID:        accommodation.ID,
		UpdatedAt: utiltime.GetTimeNow(),
	})
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for delete accommodation: %s", err)
	}

	return response.ErrCodeDeleteAccommodationSuccess, nil
}

func (t *serviceImpl) UpdateAccommodation(ctx *gin.Context, in *vo.UpdateAccommodationInput) (codeResult int, out *vo.UpdateAccommodationOutput, err error) {
	out = &vo.UpdateAccommodationOutput{}

	// TODO: get userId from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check manager exists in database
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, nil, fmt.Errorf("manager not found")
	}

	// TODO: get accommodation in database
	accommodation, err := t.sqlc.GetAccommodationByIdNoVerify(ctx, in.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userID {
		return response.ErrCodeForbidden, nil, fmt.Errorf("user is not the owner of the accommodation")
	}

	// TODO: update accommodation
	now := utiltime.GetTimeNow()
	facilitiesJSON, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	rulesJSON, err := json.Marshal(in.Rules)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for marshal rules: %s", err)
	}

	err = t.sqlc.UpdateAccommodation(ctx, database.UpdateAccommodationParams{
		ID:          accommodation.ID,
		Name:        in.Name,
		Country:     in.Country,
		City:        in.City,
		District:    in.District,
		Description: in.Description,
		GgMap:       in.GoogleMap,
		Address:     in.Address,
		Facilities:  facilitiesJSON,
		Rules:       rulesJSON,
		UpdatedAt:   now,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for update accommodation: %s", err)
	}

	// TODO: get facility
	for _, facilityID := range in.Facilities {
		facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
		if err != nil {
			// TODO: Nếu không tìm thấy facility thì bỏ qua luôn thay vì báo lỗi
			fmt.Printf("Cannot found facility: %s", err.Error())
			global.Logger.Error("Cannot found facility: ", zap.Error(err))
			break
		}

		out.Facilities = append(out.Facilities, vo.FacilitiesOutput{
			ID:    facility.ID,
			Name:  facility.Name,
			Image: facility.Image,
		})
	}

	// TODO: return response
	out.ID = accommodation.ID
	out.ManagerID = accommodation.ManagerID
	out.Name = in.Name
	out.City = in.City
	out.Country = in.Country
	out.District = in.District
	out.Description = in.Description
	out.Address = in.Address
	out.GoogleMap = in.GoogleMap
	out.Rules = in.Rules
	out.Rating = accommodation.Rating
	out.IsDeleted = accommodation.IsDeleted == 1
	out.IsVerified = accommodation.IsVerified == 1

	return response.ErrCodeUpdateAccommodationSuccess, out, nil
}

func (t *serviceImpl) GetAccommodations(ctx *gin.Context, in *vo.GetAccommodationsInput) (codeStatus int, out []*vo.GetAccommodationsOutput, pagination *vo.BasePaginationOutput, err error) {
	out = []*vo.GetAccommodationsOutput{}

	page := in.GetPage()
	limit := in.GetLimit()

	var totalAccommodation int64
	var accommodationData []vo.AccommodationData

	// TODO: get accommodations by city
	if in.City != "" {
		totalAccommodation, err = t.sqlc.CountAccommodationByCity(ctx, in.City)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("count reviews failed: %s", err)
		}

		offset := (page - 1) * limit

		accommodations, err := t.sqlc.GetAccommodationsByCityWithPagination(ctx, database.GetAccommodationsByCityWithPaginationParams{
			City:   in.City,
			Limit:  limit,
			Offset: offset,
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
			})
		}
	} else {
		// TODO: get accommodations
		totalAccommodation, err = t.sqlc.CountAccommodation(ctx)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("count reviews failed: %s", err)
		}

		offset := (page - 1) * limit

		accommodations, err := t.sqlc.GetAccommodationsWithPagination(ctx, database.GetAccommodationsWithPaginationParams{
			Limit:  limit,
			Offset: offset,
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
			})
		}
	}

	for _, accommodation := range accommodationData {
		// TODO: get facility
		var facilityIDs []string
		if err := json.Unmarshal(accommodation.Facilities, &facilityIDs); err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		facilities := []vo.FacilitiesOutput{}

		for _, facilityID := range facilityIDs {
			facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
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
		accommodationImages, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		var imagePaths []string
		for _, i := range accommodationImages {
			imagePaths = append(imagePaths, i.Image)
		}

		out = append(out, &vo.GetAccommodationsOutput{
			ID:          accommodation.ID,
			ManagerID:   accommodation.ManagerID,
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

func (t *serviceImpl) CreateAccommodation(ctx *gin.Context, in *vo.CreateAccommodationInput) (codeResult int, out *vo.CreateAccommodationOutput, err error) {
	out = &vo.CreateAccommodationOutput{}

	// TODO: get userId from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check manager exists in database
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, nil, fmt.Errorf("manager not found")
	}

	// TODO: convert struct to json
	now := utiltime.GetTimeNow()
	id := uuid.New().String()

	facilitiesJSON, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	rulesJSON, err := json.Marshal(in.Rules)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for marshal rules: %s", err)
	}

	// TODO: create accommodation
	err = t.sqlc.CreateAccommodation(ctx, database.CreateAccommodationParams{
		ID:          id,
		ManagerID:   userID,
		Name:        in.Name,
		Country:     in.Country,
		City:        in.City,
		District:    in.District,
		Description: in.Description,
		Address:     in.Address,
		GgMap:       in.GoogleMap,
		Facilities:  facilitiesJSON,
		Rules:       rulesJSON,
		Rating:      in.Rating,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for create accommodation: %s", err)
	}

	// TODO: get accommodation
	accommodation, err := t.sqlc.GetAccommodationByIdNoVerify(ctx, id)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get accommodation failed: %s", err)
	}

	out.ID = accommodation.ID
	out.ManagerID = accommodation.ManagerID
	out.Name = accommodation.Name
	out.City = accommodation.City
	out.Country = accommodation.Country
	out.District = accommodation.District
	out.Description = accommodation.Description
	out.GoogleMap = accommodation.GgMap
	out.Address = accommodation.Address
	out.Rating = accommodation.Rating
	out.IsDeleted = accommodation.IsDeleted == 1
	out.IsVerified = accommodation.IsVerified == 1

	// TODO: get facility
	var facilitieIDs []string
	if err := json.Unmarshal(accommodation.Facilities, &facilitieIDs); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
	}

	for _, facilityID := range facilitieIDs {
		facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("get facility failed: %s", err)
		}

		out.Facilities = append(out.Facilities, vo.FacilitiesOutput{
			ID:    facility.ID,
			Name:  facility.Name,
			Image: facility.Image,
		})
	}

	err = json.Unmarshal(accommodation.Rules, &out.Rules)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for unmarshal rules: %s", err)
	}

	// TODO: get images of accommodation
	accommodationImages, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get images of accommodation failed: %s", err)
	}

	for _, i := range accommodationImages {
		out.Images = append(out.Images, i.Image)
	}

	return response.ErrCodeCreateAccommodationSuccess, out, nil
}
