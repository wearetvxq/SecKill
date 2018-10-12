package setup

import (
	"github.com/wearetvxq/SecKill/sk_admin/controller/activity"
	"github.com/wearetvxq/SecKill/sk_admin/controller/product"
	"github.com/gin-gonic/gin"
)

//设置路由
func setupRouter(router *gin.Engine) {
	//商品
	router.GET("/product/list", product.GetPorductList)
	router.POST("/product/create", product.CreateProduct)

	//活动
	router.GET("/activity/list", activity.GetActivityList)
	router.POST("/activity/create", activity.CreateActivity)
}
