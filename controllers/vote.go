package controllers

import (
	"bluebell/logic"
	"bluebell/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// 投票

func PostVoteController(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)

	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("投票问题:", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除错误提示中的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	// 获取当前请求的用户id
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
