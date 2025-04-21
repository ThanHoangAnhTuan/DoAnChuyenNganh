package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

type TestImpl struct {
	sqlc *database.Queries
}

// CreateAccommodation implements services.ITest.
func (t *TestImpl) CreateAccommodation(ctx context.Context, in *vo.CreateAccommodationInput) (codeResult int, out *vo.CreateAccommodationOutput, err error) {
	out = &vo.CreateAccommodationOutput{}
	id := uuid.New().String()
	now := utiltime.GetTimeNow()

	fmt.Printf("CreateAccommodation: %s", in)

	err = t.sqlc.CreateAccommodation(ctx, database.CreateAccommodationParams{
		Name:                 in.Name,
		City:                 in.City,
		Provine:              in.Provine,
		District:             in.District,
		Images:               in.Images.Filename,
		Description:          in.Description,
		Facilities:           json.RawMessage(in.Facilities),
		ID:                   id,
		ManagerID:            ctx.Value("user_id").(string),
		GgMap:                in.GoogleMap,
		PropertySurroundings: json.RawMessage(""),
		Rules:                in.Rules,
		CreatedAt:            now,
		UpdatedAt:            now,
	})

	if err != nil {
		return response.ErrCodeCreateAccommodationFailed, nil, fmt.Errorf("error for create accommodation: %s", err)
	}

	out.Id = id
	out.ManagerId = ctx.Value("user_id").(string)
	out.Name = in.Name
	out.City = in.City
	out.Provine = in.Provine
	out.District = in.District
	out.Images = in.Images.Filename
	out.Description = in.Description
	out.Facilities = json.RawMessage(in.Facilities)
	out.GoogleMap = in.GoogleMap
	out.PropertySurroundings = json.RawMessage("")
	out.Rules = in.Rules
	out.Rating = "0"

	fmt.Println("CreateAccommodation: ", out)

	return response.ErrCodeCreateAccommodationSuccess, out, nil
}

// CreateAccommodation implements services.ITest.

// LoginAdmin implements services.ITest.
func (t *TestImpl) LoginAdmin(ctx context.Context, in *vo.LoginInput) (codeResult int, out *vo.LoginOutput, err error) {
	panic("unimplemented")
}

// GetAccommodations implements services.ITest.
func (t *TestImpl) GetAccommodations(ctx context.Context) {
	panic("unimplemented")
}

func NewTestImpl(sqlc *database.Queries) *TestImpl {
	return &TestImpl{
		sqlc: sqlc,
	}
}
