package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	{
		"code":10000, // 程序中的错误码
		"msg": 		//提示信息
		"data": {}  // 数据

	}

*/

// 统一封装了返回的数据结构  本质上都是gin.H也就是map[string]interface{}{}

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	// 给前端的流程状态码都是200，前端显示的才是真的业务状态码
	c.JSON(http.StatusOK, rd)
}
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	// 给前端的流程状态码都是200，前端显示的才是真的业务状态码
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	// 给前端的流程状态码都是200，前端显示的才是真的业务状态码
	// 反序列化了
	c.JSON(http.StatusOK, rd)
}
