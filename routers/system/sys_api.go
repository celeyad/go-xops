package system

import (
	"go-xops/api/v1/system"
	"go-xops/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 接口路由
func InitApiRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("api").Use(authMiddleware.MiddlewareFunc()).Use(middleware.CasbinMiddleware)
	{
		router.GET("/list", system.GetApis)
		router.POST("/create", system.CreateApi)
		router.PATCH("/update/:apiId", system.UpdateApiById)
		router.DELETE("/delete", system.BatchDeleteApiByIds)
	}
	return router
}
