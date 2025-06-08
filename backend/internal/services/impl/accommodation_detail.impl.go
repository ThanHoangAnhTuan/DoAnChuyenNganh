package impl

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

type AccommodationDetailImpl struct {
	sqlc *database.Queries
}

func (a *AccommodationDetailImpl) CreateAccommodationDetail(ctx *gin.Context, in *vo.CreateAccommodationDetailInput) (codeStatus int, out *vo.CreateAccommodationDetailOutput, err error) {
	out = &vo.CreateAccommodationDetailOutput{}

	// TODO: get userID from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeGetManagerFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// TODO: check accommodation exists
	accommodation, err := a.sqlc.GetAccommodationById(ctx, in.AccommodationID)
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

	bedsJson, err := json.Marshal(in.Beds)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	facilitiesJson, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	accommodationDetailID := uuid.New().String()
	now := utiltime.GetTimeNow()

	// TODO: create accommodation detail
	err = a.sqlc.CreateAccommodationDetail(ctx, database.CreateAccommodationDetailParams{
		ID:              accommodationDetailID,
		AccommodationID: accommodation.ID,
		Name:            in.Name,
		Guests:          in.Guests,
		AvailableRooms:  in.AvailableRooms,
		Price:           in.Price,
		Beds:            bedsJson,
		Facilities:      facilitiesJson,
		CreatedAt:       now,
		UpdatedAt:       now,
	})
	if err != nil {
		return response.ErrCodeCreateAccommodationDetailFailed, nil, fmt.Errorf("error for create accommodation details: %s", err)
	}

	// TODO: get facility detail
	var facilityIds []string
	if err := json.Unmarshal(facilitiesJson, &facilityIds); err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facility detail: %s", err)
	}

	for _, facilityId := range facilityIds {
		facility, err := a.sqlc.GetAccommodationFacilityDetailById(ctx, facilityId)
		if err != nil {
			return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility detail failed: %s", err)
		}

		out.Facilities = append(out.Facilities, vo.FacilityDetailOutput{
			ID:   facility.ID,
			Name: facility.Name,
		})
	}

	out.ID = accommodationDetailID
	out.AccommodationID = in.AccommodationID
	out.AvailableRooms = in.AvailableRooms
	out.Beds = in.Beds
	out.DiscountID = in.DiscountID
	out.Guests = in.Guests
	out.Name = in.Name
	out.Price = in.Price
	return response.ErrCodeCreateAccommodationDetailSuccess, out, nil
}

func (a *AccommodationDetailImpl) DeleteAccommodationDetail(ctx *gin.Context, in *vo.DeleteAccommodationDetailInput) (codeResult int, err error) {
	// TODO: get user from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeGetManagerFailed, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeManagerNotFound, fmt.Errorf("manager not found")
	}

	// TODO: check the accommodation detail exists
	exists, err := a.sqlc.CheckAccommodationDetailExists(ctx, in.ID)
	if err != nil {
		return response.ErrCodeAccommodationDetailNotFound, fmt.Errorf("error for get accommodation detail: %s", err)
	}

	if !exists {
		return response.ErrCodeAccommodationDetailNotFound, nil
	}

	// TODO: check the accommodation detail belongs to manager
	isBelongs, err := a.sqlc.IsAccommodationDetailBelongsToManager(ctx, database.IsAccommodationDetailBelongsToManagerParams{
		ID:   userID,
		ID_2: in.ID,
	})
	if err != nil {
		return response.ErrCodeDeleteAccommodationDetailFailed, fmt.Errorf("error for delete accommodation detail: %s", err)
	}
	if !isBelongs {
		return response.ErrCodeForbidden, fmt.Errorf("error for do not have permission to delete accommodation detail")
	}

	// TODO: delete accommodation detail
	err = a.sqlc.DeleteAccommodationDetail(ctx, in.ID)
	if err != nil {
		return response.ErrCodeDeleteAccommodationDetailFailed, fmt.Errorf("error for delete accommodation detail: %s", err)
	}
	return response.ErrCodeDeleteAccommodationDetailSuccess, nil
}

func (a *AccommodationDetailImpl) GetAccommodationDetails(ctx *gin.Context, in *vo.GetAccommodationDetailsInput) (codeStatus int, out []*vo.GetAccommodationDetailsOutput, err error) {
	out = []*vo.GetAccommodationDetailsOutput{}

	// TODO: check accommodation exists
	accommodation, err := a.sqlc.GetAccommodationById(ctx, in.AccommodationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: get accommodation details by accommodation id
	accommodationDetails, err := a.sqlc.GetAccommodationDetails(ctx, accommodation.ID)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation by id failed: %s", err)
	}

	for _, accommodationDetail := range accommodationDetails {
		beds := vo.Beds{}
		if err := json.Unmarshal(accommodationDetail.Beds, &beds); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling beds: %s", err)
		}

		// TODO: get facility
		var facilityIDs []string
		if err := json.Unmarshal(accommodationDetail.Facilities, &facilityIDs); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		facilities := []vo.FacilityDetailOutput{}

		for _, facilityID := range facilityIDs {
			facility, err := a.sqlc.GetAccommodationFacilityDetailById(ctx, facilityID)
			if err != nil {
				return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
			}

			facilities = append(facilities, vo.FacilityDetailOutput{
				ID:   facility.ID,
				Name: facility.Name,
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
			ID:              accommodationDetail.ID,
			AccommodationID: accommodationDetail.AccommodationID,
			Name:            accommodationDetail.Name,
			Guests:          accommodationDetail.Guests,
			Beds:            beds,
			Facilities:      facilities,
			AvailableRooms:  accommodationDetail.AvailableRooms,
			Price:           accommodationDetail.Price,
			DiscountID:      accommodationDetail.DiscountID.String,
			Images:          pathNames,
		})
	}
	return response.ErrCodeGetAccommodationDetailsSuccess, out, nil
}

func (a *AccommodationDetailImpl) UpdateAccommodationDetail(ctx *gin.Context, in *vo.UpdateAccommodationDetailInput) (codeResult int, out *vo.UpdateAccommodationDetailOutput, err error) {
	out = &vo.UpdateAccommodationDetailOutput{}

	// TODO: get user from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeGetManagerFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// TODO: check accommodation exists
	accommodation, err := a.sqlc.GetAccommodationById(ctx, in.AccommodationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: check the accommodation detail exists
	isExists, err := a.sqlc.CheckAccommodationDetailExists(ctx, in.ID)
	if err != nil {
		return response.ErrCodeGetAccommodationDetailFailed, nil, fmt.Errorf("error for get accommodation detail: %s", err)
	}

	if !isExists {
		return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("get accommodation detail not found")
	}

	// TODO: Check the user is the owner of the accommodation detail
	isBelongs, err := a.sqlc.IsAccommodationDetailBelongsToManager(ctx, database.IsAccommodationDetailBelongsToManagerParams{
		ID:   accommodation.ManagerID,
		ID_2: in.ID,
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
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal beds: %s", err)
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
		ID:              in.ID,
		AccommodationID: in.AccommodationID,
	})
	if err != nil {
		return response.ErrCodeUpdateAccommodationDetailFailed, nil, fmt.Errorf("error for update accommodation detail failed: %s", err)
	}

	// TODO: get images of accommodation detail
	accommodationDetailImages, err := a.sqlc.GetAccommodationDetailImages(ctx, in.ID)
	if err != nil {
		return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
	}

	var pathNames []string
	for _, img := range accommodationDetailImages {
		pathNames = append(pathNames, img.Image)
	}

	// TODO: get facility of accommodation detail
	for _, facilityID := range in.Facilities {
		facility, err := a.sqlc.GetAccommodationFacilityDetailById(ctx, facilityID)
		if err != nil {
			// TODO: Nếu không tìm thấy facility thì bỏ qua luôn thay vì báo lỗi
			fmt.Printf("Cannot found facility detail: %s", err.Error())
			global.Logger.Error("Cannot found facility detail: ", zap.Error(err))
			break
		}

		out.Facilities = append(out.Facilities, vo.FacilityDetailOutput{
			ID:   facility.ID,
			Name: facility.Name,
		})
	}

	out.AccommodationID = in.AccommodationID
	out.AvailableRooms = in.AvailableRooms
	out.Beds = in.Beds
	out.DiscountID = in.DiscountID
	out.Guests = in.Guests
	out.ID = in.ID
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
