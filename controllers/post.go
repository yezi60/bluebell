package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数以及参数的校验

	// validator --> binding tag
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJson(p) error", zap.Any("err", err))
		zap.L().Error("create post with binding fail", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从c取到当前发请求的用户id值
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从URL中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 根据id取出帖子数据（查数据库）
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)

}

// GetPostListHandler获取帖子列表的处理函数
func GetPostListHandler(c *gin.Context) {

	// 获取分页参数
	page, size := getPageInfo(c)

	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}

/*
根据前端传来的参数 获取帖子列表的处理函数
按创建时间排序 或者 按照分数排序
 1. 获取参数
 2. 去redis查询id值
 3. 根据id去数据库查询帖子详细信息
*/

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(分组)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ParamPostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	//GET请求参数：/api/v1/posts2?page=1&size=10&order=time   query
	//GET请求参数：/api/v1/posts2?p=1&s=10&o=time

	// 初始化结构体时指定一些具体的参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}

	// 还有c.ShouldBind(),可以自动根据对应的参数进行查询
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取数据
	data, err := logic.GetPostListNew(p) // 更新：合二为一

	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应

}

// 根据社区id查询帖子列表 与前面的进行整合了
//func GetCommunityPostListHandler(c *gin.Context) {
//
//	p := &models.ParamPostList{
//		Page:  1,
//		Size:  10,
//		Order: models.OrderTime,
//	}
//
//	// 还有c.ShouldBind(),可以自动根据对应的参数进行查询
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	// 获取数据
//
//	if err != nil {
//		zap.L().Error("logic.GetCommunityPostList() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//	// 返回响应
//
//}
