package accommodationroom

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	GetAccommodationRooms(ctx *gin.Context, in *vo.GetAccommodationRoomsInput) (codeStatus int, out []*vo.GetAccommodationRoomsOutput, err error)
	CreateAccommodationRoom(ctx *gin.Context, in *vo.CreateAccommodationRoomInput) (codeStatus int, out []*vo.CreateAccommodationRoomOutput, err error)
	UpdateAccommodationRoom(ctx *gin.Context, in *vo.UpdateAccommodationRoomInput) (codeResult int, out *vo.UpdateAccommodationRoomOutput, err error)
	DeleteAccommodationRoom(ctx *gin.Context, in *vo.DeleteAccommodationRoomInput) (codeResult int, err error)
}
