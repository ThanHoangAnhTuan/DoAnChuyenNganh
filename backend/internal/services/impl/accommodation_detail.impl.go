package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

type AccommodationDetailImpl struct {
	sqlc *database.Queries
}

func (a *AccommodationDetailImpl) CreateAccommodationDetail(ctx *gin.Context, in *vo.CreateAccommodationDetailInput) (codeStatus int, out *vo.CreateAccommodationDetailOutput, err error) {
	out = &vo.CreateAccommodationDetailOutput{}

	// TODO: get userId from context
	userId, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userId not found in context")
	}

	// TODO: check manager exists
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userId)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if manager == 0 {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// TODO: check accommodation exists
	accommodation, err := a.sqlc.GetAccommodationById(ctx, in.AccommodationId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userId {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user is not the owner of the accommodation")
	}

	// TODO: create accommodation detail
	accommodationDetailId := uuid.New().String()
	bedsJson, err := json.Marshal(in.Beds)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	facilitiesJson, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	now := utiltime.GetTimeNow()

	params := database.CreateAccommodationDetailParams{
		ID:              accommodationDetailId,
		AccommodationID: accommodation.ID,
		Name:            in.Name,
		Guests:          in.Guests,
		AvailableRooms:  in.AvailableRooms,
		Price:           in.Price,
		Beds:            bedsJson,
		Facilities:      facilitiesJson,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	err = a.sqlc.CreateAccommodationDetail(ctx, params)
	if err != nil {
		return response.ErrCodeCreateAccommodationDetailFailed, nil, fmt.Errorf("error for create accommodation details: %s", err)
	}

	// TODO: get facility
	var facilitieIds []vo.FacilitiesInput
	if err := json.Unmarshal(facilitiesJson, &facilitieIds); err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
	}

	for _, facilityId := range facilitieIds {
		facility, err := a.sqlc.GetAccommodationFacilityById(ctx, facilityId.Id)
		if err != nil {
			return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
		}

		out.Facilities = append(out.Facilities, vo.FacilitiesOutput{
			Name:  facility.Name,
			Image: facility.Image,
		})
	}

	out.Id = accommodationDetailId
	out.AccommodationId = in.AccommodationId
	out.AvailableRooms = in.AvailableRooms
	out.Beds = in.Beds
	out.DiscountId = in.DiscountId
	out.Guests = in.Guests
	out.Name = in.Name
	out.Price = in.Price
	return response.ErrCodeCreateAccommodationDetailSuccess, out, nil
}

func (a *AccommodationDetailImpl) DeleteAccommodationDetail(ctx context.Context, in *vo.DeleteAccommodationDetailInput) (codeResult int, err error) {
	// TODO: Check the user is manager
	userId, ok := utils.GetUserIDFromContext(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userId not found in context")
	}

	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userId)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, fmt.Errorf("error for get manager: %s", err)
	}

	if manager == 0 {
		return response.ErrCodeManagerNotFound, fmt.Errorf("manager not found")
	}

	// TODO: check the accommodation detail exists
	isExists, err := a.sqlc.CheckAccommodationDetailExists(ctx, in.Id)
	if err != nil {
		return response.ErrCodeAccommodationDetailNotFound, fmt.Errorf("error for get accommodation detail: %s", err)
	}

	if !isExists {
		return response.ErrCodeAccommodationDetailNotFound, nil
	}

	// TODO: check the accommodation detail belongs to manager
	isBelongs, err := a.sqlc.IsAccommodationDetailBelongsToManager(ctx, database.IsAccommodationDetailBelongsToManagerParams{
		ID:   userId,
		ID_2: in.Id,
	})
	if err != nil {
		return response.ErrCodeDeleteAccommodationDetailFailed, fmt.Errorf("error for delete accommodation detail: %s", err)
	}
	if !isBelongs {
		return response.ErrCodeForbidden, fmt.Errorf("error for do not have permission to delete accommodation detail")
	}

	// TODO: delete accommodation detail
	err = a.sqlc.DeleteAccommodationDetail(ctx, in.Id)
	if err != nil {
		return response.ErrCodeDeleteAccommodationDetailFailed, fmt.Errorf("error for delete accommodation detail: %s", err)
	}
	return response.ErrCodeDeleteAccommodationDetailSuccess, nil
}

func (a *AccommodationDetailImpl) GetAccommodationDetails(ctx context.Context, in *vo.GetAccommodationDetailsInput) (codeStatus int, out []*vo.GetAccommodationDetailsOutput, err error) {
	out = []*vo.GetAccommodationDetailsOutput{}

	// TODO: get accommodation details by accommodation id
	accommodationDetails, err := a.sqlc.GetAccommodationDetails(ctx, in.AccommodationId)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation by id failed: %s", err)
	}

	for _, accommodationDetail := range accommodationDetails {
		beds := vo.Beds{}
		if err := json.Unmarshal(accommodationDetail.Beds, &beds); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling beds: %s", err)
		}

		// TODO: get facility
		var facilityIds []vo.FacilitiesInput
		if err := json.Unmarshal(accommodationDetail.Facilities, &facilityIds); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		facilities := []vo.FacilitiesOutput{}

		for _, facilityId := range facilityIds {
			facility, err := a.sqlc.GetAccommodationFacilityById(ctx, facilityId.Id)
			if err != nil {
				return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
			}

			facilities = append(facilities, vo.FacilitiesOutput{
				Name:  facility.Name,
				Image: facility.Image,
			})
		}

		// TODO: get images of accommodation detail
		accommodationDetailImages, err := a.sqlc.GetAccommodationDetailImages(ctx, accommodationDetail.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		var pathNames []string
		for _, img := range accommodationDetailImages {
			pathNames = append(pathNames, img.Image)
		}

		out = append(out, &vo.GetAccommodationDetailsOutput{
			Id:              accommodationDetail.ID,
			AccommodationId: accommodationDetail.AccommodationID,
			Name:            accommodationDetail.Name,
			Guests:          accommodationDetail.Guests,
			Beds:            beds,
			Facilities:      facilities,
			AvailableRooms:  accommodationDetail.AvailableRooms,
			Price:           accommodationDetail.Price,
			DiscountId:      accommodationDetail.DiscountID.String,
			Images:          pathNames,
		})
	}
	return response.ErrCodeGetAccommodationDetailsSuccess, out, nil
}

func (a *AccommodationDetailImpl) UpdateAccommodationDetail(ctx *gin.Context, in *vo.UpdateAccommodationDetailInput) (codeResult int, out *vo.UpdateAccommodationDetailOutput, err error) {
	out = &vo.UpdateAccommodationDetailOutput{}

	// TODO: Check the user is manager
	userId, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userId not found in context")
	}

	// TODO: check manager exists
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userId)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if manager == 0 {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// TODO: check the accommodation detail exists
	isExists, err := a.sqlc.CheckAccommodationDetailExists(ctx, in.Id)
	if err != nil {
		return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("error for get accommodation detail: %s", err)
	}

	if !isExists {
		return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("get accommodation detail not found")
	}

	// TODO: Check the user is the owner of the accommodation detail
	isBelongs, err := a.sqlc.IsAccommodationDetailBelongsToManager(ctx, database.IsAccommodationDetailBelongsToManagerParams{
		ID:   userId,
		ID_2: in.Id,
	})
	if err != nil {
		return response.ErrCodeUpdateAccommodationDetailFailed, nil, fmt.Errorf("error for update accommodation detail: %s", err)
	}
	if !isBelongs {
		return response.ErrCodeForbidden, nil, fmt.Errorf("error for do not have permission to update accommodation detail")
	}

	// TODO: update accommodation detail
	bedsJson, err := json.Marshal(in.Beds)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	facilitiesJson, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	now := utiltime.GetTimeNow()
	err = a.sqlc.UpdateAccommodationDetail(ctx, database.UpdateAccommodationDetailParams{
		Name:            in.Name,
		Guests:          in.Guests,
		Beds:            bedsJson,
		Facilities:      facilitiesJson,
		AvailableRooms:  in.AvailableRooms,
		Price:           in.Price,
		UpdatedAt:       now,
		ID:              in.Id,
		AccommodationID: in.AccommodationId,
	})
	if err != nil {
		return response.ErrCodeUpdateAccommodationDetailFailed, nil, fmt.Errorf("error for update accommodation detail failed: %s", err)
	}

	// TODO: get images of accommodation detail
	accommodationDetailImages, err := a.sqlc.GetAccommodationDetailImages(ctx, in.Id)
	if err != nil {
		return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
	}

	var pathNames []string
	for _, img := range accommodationDetailImages {
		pathNames = append(pathNames, img.Image)
	}

	for _, facilityId := range in.Facilities {
		facility, err := a.sqlc.GetAccommodationFacilityById(ctx, facilityId.Id)
		if err != nil {
			return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
		}
		out.Facilities = append(out.Facilities, vo.FacilitiesOutput{
			Name:  facility.Name,
			Image: facility.Image,
		})
	}

	out.AccommodationId = in.AccommodationId
	out.AvailableRooms = in.AvailableRooms
	out.Beds = in.Beds
	out.DiscountId = in.DiscountId
	out.Guests = in.Guests
	out.Id = in.Id
	out.Name = in.Name
	out.Price = in.Price
	out.Images = pathNames

	return response.ErrCodeUpdateAccommodationDetailSuccess, out, nil
}

func NewAccommodationDetailImpl(sqlc *database.Queries) *AccommodationDetailImpl {
	return &AccommodationDetailImpl{
		sqlc: sqlc,
	}
}
