package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.uber.org/zap"

	// _ "github.com/thanhoanganhtuan/DoAnChuyenNganh/docs"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/initializer"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/tracing"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
// @schema http
func main() {
	// Initialize OpenTelemetry tracing
	cleanup, err := tracing.InitTracer("doan-chuyen-nganh-api")
	if err != nil {
		global.Logger.Error("Failed to initialize tracing", zap.Error(err))
	}
	defer cleanup()

	r := initializer.Run()
	port := strconv.Itoa(global.Config.Server.Port)

	// Add metrics endpoint to Gin router
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Add swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	global.Logger.Info("Server starting", zap.String("port", port))
	r.Run(":" + port)
}
