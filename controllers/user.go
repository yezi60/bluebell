package controllers

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	// 从请求的参数绑定到结构体
	var p = new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应

		// 1. 字段为空
		// 2. password与repassword不相等

		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是vaildator类型的错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}

		// 翻译错误  去除多余的结构体提示内容
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 使用第三方库改进请求校验

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数以及参数校验
	var p = new(models.ParamsLogin)

	if err := c.ShouldBindJSON(p); err != nil {

		zap.L().Error("login params wrong", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login fail", zap.String("username", p.Username), zap.Error(err))

		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}

		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型最大值1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
