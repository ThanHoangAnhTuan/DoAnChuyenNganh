package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type AccommodationImpl struct {
	sqlc *database.Queries
}

func (t *AccommodationImpl) GetAccommodationById(ctx context.Context, in *vo.GetAccommodationByIdInput) (codeStatus int, out *vo.GetAccommodationByIdOutput, err error) {
	out = &vo.GetAccommodationByIdOutput{}

	// TODO: get accommodation by id
	accommodation, err := t.sqlc.GetAccommodationById(ctx, in.ID)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation by id: %s", err)
	}

	// TODO: get facility
	var facilityIDs []string
	if err := json.Unmarshal(accommodation.Facilities, &facilityIDs); err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
	}

	facilities := []vo.FacilitiesOutput{}

	for _, facilityID := range facilityIDs {
		facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
		if err != nil {
			return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
		}

		facilities = append(facilities, vo.FacilitiesOutput{
			ID:    facility.ID,
			Name:  facility.Name,
			Image: facility.Image,
		})
	}

	rules := vo.Rule{}
	if err := json.Unmarshal(accommodation.Rules, &rules); err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling rules: %s", err)
	}

	// TODO: get images of accommodation
	images, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
	if err != nil {
		return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("error for get images of accommodation by id failed: %s", err)
	}

	var imagesName []string
	for _, img := range images {
		imagesName = append(imagesName, img.Image)
	}

	out = &vo.GetAccommodationByIdOutput{
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

func (t *AccommodationImpl) GetAccommodationByCity(ctx context.Context, in *vo.GetAccommodationByCityInput) (codeStatus int, out []*vo.GetAccommodationsByCity, err error) {
	out = []*vo.GetAccommodationsByCity{}
	accommodations, err := t.sqlc.GetAccommodationsByCity(ctx, in.City)
	if err != nil {
		return response.ErrCodeGetAccommodationsFailed, nil, fmt.Errorf("error for get accommodations: %s", err)
	}

	for _, accommodation := range accommodations {
		// TODO: get image of accommodation
		images, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailImagesFailed, nil, fmt.Errorf("get accommodation image failed: %s", err)
		}

		var imagesName []string
		for _, img := range images {
			imagesName = append(imagesName, img.Image)
		}

		out = append(out, &vo.GetAccommodationsByCity{
			ID:        accommodation.ID,
			Name:      accommodation.Name,
			Country:   accommodation.Country,
			City:      accommodation.City,
			Address:   accommodation.Address,
			District:  accommodation.District,
			Rating:    accommodation.Rating,
			GoogleMap: accommodation.GgMap,
			Images:    imagesName,
		})
	}
	return response.ErrCodeGetAccommodationSuccess, out, nil
}

func (t *AccommodationImpl) GetAccommodationsByManager(ctx context.Context) (codeStatus int, out []*vo.GetAccommodations, err error) {
	out = []*vo.GetAccommodations{}

	// TODO: get userId from context
	userID, ok := utils.GetUserIDFromContext(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	accommodations, err := t.sqlc.GetAccommodationsByManager(ctx, userID)
	if err != nil {
		return response.ErrCodeGetAccommodationsFailed, nil, fmt.Errorf("error for get accommodations by manager: %s", err)
	}

	for _, accommodation := range accommodations {
		// TODO: get facility
		var facilityIDs []string
		if err := json.Unmarshal(accommodation.Facilities, &facilityIDs); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		facilities := []vo.FacilitiesOutput{}

		for _, facilityID := range facilityIDs {
			facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
			if err != nil {
				return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
			}

			facilities = append(facilities, vo.FacilitiesOutput{
				ID:    facility.ID,
				Name:  facility.Name,
				Image: facility.Image,
			})
		}

		rules := vo.Rule{}
		if err := json.Unmarshal(accommodation.Rules, &rules); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling rules: %s", err)
		}

		// TODO: get images of accommodation
		accommodationImages, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		var imagePaths []string
		for _, i := range accommodationImages {
			imagePaths = append(imagePaths, i.Image)
		}

		out = append(out, &vo.GetAccommodations{
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
		})
	}
	return response.ErrCodeGetAccommodationSuccess, out, nil
}

func (t *AccommodationImpl) DeleteAccommodation(ctx context.Context, in *vo.DeleteAccommodationInput) (codeResult int, err error) {
	// TODO: get userId from context
	userID, ok := utils.GetUserIDFromContext(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check manager exists in database
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeManagerNotFound, fmt.Errorf("manager not found")
	}

	// TODO: check accommodation exists in database
	accommodation, err := t.sqlc.GetAccommodationById(ctx, in.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeGetAccommodationFailed, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userID {
		return response.ErrCodeUnauthorized, fmt.Errorf("user is not the owner of the accommodation")
	}

	// TODO: delete accommodation
	err = t.sqlc.DeleteAccommodation(ctx, database.DeleteAccommodationParams{
		ID:        accommodation.ID,
		UpdatedAt: utiltime.GetTimeNow(),
	})
	if err != nil {
		return response.ErrCodeDeleteAccommodationFailed, fmt.Errorf("error for delete accommodation: %s", err)
	}

	return response.ErrCodeDeleteAccommodationSuccess, nil
}

func (t *AccommodationImpl) UpdateAccommodation(ctx *gin.Context, in *vo.UpdateAccommodationInput) (codeResult int, out *vo.UpdateAccommodationOutput, err error) {
	out = &vo.UpdateAccommodationOutput{}

	// TODO: get userId from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check manager exists in database
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// TODO: get accommodation in database
	accommodation, err := t.sqlc.GetAccommodationById(ctx, in.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userID {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user is not the owner of the accommodation")
	}

	// TODO: update accommodation
	now := utiltime.GetTimeNow()
	facilitiesJSON, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	rulesJSON, err := json.Marshal(in.Rules)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal rules: %s", err)
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
		return response.ErrCodeUpdateAccommodationFailed, nil, fmt.Errorf("error for update accommodation: %s", err)
	}

	// TODO: get facility
	for _, facilityID := range in.Facilities {
		facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
		if err != nil {
			return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
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

	return response.ErrCodeUpdateAccommodationSuccess, out, nil
}

func (t *AccommodationImpl) GetAccommodations(ctx context.Context) (codeStatus int, out []*vo.GetAccommodations, err error) {
	out = []*vo.GetAccommodations{}

	accommodations, err := t.sqlc.GetAccommodations(ctx)
	if err != nil {
		return response.ErrCodeGetAccommodationsFailed, nil, fmt.Errorf("error for get accommodations: %s", err)
	}

	for _, accommodation := range accommodations {
		// TODO: get facility
		var facilityIDs []string
		if err := json.Unmarshal(accommodation.Facilities, &facilityIDs); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		facilities := []vo.FacilitiesOutput{}

		for _, facilityID := range facilityIDs {
			facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
			if err != nil {
				return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
			}

			facilities = append(facilities, vo.FacilitiesOutput{
				ID:    facility.ID,
				Name:  facility.Name,
				Image: facility.Image,
			})
		}

		rules := vo.Rule{}
		if err := json.Unmarshal(accommodation.Rules, &rules); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling property surroundings: %s", err)
		}

		// TODO: get images of accommodation
		accommodationImages, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		var imagePaths []string
		for _, i := range accommodationImages {
			imagePaths = append(imagePaths, i.Image)
		}

		out = append(out, &vo.GetAccommodations{
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
		})
	}
	return response.ErrCodeGetAccommodationSuccess, out, nil
}

func (t *AccommodationImpl) CreateAccommodation(ctx *gin.Context, in *vo.CreateAccommodationInput) (codeResult int, out *vo.CreateAccommodationOutput, err error) {
	out = &vo.CreateAccommodationOutput{}

	// TODO: get userId from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check manager exists in database
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeGetManagerFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// TODO: convert struct to json
	now := utiltime.GetTimeNow()
	id := uuid.New().String()

	facilitiesJSON, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	rulesJSON, err := json.Marshal(in.Rules)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal rules: %s", err)
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
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for create accommodation: %s", err)
	}

	// TODO: get accommodation
	accommodation, err := t.sqlc.GetAccommodationById(ctx, id)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("get accommodation failed: %s", err)
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

	// TODO: get facility
	var facilitieIDs []string
	if err := json.Unmarshal(accommodation.Facilities, &facilitieIDs); err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
	}

	for _, facilityID := range facilitieIDs {
		facility, err := t.sqlc.GetAccommodationFacilityById(ctx, facilityID)
		if err != nil {
			return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
		}

		out.Facilities = append(out.Facilities, vo.FacilitiesOutput{
			ID:    facility.ID,
			Name:  facility.Name,
			Image: facility.Image,
		})
	}

	err = json.Unmarshal(accommodation.Rules, &out.Rules)
	if err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error for unmarshal rules: %s", err)
	}

	// TODO: get images of accommodation
	accommodationImages, err := t.sqlc.GetAccommodationImages(ctx, accommodation.ID)
	if err != nil {
		return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
	}

	for _, i := range accommodationImages {
		out.Images = append(out.Images, i.Image)
	}

	return response.ErrCodeCreateAccommodationSuccess, out, nil
}

func NewAccommodationImpl(sqlc *database.Queries) *AccommodationImpl {
	return &AccommodationImpl{
		sqlc: sqlc,
	}
}
