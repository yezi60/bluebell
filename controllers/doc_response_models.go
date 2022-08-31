package controllers

import "bluebell/models"

// 专门用来放接口文档用到的model
// 因为我们的接口文档返回的数据格式是一致的，但是具体的data类型不一致

// 写在字段后面的注释可以
type _ResponseData struct {
	Code ResCode                 `json:"code"` // 业务响应状态码
	Msg  string                  `json:"msg"`  // 提示信息
	Data []*models.ApiPostDetail `json:"data"` // 数据
}

type _ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page"`                   // 页码
	Size        int64  `json:"size" form:"size"`                   // 每页的数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}
