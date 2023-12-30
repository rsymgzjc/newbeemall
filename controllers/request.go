package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const CtxUserID = "userID"

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserID)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int64, int64) {
	PageStr := c.Query("page")
	SizeStr := c.Query("size")
	Page, err := strconv.ParseInt(PageStr, 10, 64)
	if err != nil {
		Page = 1
	}
	Size, err := strconv.ParseInt(SizeStr, 10, 64)
	if err != nil {
		Size = 10
	}
	return Page, Size
}
