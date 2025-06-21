package upload

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/uploader"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{
		sqlc: sqlc,
	}
}

func (i *serviceImpl) GetImages(ctx *gin.Context, in *vo.GetImagesInput) (codeStatus int, imagesPath []string, err error) {
	// TODO: Get images of accommodation detail
	if in.IsDetail {

		// TODO: Check accommodation detail exists
		isExist, err := i.sqlc.CheckAccommodationDetailExists(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailFailed, imagesPath, fmt.Errorf("get accommodation detail failed: %s", err)
		}

		if !isExist {
			return response.ErrCodeAccommodationDetailNotFound, imagesPath, fmt.Errorf("accommodation detail not found")
		}

		// TODO: get images
		accommodationDetailImages, err := i.sqlc.GetAccommodationDetailImages(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailImagesFailed, imagesPath, fmt.Errorf("get images of accommodation detail failed: %s", err)
		}

		for _, inage := range accommodationDetailImages {
			imagesPath = append(imagesPath, inage.Image)
		}

		return response.ErrCodeGetAccommodationDetailImagesSuccess, imagesPath, nil

	} else {
		// TODO: Check accommodation exists
		isExist, err := i.sqlc.CheckAccommodationExists(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationFailed, imagesPath, fmt.Errorf("get accommodation failed: %s", err)
		}

		if !isExist {
			return response.ErrCodeAccommodationNotFound, imagesPath, fmt.Errorf("accommodation not found")
		}

		// TODO: get images
		accommodationImages, err := i.sqlc.GetAccommodationImages(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationImagesFailed, imagesPath, fmt.Errorf("get images of accommodation  failed: %s", err)
		}

		for _, inage := range accommodationImages {
			imagesPath = append(imagesPath, inage.Image)
		}

		return response.ErrCodeGetAccommodationImagesSuccess, imagesPath, nil
	}
}

func (i *serviceImpl) DeleteImage(ctx *gin.Context, fileName string) (err error) {
	panic("unimplemented")
}

func (i *serviceImpl) UploadImages(ctx *gin.Context, in *vo.UploadImages) (codeStatus int, savedImagePaths []string, err error) {
	// TODO: check accommodation exists in db
	if !in.IsDetail {
		isExists, err := i.sqlc.CheckAccommodationExists(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("get accommodation failed: %s", err)
		}
		if !isExists {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}

		// TODO: Get all image of accommodation
		accommodationImages, err := i.sqlc.GetAccommodationImages(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		// TODO: Remove image
		if len(in.OldImages) > 0 {
			deleteFileNames := []string{}

			for _, accommodationImage := range accommodationImages {
				is_deleted := false
				for _, deteleImage := range in.OldImages {
					if deteleImage == accommodationImage.Image {
						is_deleted = true
						break
					}
				}

				if !is_deleted {
					err := i.sqlc.DeleteAccommodationImage(ctx, accommodationImage.Image)
					if err != nil {
						return response.ErrCodeDeleteAccommodationImagesFailed, nil, fmt.Errorf("delete images in db of accommodation failed: %s", err)
					}
					deleteFileNames = append(deleteFileNames, accommodationImage.Image)
				}
			}

			err = uploader.DeleteImageToDisk(deleteFileNames)
			if err != nil {
				return response.ErrCodeDeleteAccommodationImagesFailed, nil, fmt.Errorf("delete images in disk of accommodation failed: %s", err)
			}
		}

		// TODO: Save image to disk
		if len(in.Images) > 0 {
			codeStatus, imagesFileName, err := uploader.SaveImageToDisk(ctx, in.Images)
			if err != nil {
				return codeStatus, nil, err
			}
			// TODO: Save image to db
			for _, image := range imagesFileName {
				id := uuid.New().String()
				now := utiltime.GetTimeNow()
				err := i.sqlc.UpdateAccommodationImages(ctx, database.UpdateAccommodationImagesParams{
					ID:              id,
					AccommodationID: in.ID,
					Image:           image,
					CreatedAt:       now,
					UpdatedAt:       now,
				})
				if err != nil {
					return response.ErrCodeSaveAccommodationImagesFailed, nil, fmt.Errorf("save images in db of accommodation failed: %s", err)
				}
			}
		}

		// TODO: Get all image
		accommodationImages, err = i.sqlc.GetAccommodationImages(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		for _, i := range accommodationImages {
			savedImagePaths = append(savedImagePaths, i.Image)
		}

	} else {
		isExists, err := i.sqlc.CheckAccommodationDetailExists(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailFailed, nil, fmt.Errorf("get accommodation detail failed: %s", err)
		}
		if !isExists {
			return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("accommodation detail not found")
		}

		// TODO: Get all image of accommodation detail
		accommodationDetailImages, err := i.sqlc.GetAccommodationDetailImages(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailImagesFailed, nil, fmt.Errorf("get images of accommodation detail failed: %s", err)
		}

		// TODO: Remove old image

		deleteFileNames := []string{}

		for _, deteleImage := range in.OldImages {
			is_deleted := false
			for _, accommodationDetailImage := range accommodationDetailImages {
				if deteleImage == accommodationDetailImage.Image {
					is_deleted = true
					break
				}
			}

			if !is_deleted {
				err := i.sqlc.DeleteAccommodationDetailImage(ctx, deteleImage)
				if err != nil {
					return response.ErrCodeDeleteAccommodationDetailImagesFailed, nil, fmt.Errorf("delete images in db of accommodation detail failed: %s", err)
				}
				deleteFileNames = append(deleteFileNames, deteleImage)
			}
		}

		err = uploader.DeleteImageToDisk(deleteFileNames)
		if err != nil {
			return response.ErrCodeDeleteAccommodationDetailImagesFailed, nil, fmt.Errorf("delete images in disk of accommodation detail failed: %s", err)
		}

		// TODO: Save image to disk
		codeStatus, imagesFileName, err := uploader.SaveImageToDisk(ctx, in.Images)
		if err != nil {
			return codeStatus, nil, err
		}

		// TODO: Save image to db
		for _, image := range imagesFileName {
			id := uuid.New().String()
			now := utiltime.GetTimeNow()
			err := i.sqlc.UpdateAccommodationDetailImages(ctx, database.UpdateAccommodationDetailImagesParams{
				ID:                    id,
				AccommodationDetailID: in.ID,
				Image:                 image,
				CreatedAt:             now,
				UpdatedAt:             now,
			})
			if err != nil {
				return response.ErrCodeSaveAccommodationDetailImagesFailed, nil, fmt.Errorf("save images in db of accommodation detail failed: %s", err)
			}
		}

		// TODO: Get all image
		accommodationDetailImages, err = i.sqlc.GetAccommodationDetailImages(ctx, in.ID)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailImagesFailed, nil, fmt.Errorf("get images of accommodation detail failed: %s", err)
		}

		for _, i := range accommodationDetailImages {
			savedImagePaths = append(savedImagePaths, i.Image)
		}
	}
	return response.ErrCodeUploadFileSuccess, savedImagePaths, nil
}
