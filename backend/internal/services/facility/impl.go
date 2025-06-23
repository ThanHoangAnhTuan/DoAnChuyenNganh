package facility

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/uploader"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{sqlc: sqlc}
}

func (f *serviceImpl) DeleteFacility(ctx *gin.Context, in *vo.DeleteFacilityInput) (codeStatus int, err error) {
	// TODO: get userID from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is admin
	isExists, err := f.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeGetAdminFailed, fmt.Errorf("get admin failed")
	}

	if !isExists {
		return response.ErrCodeUnauthorized, fmt.Errorf("user not admin")
	}

	// TODO: get facility
	facility, err := f.sqlc.GetAccommodationFacilityById(ctx, in.ID)
	if err != nil {
		return response.ErrCodeGetFacilityFailed, fmt.Errorf("get facility failed: %s", err)
	}

	// TODO: delete image
	err = uploader.DeleteImageToDisk([]string{facility.Image})
	if err != nil {
		return response.ErrCodeDeleteFacilityImageFailed, fmt.Errorf("delete images in disk of facitliy failed: %s", err)
	}

	err = f.sqlc.DeleteFacility(ctx, facility.ID)
	if err != nil {
		return response.ErrCodeDeleteFacilityFailed, fmt.Errorf("delete facility failed: %s", err)
	}
	return response.ErrCodeDeleteFacilitySuccess, nil
}

func (f *serviceImpl) UpdateFacility(ctx *gin.Context, in *vo.UpdateFacilityInput) (codeStatus int, out *vo.UpdateFacilityOutput, err error) {
	out = &vo.UpdateFacilityOutput{}

	// TODO: get userID from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is admin
	isExists, err := f.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeGetAdminFailed, nil, fmt.Errorf("get admin failed")
	}

	if !isExists {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user not admin")
	}

	// TODO: get facility
	facility, err := f.sqlc.GetAccommodationFacilityById(ctx, in.ID)
	if err != nil {
		return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
	}

	// TODO:
	if in.Image != nil {
		// TODO: delete image
		err = uploader.DeleteImageToDisk([]string{facility.Image})
		if err != nil {
			return response.ErrCodeDeleteFacilityImageFailed, nil, fmt.Errorf("delete images in disk of facitliy failed: %s", err)
		}

		// TODO: save image to disk
		codeStatus, imagesFileName, err := uploader.SaveImageToDisk(ctx, []*multipart.FileHeader{in.Image})
		if err != nil {
			return codeStatus, nil, err
		}
		// TODO: save facility
		now := utiltime.GetTimeNow()
		err = f.sqlc.UpdateFacility(ctx, database.UpdateFacilityParams{
			Name:      in.Name,
			Image:     imagesFileName[0],
			UpdatedAt: now,
			ID:        in.ID,
		})
		if err != nil {
			return response.ErrCodeUpdateFacilityFailed, nil, fmt.Errorf("update facility failed: %s", err)
		}
	} else {
		now := utiltime.GetTimeNow()
		err = f.sqlc.UpdateNameFacility(ctx, database.UpdateNameFacilityParams{
			ID:        in.ID,
			Name:      in.Name,
			UpdatedAt: now,
		})
		if err != nil {
			return response.ErrCodeUpdateFacilityFailed, nil, fmt.Errorf("update facility failed: %s", err)
		}
	}

	// TODO: get facility
	facility, err = f.sqlc.GetAccommodationFacilityById(ctx, in.ID)
	if err != nil {
		return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
	}

	out.ID = facility.ID
	out.Image = facility.Image
	out.Name = facility.Name

	return response.ErrCodeUpdateFacilitySuccess, out, nil
}

func (f *serviceImpl) CreateFacility(ctx *gin.Context, in *vo.CreateFacilityInput) (codeStatus int, out *vo.CreateFacilityOutput, err error) {
	out = &vo.CreateFacilityOutput{}

	// TODO: get userID from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
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
	codeStatus, imagesFileName, err := uploader.SaveImageToDisk(ctx, []*multipart.FileHeader{in.Image})
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

	out.ID = facility.ID
	out.Image = facility.Image
	out.Name = facility.Name

	return response.ErrCodeCreateFacilitySuccess, out, nil
}

func (f *serviceImpl) GetFacilities(ctx *gin.Context) (codeStatus int, out []*vo.GetFacilitiesOutput, err error) {
	// TODO: get user id in gin.Context
	// val := ctx.Value("userId")
	// if val == nil {
	// 	return response.ErrCodeUnauthorized, nil, fmt.Errorf("unauthorized")
	// }
	// userID, ok := val.(string)
	// if !ok {
	// 	return response.ErrCodeUnauthorized, nil, fmt.Errorf("invalid user id format")
	// }

	// TODO: check user is admin
	// isExists, err := f.sqlc.CheckUserAdminExistsById(ctx, userID)
	// if err != nil {
	// 	return response.ErrCodeGetAdminFailed, nil, fmt.Errorf("get admin failed")
	// }

	// if !isExists {
	// 	return response.ErrCodeUnauthorized, nil, fmt.Errorf("user not admin")
	// }

	// TODO: get facilities
	facilities, err := f.sqlc.GetAccommodationFacilityNames(ctx)
	if err != nil {
		return response.ErrCodeGetFacilityFailed, nil, fmt.Errorf("get facility failed: %s", err)
	}

	for _, facility := range facilities {
		out = append(out, &vo.GetFacilitiesOutput{
			ID:    facility.ID,
			Name:  facility.Name,
			Image: facility.Image,
		})
	}

	return response.ErrCodeGetFacilitySuccess, out, nil
}
