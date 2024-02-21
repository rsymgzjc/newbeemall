package user

import (
	"newbeemall/controllers"
	"newbeemall/logic/user"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SearchGoodsHandler(c *gin.Context) {
	page, _ := controllers.GetPageInfo(c)
	categorystr := c.Query("goodsCategoryId")
	categoryid, err := strconv.ParseInt(categorystr, 10, 64)
	if err != nil {
		return
	}
	keyword := c.Query("keyword")
	orderby := c.Query("orderby")
	datas, err := user.SearchGoodsInfo(page, categoryid, keyword, orderby)
	if err != nil {
		zap.L().Error("SearchGoodsInfo with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}
