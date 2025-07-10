package accommodationdetail

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
	"go.uber.org/zap"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{sqlc: sqlc}
}

func (a *serviceImpl) CreateAccommodationDetail(ctx *gin.Context, in *vo.CreateAccommodationDetailInput) (codeStatus int, out *vo.CreateAccommodationDetailOutput, err error) {
	out = &vo.CreateAccommodationDetailOutput{}

	// TODO: get userID from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, nil, fmt.Errorf("manager not found")
	}

	// TODO: check accommodation exists
	accommodation, err := a.sqlc.GetAccommodationByIdNoVerify(ctx, in.AccommodationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: check if the manager is the owner of the accommodation
	if accommodation.ManagerID != userID {
		return response.ErrCodeForbidden, nil, fmt.Errorf("user is not the owner of the accommodation")
	}

	bedsJson, err := json.Marshal(in.Beds)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	facilitiesJson, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	accommodationDetailID := uuid.New().String()
	now := utiltime.GetTimeNow()

	price, err := decimal.NewFromString(strings.TrimSpace(in.Price))
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("invalid price format: %v", err)
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("price must be positive")
	}

	// TODO: create accommodation detail
	err = a.sqlc.CreateAccommodationDetail(ctx, database.CreateAccommodationDetailParams{
		ID:              accommodationDetailID,
		AccommodationID: accommodation.ID,
		Name:            in.Name,
		Guests:          in.Guests,
		Price:           price,
		Beds:            bedsJson,
		Facilities:      facilitiesJson,
		CreatedAt:       now,
		UpdatedAt:       now,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for create accommodation details: %s", err)
	}

	// TODO: get facility detail
	var facilityIds []string
	if err := json.Unmarshal(facilitiesJson, &facilityIds); err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error unmarshaling facility detail: %s", err)
	}

	for _, facilityId := range facilityIds {
		facility, err := a.sqlc.GetAccommodationFacilityDetailById(ctx, facilityId)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("get facility detail failed: %s", err)
		}

		out.Facilities = append(out.Facilities, vo.FacilityDetailOutput{
			ID:   facility.ID,
			Name: facility.Name,
		})
	}

	out.ID = accommodationDetailID
	out.AccommodationID = in.AccommodationID
	out.Beds = in.Beds
	out.DiscountID = in.DiscountID
	out.Guests = in.Guests
	out.Name = in.Name
	out.Price = price.String()
	return response.ErrCodeCreateAccommodationTypeSuccess, out, nil
}

func (a *serviceImpl) DeleteAccommodationDetail(ctx *gin.Context, in *vo.DeleteAccommodationDetailInput) (codeResult int, err error) {
	// TODO: get user from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, fmt.Errorf("manager not found")
	}

	// TODO: check the accommodation detail exists
	exists, err := a.sqlc.CheckAccommodationDetailExists(ctx, in.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for get accommodation detail: %s", err)
	}

	if !exists {
		return response.ErrCodeAccommodationTypeNotFound, nil
	}

	// TODO: check the accommodation detail belongs to manager
	isBelongs, err := a.sqlc.IsAccommodationDetailBelongsToManager(ctx, database.IsAccommodationDetailBelongsToManagerParams{
		ID:   userID,
		ID_2: in.ID,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for delete accommodation detail: %s", err)
	}
	if !isBelongs {
		return response.ErrCodeForbidden, fmt.Errorf("error for do not have permission to delete accommodation detail")
	}

	// TODO: delete accommodation detail
	err = a.sqlc.DeleteAccommodationDetail(ctx, in.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for delete accommodation detail: %s", err)
	}
	return response.ErrCodeDeleteAccommodationTypeSuccess, nil
}

func (a *serviceImpl) GetAccommodationDetails(ctx *gin.Context, in *vo.GetAccommodationDetailsInput) (codeStatus int, out []*vo.GetAccommodationDetailsOutput, err error) {
	out = []*vo.GetAccommodationDetailsOutput{}

	// 1. Lấy thông tin accommodation
	accommodation, err := a.sqlc.GetAccommodationByIdNoVerify(ctx, in.AccommodationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get accommodation failed: %w", err)
	}

	// 2. Lấy danh sách accommodation detail
	accommodationDetails, err := a.sqlc.GetAccommodationDetails(ctx, accommodation.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get accommodation details failed: %s", err)
	}

	var (
		allFacilityIDs         []string
		accommodationDetailIDs []string
		facilityIDsMap         = make(map[string][]string)
	)

	// 3. Gom facilityIDs và detailIDs
	for _, detail := range accommodationDetails {
		accommodationDetailIDs = append(accommodationDetailIDs, detail.ID)

		var facilityIDs []string
		if err := json.Unmarshal(detail.Facilities, &facilityIDs); err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("unmarshal facilities failed: %s", err)
		}
		facilityIDsMap[detail.ID] = facilityIDs
		allFacilityIDs = append(allFacilityIDs, facilityIDs...)
	}

	// 4. Truy vấn batch facility và đưa vào map
	uniqueFacilityIDs := removeDuplicateStrings(allFacilityIDs)
	allFacilities, err := a.sqlc.GetAccommodationFacilitiesByIds(ctx, uniqueFacilityIDs)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get facilities failed: %s", err)
	}

	facilityMap := make(map[string]vo.FacilityDetailOutput)
	for _, facility := range allFacilities {
		facilityMap[facility.ID] = vo.FacilityDetailOutput{
			ID:   facility.ID,
			Name: facility.Name,
		}
	}

	// 5. Lấy danh sách ảnh accommodation detail
	imageMap := make(map[string][]string)
	for _, detailID := range accommodationDetailIDs {
		images, err := a.sqlc.GetAccommodationDetailImages(ctx, detailID)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("get images failed: %s", err)
		}
		for _, img := range images {
			imageMap[detailID] = append(imageMap[detailID], img.Image)
		}
	}

	// 6. Đếm số lượng phòng còn trống nếu có ngày checkin - checkout
	availableRoomsMap := make(map[string]int64)
	if in.CheckIn != "" && in.CheckOut != "" {
		checkIn, err := utiltime.ConvertISOToUnixTimestamp(in.CheckIn)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("convert checkIn failed: %s", err)
		}
		checkOut, err := utiltime.ConvertISOToUnixTimestamp(in.CheckOut)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("convert checkOut failed: %s", err)
		}

		// Sử dụng batch query
		results, err := a.sqlc.BatchCountAccommodationRoomAvailable(ctx, database.BatchCountAccommodationRoomAvailableParams{
			CheckIn:  checkIn,
			CheckOut: checkOut,
			Ids:      accommodationDetailIDs,
		})
		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("batch count available room failed: %s", err)
		}
		for _, result := range results {
			availableRoomsMap[result.AccommodationTypeID] = result.AvailableCount
		}
	}

	// 7. Tạo output
	for _, detail := range accommodationDetails {
		var beds vo.Beds
		if err := json.Unmarshal(detail.Beds, &beds); err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("unmarshal beds failed: %s", err)
		}

		// Map facilities
		var facilities []vo.FacilityDetailOutput
		for _, fid := range facilityIDsMap[detail.ID] {
			if f, ok := facilityMap[fid]; ok {
				facilities = append(facilities, f)
			}
		}

		// Available rooms
		availableRooms := availableRoomsMap[detail.ID]

		out = append(out, &vo.GetAccommodationDetailsOutput{
			ID:             detail.ID,
			Name:           detail.Name,
			Guests:         detail.Guests,
			Beds:           beds,
			Facilities:     facilities,
			AvailableRooms: uint8(availableRooms),
			Price:          detail.Price.String(),
			DiscountID:     detail.DiscountID.String,
			Images:         imageMap[detail.ID],
		})
	}

	return response.ErrCodeGetAccommodationTypeSuccess, out, nil
}

func (a *serviceImpl) GetAccommodationDetailsByManager(ctx *gin.Context, in *vo.GetAccommodationDetailsByManagerInput) (codeStatus int, out []*vo.GetAccommodationDetailsByManagerOutput, err error) {
	out = []*vo.GetAccommodationDetailsByManagerOutput{}

	accommodation, err := a.sqlc.GetAccommodationByIdNoVerify(ctx, in.AccommodationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	accommodationDetails, err := a.sqlc.GetAccommodationDetails(ctx, accommodation.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get accommodation by id failed: %s", err)
	}

	var allFacilityIDs []string
	var accommodationDetailIDs []string
	facilityIDsMap := make(map[string][]string) // detail_id -> facility_ids

	for _, detail := range accommodationDetails {
		accommodationDetailIDs = append(accommodationDetailIDs, detail.ID)

		var facilityIDs []string
		if err := json.Unmarshal(detail.Facilities, &facilityIDs); err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("error unmarshaling facilities: %s", err)
		}

		facilityIDsMap[detail.ID] = facilityIDs
		allFacilityIDs = append(allFacilityIDs, facilityIDs...)
	}

	// Loại bỏ duplicate facility IDs
	uniqueFacilityIDs := removeDuplicateStrings(allFacilityIDs)

	// 4. Batch query facilities
	allFacilities, err := a.sqlc.GetAccommodationFacilitiesByIds(ctx, uniqueFacilityIDs)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get facilities failed: %s", err)
	}

	// Tạo map để lookup nhanh
	facilityMap := make(map[string]vo.FacilityDetailOutput)
	for _, facility := range allFacilities {
		facilityMap[facility.ID] = vo.FacilityDetailOutput{
			ID:   facility.ID,
			Name: facility.Name,
		}
	}

	availableRoomsResult, err := a.sqlc.BatchCountAccommodationRoomAvailableByManager(ctx, accommodationDetailIDs)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("batch count rooms failed: %s", err)
	}

	// Tạo map để lookup nhanh
	availableRoomsMap := make(map[string]int64)
	for _, result := range availableRoomsResult {
		availableRoomsMap[result.AccommodationTypeID] = result.AvailableCount
	}

	// Đảm bảo tất cả detail đều có count (set 0 nếu không tìm thấy)
	for _, detailID := range accommodationDetailIDs {
		if _, exists := availableRoomsMap[detailID]; !exists {
			availableRoomsMap[detailID] = 0
		}
	}

	// 6. Xử lý và tạo output
	for _, accommodationDetail := range accommodationDetails {
		// Unmarshal beds
		beds := vo.Beds{}
		if err := json.Unmarshal(accommodationDetail.Beds, &beds); err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("error unmarshaling beds: %s", err)
		}

		// Lấy facilities từ map (không cần query)
		facilities := []vo.FacilityDetailOutput{}
		for _, facilityID := range facilityIDsMap[accommodationDetail.ID] {
			if facility, exists := facilityMap[facilityID]; exists {
				facilities = append(facilities, facility)
			}
		}

		// Lấy available rooms từ map
		availableRooms := availableRoomsMap[accommodationDetail.ID]

		out = append(out, &vo.GetAccommodationDetailsByManagerOutput{
			ID:             accommodationDetail.ID,
			Name:           accommodationDetail.Name,
			Guests:         accommodationDetail.Guests,
			Beds:           beds,
			Facilities:     facilities,
			AvailableRooms: uint8(availableRooms),
			Price:          accommodationDetail.Price.String(),
			DiscountID:     accommodationDetail.DiscountID.String,
		})
	}
	return response.ErrCodeGetAccommodationTypeSuccess, out, nil
}

func (a *serviceImpl) UpdateAccommodationDetail(ctx *gin.Context, in *vo.UpdateAccommodationDetailInput) (codeResult int, out *vo.UpdateAccommodationDetailOutput, err error) {
	out = &vo.UpdateAccommodationDetailOutput{}

	// TODO: get user from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is manager
	manager, err := a.sqlc.CheckUserManagerExistsByID(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get manager: %s", err)
	}

	if !manager {
		return response.ErrCodeForbidden, nil, fmt.Errorf("manager not found")
	}

	// TODO: check accommodation exists
	accommodation, err := a.sqlc.GetAccommodationByIdNoVerify(ctx, in.AccommodationID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get accommodation: %w", err)
	}

	// TODO: check the accommodation detail exists
	isExists, err := a.sqlc.CheckAccommodationDetailExists(ctx, in.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for get accommodation detail: %s", err)
	}

	if !isExists {
		return response.ErrCodeAccommodationTypeNotFound, nil, fmt.Errorf("get accommodation detail not found")
	}

	// TODO: Check the user is the owner of the accommodation detail
	isBelongs, err := a.sqlc.IsAccommodationDetailBelongsToManager(ctx, database.IsAccommodationDetailBelongsToManagerParams{
		ID:   accommodation.ManagerID,
		ID_2: in.ID,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for update accommodation detail: %s", err)
	}
	if !isBelongs {
		return response.ErrCodeForbidden, nil, fmt.Errorf("error for do not have permission to update accommodation detail")
	}

	// TODO: update accommodation detail
	bedsJson, err := json.Marshal(in.Beds)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for marshal beds: %s", err)
	}

	facilitiesJson, err := json.Marshal(in.Facilities)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for marshal facilities: %s", err)
	}

	price, err := decimal.NewFromString(strings.TrimSpace(in.Price))
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("invalid price format: %v", err)
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return response.ErrCodeInternalServerError, nil, errors.New("price must be positive")
	}

	now := utiltime.GetTimeNow()
	err = a.sqlc.UpdateAccommodationDetail(ctx, database.UpdateAccommodationDetailParams{
		Name:            in.Name,
		Guests:          in.Guests,
		Beds:            bedsJson,
		Facilities:      facilitiesJson,
		Price:           price,
		UpdatedAt:       now,
		ID:              in.ID,
		AccommodationID: in.AccommodationID,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for update accommodation detail failed: %s", err)
	}

	// TODO: get images of accommodation detail
	accommodationDetailImages, err := a.sqlc.GetAccommodationDetailImages(ctx, in.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get images of accommodation failed: %s", err)
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
	out.Beds = in.Beds
	out.DiscountID = in.DiscountID
	out.Guests = in.Guests
	out.ID = in.ID
	out.Name = in.Name
	out.Price = price.String()
	out.Images = pathNames

	return response.ErrCodeUpdateAccommodationTypeSuccess, out, nil
}

func removeDuplicateStrings(slice []string) []string {
	keys := make(map[string]bool)
	result := []string{}

	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}
