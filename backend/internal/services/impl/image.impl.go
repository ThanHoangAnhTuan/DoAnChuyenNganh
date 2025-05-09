package impl

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

type ImageImpl struct {
	sqlc *database.Queries
}

// DeleteImage implements services.IImage.
func (i *ImageImpl) DeleteImage(ctx *gin.Context, fileName string) (err error) {
	panic("unimplemented")
}
func (i *ImageImpl) UploadImages(ctx *gin.Context, in *vo.UploadImages) (codeStatus int, savedImagePaths []string, err error) {
	// TODO: check accommodation exists in db
	if in.Accommodation != "" {
		isExists, err := i.sqlc.CheckAccommodationExists(ctx, in.Accommodation)
		if err != nil {
			return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("get accommodation failed: %s", err)
		}
		if !isExists {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}

		// TODO: Get all image of accommodation
		accommodationImages, err := i.sqlc.GetAccommodationImages(ctx, in.Accommodation)
		if err != nil {
			return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		// TODO: Remove old image
		if len(in.DeleteImages) > 0 {
			is_deleted := false
			deleteFileNames := []string{}

			for _, deteleImage := range in.DeleteImages {
				for _, accommodationImage := range accommodationImages {
					if deteleImage == accommodationImage.Image {
						is_deleted = true
						break
					}
				}

				if is_deleted {
					err := i.sqlc.DeleteAccommodationImage(ctx, deteleImage)
					if err != nil {
						return response.ErrCodeDeleteAccommodationImagesFailed, nil, fmt.Errorf("delete images in db of accommodation failed: %s", err)
					}
					deleteFileNames = append(deleteFileNames, deteleImage)
				}
			}

			err = deleteImageToDisk(deleteFileNames)
			if err != nil {
				return response.ErrCodeDeleteAccommodationImagesFailed, nil, fmt.Errorf("delete images in disk of accommodation failed: %s", err)
			}
		}

		// TODO: Save image to disk
		if len(in.Images) > 0 {
			codeStatus, imagesFileName, err := saveImageToDisk(ctx, in.Images)
			if err != nil {
				return codeStatus, nil, err
			}
			// TODO: Save image to db
			for _, image := range imagesFileName {
				id := uuid.New().String()
				now := utiltime.GetTimeNow()
				err := i.sqlc.UpdateAccommodationImages(ctx, database.UpdateAccommodationImagesParams{
					ID:              id,
					AccommodationID: in.Accommodation,
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
		accommodationImages, err = i.sqlc.GetAccommodationImages(ctx, in.Accommodation)
		if err != nil {
			return response.ErrCodeGetAccommodationImagesFailed, nil, fmt.Errorf("get images of accommodation failed: %s", err)
		}

		for _, i := range accommodationImages {
			savedImagePaths = append(savedImagePaths, i.Image)
		}

	} else if in.AccommodationDetail != "" {
		isExists, err := i.sqlc.CheckAccommodationDetailExists(ctx, in.AccommodationDetail)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailFailed, nil, fmt.Errorf("get accommodation detail failed: %s", err)
		}
		if !isExists {
			return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("accommodation detail not found")
		}

		// TODO: Get all image of accommodation detail
		accommodationDetailImages, err := i.sqlc.GetAccommodationDetailImages(ctx, in.AccommodationDetail)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailImagesFailed, nil, fmt.Errorf("get images of accommodation detail failed: %s", err)
		}

		// TODO: Remove old image
		is_deleted := false
		deleteFileNames := []string{}

		for _, deteleImage := range in.DeleteImages {
			for _, accommodationDetailImage := range accommodationDetailImages {
				if deteleImage == accommodationDetailImage.Image {
					is_deleted = true
					break
				}
			}

			if is_deleted {
				err := i.sqlc.DeleteAccommodationDetailImage(ctx, deteleImage)
				if err != nil {
					return response.ErrCodeDeleteAccommodationDetailImagesFailed, nil, fmt.Errorf("delete images in db of accommodation detail failed: %s", err)
				}
				deleteFileNames = append(deleteFileNames, deteleImage)
			}
		}

		err = deleteImageToDisk(deleteFileNames)
		if err != nil {
			return response.ErrCodeDeleteAccommodationDetailImagesFailed, nil, fmt.Errorf("delete images in disk of accommodation detail failed: %s", err)
		}

		// TODO: Save image to disk
		codeStatus, imagesFileName, err := saveImageToDisk(ctx, in.Images)
		if err != nil {
			return codeStatus, nil, err
		}

		// TODO: Save image to db
		for _, image := range imagesFileName {
			id := uuid.New().String()
			now := utiltime.GetTimeNow()
			err := i.sqlc.UpdateAccommodationDetailImages(ctx, database.UpdateAccommodationDetailImagesParams{
				ID:                    id,
				AccommodationDetailID: in.AccommodationDetail,
				Image:                 image,
				CreatedAt:             now,
				UpdatedAt:             now,
			})
			if err != nil {
				return response.ErrCodeSaveAccommodationDetailImagesFailed, nil, fmt.Errorf("save images in db of accommodation detail failed: %s", err)
			}
		}

		// TODO: Get all image
		accommodationDetailImages, err = i.sqlc.GetAccommodationDetailImages(ctx, in.AccommodationDetail)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailImagesFailed, nil, fmt.Errorf("get images of accommodation detail failed: %s", err)
		}

		for _, i := range accommodationDetailImages {
			savedImagePaths = append(savedImagePaths, i.Image)
		}
	}
	return response.ErrCodeUploadFileSuccess, savedImagePaths, nil
}

func saveImageToDisk(ctx *gin.Context, images []*multipart.FileHeader) (codeStatus int, savedImagePaths []string, err error) {
	uploadDir := "storage/uploads"

	// TODO: Make sure the directory exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return response.ErrCodeCreateFolderFailed, nil, fmt.Errorf("cannot create upload directory: %s", err)
		}
	}

	for _, file := range images {
		// TODO: create unique file name
		fileName := strconv.Itoa(int(utiltime.GetTimeNow())) + uuid.New().String()
		fileName += filepath.Ext(file.Filename)

		// TODO: create path
		savePath := filepath.Join(uploadDir, fileName)

		// TODO: save file
		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			return response.ErrCodeUploadFileFailed, nil, fmt.Errorf("error upload images: %s", err)
		}
		savedImagePaths = append(savedImagePaths, fileName)
	}

	return response.ErrCodeUploadFileSuccess, savedImagePaths, nil
}

func deleteImageToDisk(fileNames []string) error {
	uploadDir := "storage/uploads"
	var failedDeletes []string

	for _, name := range fileNames {
		imagePath := filepath.Join(uploadDir, name)
		if err := os.Remove(imagePath); err != nil {
			failedDeletes = append(failedDeletes, name)
			continue
		}
	}

	if len(failedDeletes) > 0 {
		return fmt.Errorf("failed to delete images: %v", failedDeletes)
	}

	return nil
}

func NewImageImpl(sqlc *database.Queries) *ImageImpl {
	return &ImageImpl{
		sqlc: sqlc,
	}
}
