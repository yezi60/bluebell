package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登陆")

// getCurrentUser 获取当前登陆用户的uuid
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	// uid, ok := c.Get(middlewares.ContextUserIDKey) 导致了循环引用

	uid, ok := c.Get(ContextUserIDKey)
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

// 获取分页参数
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}

	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
