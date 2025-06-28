package accommodationroom

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
	return &serviceImpl{sqlc: sqlc}
}

func (a *serviceImpl) DeleteAccommodationRoom(ctx *gin.Context, in *vo.DeleteAccommodationRoomInput) (codeResult int, err error) {
	// TODO: get userID from gin context
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

	// TODO: check accommodation room belongs to manager
	isBelongs, err := a.sqlc.CheckAccommodationRoomBelongsToManager(ctx, database.CheckAccommodationRoomBelongsToManagerParams{
		ManagerID:           userID,
		AccommodationRoomID: in.ID,
	})

	if err != nil {
		return response.ErrCodeCheckAccommodationRoomBelongsToManagerFailed, fmt.Errorf("check accommodation room belong to manager failed: %s", err)
	}

	if !isBelongs {
		return response.ErrCodeCheckAccommodationRoomNotBelongsToManager, fmt.Errorf("accommodaion room not belongs to manager")
	}

	// TODO: update accommodation room
	err = a.sqlc.DeleteAccommodationRoom(ctx, in.ID)
	if err != nil {
		return response.ErrCodeDeleteAccommodationRoomFailed, fmt.Errorf("delete accommodation room failed: %s", err)
	}
	return response.ErrCodeDeleteAccommodationRoomSuccess, nil
}

func (a *serviceImpl) UpdateAccommodationRoom(ctx *gin.Context, in *vo.UpdateAccommodationRoomInput) (codeResult int, out *vo.UpdateAccommodationRoomOutput, err error) {
	out = &vo.UpdateAccommodationRoomOutput{}

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

	fmt.Printf("managerID: %v", userID)
	fmt.Printf("AccommodationRoomID: %v", in.ID)

	// TODO: check accommodation room belongs to manager
	isBelongs, err := a.sqlc.CheckAccommodationRoomBelongsToManager(ctx, database.CheckAccommodationRoomBelongsToManagerParams{
		ManagerID:           userID,
		AccommodationRoomID: in.ID,
	})

	if err != nil {
		return response.ErrCodeCheckAccommodationRoomBelongsToManagerFailed, nil, fmt.Errorf("check accommodation room belong to manager failed: %s", err)
	}

	if !isBelongs {
		return response.ErrCodeCheckAccommodationRoomNotBelongsToManager, nil, fmt.Errorf("accommodaion room not belongs to manager")
	}

	// TODO: update accommodation room
	now := utiltime.GetTimeNow()
	err = a.sqlc.UpdateAccommodationRooms(ctx, database.UpdateAccommodationRoomsParams{
		Name:      in.Name,
		Status:    database.EcommerceGoAccommodationRoomStatus(in.Status),
		UpdatedAt: now,
		ID:        in.ID,
	})

	if err != nil {
		return response.ErrCodeUpdateAccommodationRoomFailed, nil, fmt.Errorf("update accommodation room failed: %s", err)
	}

	out.ID = in.ID
	out.Name = in.Name
	out.Status = string(in.Status)
	return response.ErrCodeUpdateAccommodationRoomSuccess, out, nil
}

func (a *serviceImpl) GetAccommodationRooms(ctx *gin.Context, in *vo.GetAccommodationRoomsInput) (codeStatus int, out []*vo.GetAccommodationRoomsOutput, err error) {
	out = []*vo.GetAccommodationRoomsOutput{}

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

	// TODO: check accommodation type of manager
	isBelongs, err := a.sqlc.CheckAccommodationTypeBelongToManager(ctx, database.CheckAccommodationTypeBelongToManagerParams{
		ManagerID:           userID,
		AccommodationTypeID: in.AccommodationTypeID,
	})

	if err != nil {
		return response.ErrCodeCheckAccommodationTypeBelongsToManagerFailed, nil, fmt.Errorf("check accommodation type belong to manager failed: %s", err)
	}

	if !isBelongs {
		return response.ErrCodeCheckAccommodationTypeNotBelongsToManager, nil, fmt.Errorf("accommodaion type not belongs to manager")
	}

	// TODO: get accommodation room
	accommodationRooms, err := a.sqlc.GetAccommodationRooms(ctx, in.AccommodationTypeID)
	if err != nil {
		return response.ErrCodeGetAccommodationRoomFailed, nil, fmt.Errorf("get accommodation room failed: %s", err)
	}

	for _, accommodationRoom := range accommodationRooms {
		out = append(out, &vo.GetAccommodationRoomsOutput{
			ID: accommodationRoom.ID,
			// AccommodationTypeID: accommodationRoom.AccommodationType,
			Name:   accommodationRoom.Name,
			Status: string(accommodationRoom.Status),
		})
	}
	return response.ErrCodeGetAccommodationRoomSuccess, out, nil
}

func (a *serviceImpl) CreateAccommodationRoom(ctx *gin.Context, in *vo.CreateAccommodationRoomInput) (codeStatus int, out []*vo.CreateAccommodationRoomOutput, err error) {
	out = []*vo.CreateAccommodationRoomOutput{}

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

	// TODO: check accommodation type of manager
	isBelongs, err := a.sqlc.CheckAccommodationTypeBelongToManager(ctx, database.CheckAccommodationTypeBelongToManagerParams{
		ManagerID:           userID,
		AccommodationTypeID: in.AccommodationTypeID,
	})

	if err != nil {
		return response.ErrCodeCheckAccommodationTypeBelongsToManagerFailed, nil, fmt.Errorf("check accommodation type belong to manager failed: %s", err)
	}

	if !isBelongs {
		return response.ErrCodeCheckAccommodationTypeNotBelongsToManager, nil, fmt.Errorf("accommodaion type not belongs to manager")
	}

	for i := range in.Quantity {
		id := uuid.NewString()
		now := utiltime.GetTimeNow()
		err := a.sqlc.CreateAccommodationRoom(ctx, database.CreateAccommodationRoomParams{
			ID:                id,
			AccommodationType: in.AccommodationTypeID,
			Name:              fmt.Sprintf("%s-%02d", in.Prefix, i+1),
			Status:            database.EcommerceGoAccommodationRoomStatusAvailable,
			CreatedAt:         now,
			UpdatedAt:         now,
		})

		if err != nil {
			return response.ErrCodeCreateAccommodationRoomFailed, nil, fmt.Errorf("create accommodation room failed: %s", err)
		}
	}

	// TODO: get accommodation room
	accommodationRooms, err := a.sqlc.GetAccommodationRooms(ctx, in.AccommodationTypeID)
	if err != nil {
		return response.ErrCodeGetAccommodationRoomFailed, nil, fmt.Errorf("get accommodation room failed: %s", err)
	}

	for _, accommodationRoom := range accommodationRooms {
		out = append(out, &vo.CreateAccommodationRoomOutput{
			ID: accommodationRoom.ID,
			// AccommodationTypeID: accommodationRoom.AccommodationType,
			Name:   accommodationRoom.Name,
			Status: string(accommodationRoom.Status),
		})
	}

	return response.ErrCodeCreateAccommodationRoomSuccess, out, nil
}
