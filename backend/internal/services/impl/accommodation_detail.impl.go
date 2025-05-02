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
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/services"
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
	// !. get userId from context
	userId, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userId not found in context")
	}

	// !. check manager exists
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userId)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if manager == 0 {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// !. check accommodation exists
	accommodation, err := a.sqlc.GetAccommodationById(ctx, in.AccommodationId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// !. check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userId {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user is not the owner of the accommodation")
	}

	// !. create accommodation detail
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
		Beds:            bedsJson,
		Facilities:      facilitiesJson,
		AvailableRooms:  in.AvailableRooms,
		Price:           in.Price,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	err = a.sqlc.CreateAccommodationDetail(ctx, params)
	if err != nil {
		return response.ErrCodeCreateAccommodationDetailFailed, nil, fmt.Errorf("error for create accommodation details: %s", err)
	}

	// TODO: save image
	saveImagePaths, err := services.Image().UploadImages(ctx, in.Images)
	if err != nil {
		return response.ErrCodeSaveFileFailed, nil, fmt.Errorf("error for save image to disk failed: %s", err)
	}

	for _, path := range saveImagePaths {
		imageId := uuid.New().String()
		now := utiltime.GetTimeNow()
		err := a.sqlc.SaveImage(ctx, database.SaveImageParams{
			ID:                    imageId,
			AccommodationDetailID: accommodationDetailId,
			Image:                 path,
			CreatedAt:             now,
			UpdatedAt:             now,
		})
		if err != nil {
			return response.ErrCodeSaveFileFailed, nil, fmt.Errorf("error for save image to db failed: %s", err)
		}
	}

	out.Id = accommodationDetailId
	out.AccommodationId = in.AccommodationId
	out.AvailableRooms = in.AvailableRooms
	out.Beds = in.Beds
	out.Images = saveImagePaths
	out.DiscountId = in.DiscountId
	out.Facilities = in.Facilities
	out.Guests = in.Guests
	out.Name = in.Name
	out.Price = in.Price
	return response.ErrCodeCreateAccommodationDetailSuccess, out, nil
}

func (a *AccommodationDetailImpl) DeleteAccommodationDetail(ctx context.Context, in *vo.DeleteAccommodationDetailInput) (codeResult int, err error) {
	// !. Check the user is manager
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

	// !. check the accommodation detail exists
	isExists, err := a.sqlc.CheckAccommodationDetailExists(ctx, in.Id)
	if err != nil {
		return response.ErrCodeAccommodationDetailNotFound, fmt.Errorf("error for get accommodation detail: %s", err)
	}

	if !isExists {
		return response.ErrCodeAccommodationDetailNotFound, nil
	}

	// !. check the accommodation detail belongs to manager
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

	err = a.sqlc.DeleteAccommodationDetails(ctx, in.Id)
	if err != nil {
		return response.ErrCodeDeleteAccommodationDetailFailed, fmt.Errorf("error for delete accommodation detail: %s", err)
	}
	return response.ErrCodeDeleteAccommodationDetailSuccess, nil
}

func (a *AccommodationDetailImpl) GetAccommodationDetails(ctx context.Context, in *vo.GetAccommodationDetailsInput) (codeStatus int, out []*vo.GetAccommodationDetailsOutput, err error) {
	out = []*vo.GetAccommodationDetailsOutput{}

	// !. get accommodation details by accommodation id
	accommodationDetails, err := a.sqlc.GetAccommodationDetails(ctx, in.AccommodationId)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation by id failed: %s", err)
	}

	for _, accommodationDetail := range accommodationDetails {
		beds := vo.Beds{}
		if err := json.Unmarshal(accommodationDetail.Beds, &beds); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling beds: %s", err)
		}

		facilities := vo.FacilitiesAccommodationDetail{}
		if err := json.Unmarshal(accommodationDetail.Facilities, &facilities); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		// TODO: get images
		images, err := a.sqlc.GetImages(ctx, accommodationDetail.ID)
		if err != nil {
			return response.ErrCodeGetFilesFailed, nil, fmt.Errorf("error for get images failed: %s", err)
		}

		var imagesName []string

		for _, imageName := range images {
			imagesName = append(imagesName, imageName.Image)
		}

		out = append(out, &vo.GetAccommodationDetailsOutput{
			Id:              accommodationDetail.ID,
			AccommodationId: accommodationDetail.AccommodationID,
			Name:            accommodationDetail.Name,
			Images:          imagesName,
			Guests:          accommodationDetail.Guests,
			Beds:            beds,
			Facilities:      facilities,
			AvailableRooms:  accommodationDetail.AvailableRooms,
			Price:           accommodationDetail.Price,
		})
	}
	return response.ErrCodeGetAccommodationDetailsSuccess, out, nil
}

func (a *AccommodationDetailImpl) UpdateAccommodationDetail(ctx *gin.Context, in *vo.UpdateAccommodationDetailInput) (codeResult int, out *vo.UpdateAccommodationDetailOutput, err error) {
	out = &vo.UpdateAccommodationDetailOutput{}

	// !. Check the user is manager
	userId, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userId not found in context")
	}

	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userId)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if manager == 0 {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// !. check the accommodation detail exists
	isExists, err := a.sqlc.CheckAccommodationDetailExists(ctx, in.Id)
	if err != nil {
		return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("error for get accommodation detail: %s", err)
	}

	if !isExists {
		return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("get accommodation detail not found")
	}

	// !. Check the user is the owner of the accommodation detail
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

	// !. update accommodation detail
	bedsJson, err := json.Marshal(in.Beds)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	facilitiesJson, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	now := utiltime.GetTimeNow()

	params := database.UpdateAccommodationDetailParams{
		Name:            in.Name,
		Guests:          in.Guests,
		Beds:            bedsJson,
		Facilities:      facilitiesJson,
		AvailableRooms:  in.AvailableRooms,
		Price:           in.Price,
		UpdatedAt:       now,
		ID:              in.Id,
		AccommodationID: in.AccommodationId,
	}
	err = a.sqlc.UpdateAccommodationDetail(ctx, params)
	if err != nil {
		return response.ErrCodeUpdateAccommodationDetailFailed, nil, fmt.Errorf("error for update accommodation detail failed: %s", err)
	}

	// TODO: update images
	// TODO: get all images of accommodation detail
	images, err := a.sqlc.GetImages(ctx, in.Id)
	if err != nil {
		return response.ErrCodeGetFilesFailed, nil, fmt.Errorf("error for get image failed: %s", err)
	}

	// TODO: delete image
	isDelete := false
	for _, image := range images {
		for _, oldImage := range in.OldImages {
			if oldImage == image.Image {
				isDelete = true
				break
			}
		}
		if !isDelete {
			err := a.sqlc.DeleteImage(ctx, image.ID)
			if err != nil {
				return response.ErrCodeDeleteFileFailed, nil, fmt.Errorf("error for get image failed: %s", err)
			}
		}
	}

	// TODO: add new image
	saveImagePaths, err := services.Image().UploadImages(ctx, in.Images)
	if err != nil {
		return response.ErrCodeSaveFileFailed, nil, fmt.Errorf("error for save image to disk failed: %s", err)
	}

	for _, path := range saveImagePaths {
		imageId := uuid.New().String()
		now := utiltime.GetTimeNow()
		err := a.sqlc.SaveImage(ctx, database.SaveImageParams{
			ID:                    imageId,
			AccommodationDetailID: in.Id,
			Image:                 path,
			CreatedAt:             now,
			UpdatedAt:             now,
		})
		if err != nil {
			return response.ErrCodeSaveFileFailed, nil, fmt.Errorf("error for save image to db failed: %s", err)
		}
	}

	// TODO: return new update images
	updatedImage := append(in.OldImages, saveImagePaths...)

	out.AccommodationId = in.AccommodationId
	out.AvailableRooms = in.AvailableRooms
	out.Beds = in.Beds
	out.Images = updatedImage
	out.DiscountId = in.DiscountId
	out.Facilities = in.Facilities
	out.Guests = in.Guests
	out.Id = in.Id
	out.Name = in.Name
	out.Price = in.Price
	return response.ErrCodeUpdateAccommodationDetailSuccess, out, nil
}

func NewAccommodationDetailImpl(sqlc *database.Queries) *AccommodationDetailImpl {
	return &AccommodationDetailImpl{
		sqlc: sqlc,
	}
}
