package manager

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

func (s *serviceImpl) GetManagers(ctx *gin.Context, in *vo.GetManagerInput) (codeStatus int, out []*vo.GetManagerOutput, pagination *vo.BasePaginationOutput, err error) {
	out = []*vo.GetManagerOutput{}

	// TODO: check user is admin
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	exists, err := s.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeGetUserAdminFailed, nil, nil, fmt.Errorf("get user admin failed: %s", err)
	}

	if !exists {
		return response.ErrCodeUserAdminNotFound, nil, nil, fmt.Errorf("user admin not found")
	}

	// TODO: ger all manager have pagination

	page := in.GetPage()
	limit := in.GetLimit()

	totalManagers, err := s.sqlc.CountNumberOfManagers(ctx)
	if err != nil {
		return response.ErrCodeCountNumberOfManagerFailed, nil, nil, fmt.Errorf("count number of manager failed: %s", err)
	}

	offset := (page - 1) * limit
	managers, err := s.sqlc.GetManagers(ctx, database.GetManagersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return response.ErrCodeGetManagerFailed, nil, nil, fmt.Errorf("get managers failed: %s", err)
	}

	for _, manager := range managers {
		createdAt, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(manager.CreatedAt))
		if err != nil {
			return response.ErrCodeConvertUnixToISOFailed, nil, nil, fmt.Errorf("convert ISO to Unix failed: %s", err)
		}

		updatedAt, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(manager.UpdatedAt))
		if err != nil {
			return response.ErrCodeConvertUnixToISOFailed, nil, nil, fmt.Errorf("convert ISO to Unix failed: %s", err)
		}

		out = append(out, &vo.GetManagerOutput{
			Account:   manager.Account,
			Username:  manager.UserName,
			IsDeleted: manager.IsDeleted == 1,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	totalPages := (totalManagers + int64(limit) - 1) / int64(limit)
	pagination = &vo.BasePaginationOutput{
		Page:       page,
		Limit:      limit,
		Total:      totalManagers,
		TotalPages: totalPages,
	}
	return response.ErrCodeGetManagerSuccess, out, pagination, nil
}
