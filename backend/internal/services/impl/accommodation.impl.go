package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	saveimages "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/save_images"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

type AccommodationImpl struct {
	sqlc *database.Queries
}

// DeleteAccommodation implements services.IAccommodation.
func (t *AccommodationImpl) DeleteAccommodation(ctx context.Context, in *vo.DeleteAccommodationInput) (codeResult int, err error) {
	// !. get userId from context
	// userID := ctx.Value("userId").(string)
	userID := "8ae3a3b6-db41-4dad-9ef2-57386a049c49"
	if userID == "" {
		return response.ErrCodeUnauthorized, fmt.Errorf("userId not found in context")
	}

	// !. check manager exists in database
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, fmt.Errorf("error for get manager: %s", err)
	}

	if manager == 0 {
		return response.ErrCodeManagerNotFound, fmt.Errorf("manager not found")
	}

	// !. check accommodation exists in database
	accommodation, err := t.sqlc.GetAccommodationById(ctx, in.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeGetAccommodationFailed, fmt.Errorf("error for get accommodation: %w", err)
	}

	// !. check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userID {
		return response.ErrCodeUnauthorized, fmt.Errorf("user is not the owner of the accommodation")
	}

	// !. delete accommodation
	err = t.sqlc.DeleteAccommodation(ctx, database.DeleteAccommodationParams{
		ID:        accommodation.ID,
		UpdatedAt: utiltime.GetTimeNow(),
	})
	if err != nil {
		return response.ErrCodeDeleteAccommodationFailed, fmt.Errorf("error for delete accommodation: %s", err)
	}

	return response.ErrCodeDeleteAccommodationSuccess, nil
}

// UpdateAccommodation implements services.IAccommodation.
func (t *AccommodationImpl) UpdateAccommodation(ctx *gin.Context, in *vo.UpdateAccommodationInput) (codeResult int, out *vo.UpdateAccommodationOutput, err error) {
	out = &vo.UpdateAccommodationOutput{}
	// !. get userId from context
	// userID, exists := ctx.Get("userId")
	userID := "8ae3a3b6-db41-4dad-9ef2-57386a049c49"
	// if !exists {
	// 	return response.ErrCodeUnauthorized, nil, fmt.Errorf("userId not found in context")
	// }

	// !. check manager exists in database
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if manager == 0 {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	// !. get accommodation in database
	accommodation, err := t.sqlc.GetAccommodationById(ctx, in.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// !. check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userID {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user is not the owner of the accommodation")
	}

	// !. update accommodation
	now := utiltime.GetTimeNow()
	facilitiesJSON, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	propertySurroundingsJSON, err := json.Marshal(in.PropertySurroundings)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal property surroundings: %s", err)
	}
	var images []string
	if len(in.Image) == 0 {
		images = append(images, accommodation.Image)
	} else {
		images, err = saveimages.SaveImages(ctx, in.Image)
		if err != nil {
			return response.ErrCodeSaveFileFailed, nil, fmt.Errorf("error for save images: %s", err)
		}
	}

	err = t.sqlc.UpdateAccommodation(ctx, database.UpdateAccommodationParams{
		ID:                   accommodation.ID,
		Name:                 in.Name,
		Country:              in.Country,
		City:                 in.City,
		District:             in.District,
		Description:          in.Description,
		Facilities:           facilitiesJSON,
		PropertySurroundings: propertySurroundingsJSON,
		Image:                images[0],
		GgMap:                in.GoogleMap,
		Rules:                in.Rules,
		UpdatedAt:            now,
	})
	if err != nil {
		return response.ErrCodeUpdateAccommodationFailed, nil, fmt.Errorf("error for update accommodation: %s", err)
	}

	// !. return response
	out.Id = accommodation.ID
	out.ManagerId = accommodation.ManagerID
	out.Name = in.Name
	out.City = in.City
	out.Country = in.Country
	out.District = in.District
	out.Description = in.Description

	var facilities vo.Facilities
	if err := json.Unmarshal(accommodation.Facilities, &facilities); err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
	}

	var propertySurroundings vo.PropertySurroundings
	if err := json.Unmarshal(accommodation.PropertySurroundings, &propertySurroundings); err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling property surroundings: %s", err)
	}

	out.Facilities = facilities
	out.GoogleMap = in.GoogleMap
	out.PropertySurroundings = propertySurroundings
	out.Rules = in.Rules
	out.Image = images[0]
	out.Rating = strconv.Itoa(int(accommodation.Rating))

	return response.ErrCodeUpdateAccommodationSuccess, out, nil
}

// GetAccommodations implements services.IAccommodation.
func (t *AccommodationImpl) GetAccommodations(ctx context.Context) (codeStatus int, out []*vo.GetAccommodations, err error) {
	out = make([]*vo.GetAccommodations, 0)
	accommodations, err := t.sqlc.GetAccommodations(ctx)
	if err != nil {
		return response.ErrCodeGetAccommodationsFailed, nil, fmt.Errorf("error for get accommodations: %s", err)
	}

	for _, accommodation := range accommodations {

		facilities := vo.Facilities{}
		if err := json.Unmarshal(accommodation.Facilities, &facilities); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		propertySurroundings := vo.PropertySurroundings{}
		if err := json.Unmarshal(accommodation.PropertySurroundings, &propertySurroundings); err != nil {
			return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error unmarshaling property surroundings: %s", err)
		}

		out = append(out, &vo.GetAccommodations{
			Id:                   accommodation.ID,
			Name:                 accommodation.Name,
			Country:              accommodation.Country,
			City:                 accommodation.City,
			District:             accommodation.District,
			Description:          accommodation.Description,
			Image:                accommodation.Image,
			ManagerId:            accommodation.ManagerID,
			Rating:               strconv.Itoa(int(accommodation.Rating)),
			Facilities:           facilities,
			GoogleMap:            accommodation.GgMap,
			PropertySurroundings: propertySurroundings,
			Rules:                accommodation.Rules,
		})
	}
	return response.ErrCodeGetAccommodationSuccess, out, nil
}

// CreateAccommodation implements services.ITest.
func (t *AccommodationImpl) CreateAccommodation(ctx *gin.Context, in *vo.CreateAccommodationInput) (codeResult int, out *vo.CreateAccommodationOutput, err error) {
	out = &vo.CreateAccommodationOutput{}
	// !1. check manager exists in database
	// userID, exists := ctx.Get("userId")
	userID := "8ae3a3b6-db41-4dad-9ef2-57386a049c49"

	// if !exists {
	// 	return response.ErrCodeUnauthorized, nil, fmt.Errorf("userId not found in context")
	// }
	manager, err := t.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if manager == 0 {
		return response.ErrCodeManagerNotFound, nil, fmt.Errorf("manager not found")
	}

	now := utiltime.GetTimeNow()
	id := uuid.New().String()

	facilitiesJSON, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	propertySurroundingsJSON, err := json.Marshal(in.PropertySurroundings)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("error for marshal property surroundings: %s", err)
	}

	var images []string
	if len(in.Image) > 0 {
		images, err = saveimages.SaveImages(ctx, in.Image)
		if err != nil {
			return response.ErrCodeSaveFileFailed, nil, fmt.Errorf("error for save images: %s", err)
		}
	}

	var imageToUse string
	if len(images) > 0 {
		imageToUse = images[0]
	} else {
		imageToUse = ""
	}

	// !2. create accommodation
	err = t.sqlc.CreateAccommodation(ctx, database.CreateAccommodationParams{
		ID:                   id,
		ManagerID:            userID,
		Name:                 in.Name,
		Country:              in.Country,
		City:                 in.City,
		District:             in.District,
		Description:          in.Description,
		Facilities:           facilitiesJSON,
		PropertySurroundings: propertySurroundingsJSON,
		Image:                imageToUse,
		GgMap:                in.GoogleMap,
		Rules:                in.Rules,
		CreatedAt:            now,
		UpdatedAt:            now,
	})

	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for create accommodation: %s", err)
	}

	out.Id = id
	out.ManagerId = userID
	out.Name = in.Name
	out.City = in.City
	out.Country = in.Country
	out.District = in.District
	out.Description = in.Description
	err = json.Unmarshal(facilitiesJSON, &out.Facilities)
	if err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error for unmarshal facilities: %s", err)
	}
	err = json.Unmarshal(propertySurroundingsJSON, &out.PropertySurroundings)
	if err != nil {
		return response.ErrCodeUnMarshalFailed, nil, fmt.Errorf("error for unmarshal property surroundings: %s", err)
	}
	out.GoogleMap = in.GoogleMap
	out.Rules = in.Rules
	out.Rating = "0"
	out.Image = imageToUse

	return response.ErrCodeCreateAccommodationSuccess, out, nil
}

func NewAccommodationImpl(sqlc *database.Queries) *AccommodationImpl {
	return &AccommodationImpl{
		sqlc: sqlc,
	}
}
