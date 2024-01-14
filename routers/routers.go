package routers

import (
	"net/http"
	"newbeemall/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//商城前端路由
	mallRouter := RouterGroupApp.Mall
	mallGroup := r.Group("api")
	{
		mallRouter.InitMallUserRouter(mallGroup)
		mallRouter.InitMallUserAddressRouter(mallGroup)
		mallRouter.InitMallShopCartRouter(mallGroup)
	}

	//商城管理路由
	managerRouter := RouterGroupApp.Manage
	managerGroup := r.Group("manage-api")
	{
		managerRouter.InitManageAdminUserRouter(managerGroup)
		managerRouter.InitManageGoodsInfoRouter(managerGroup)
		managerRouter.InitManageGoodsCategoryRouter(managerGroup)
		managerRouter.InitManageCarouselRouter(managerGroup)
		managerRouter.InitManageIndexConfigRouter(managerGroup)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
