package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成post id

	p.ID = snowflake.GenID()

	// 2. 保存到数据库并返回

	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	// 3. 存redis
	return redis.CreatePost(p.ID, p.CommunityID)

}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {

	// 查询数据并组合我们接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}
	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(pid) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(pid) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return

}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)

	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(pid) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(pid) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}

	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 2. 去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("get ids from redis failed")
		return
	}

	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostList2 redis query success but len == 0")
		return
	}

	// 3. 根据id去MySQL数据库查询帖子详细信息
	// 返回的顺序要按照给定的id的顺序去返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("get ids from mysql failed")
		return
	}

	// 提前查询好每一篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		zap.L().Error("get votaData from redis failed", zap.Error(err))
		return
	}

	// 初始化切片
	data = make([]*models.ApiPostDetail, 0, len(posts))

	// 将帖子的作者以及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(pid) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(pid) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx], // 顺序是一致的
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}

	return
}

// GetCommunityPostList 根据社区分类获取帖子
func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {

	// 1. 去redis查询获取社区-帖子id
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("get ids from redis failed")
		return
	}

	// 2. 判断返回切片长度是否为0
	if len(ids) == 0 {
		zap.L().Warn("redis.GetCommunityPostIDsInOrder redis query success but len == 0")
		return
	}

	// 3. 根据id去MySQL数据库查询帖子详细信息
	// 返回的顺序要按照给定的id的顺序去返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("get ids from mysql failed")
		return
	}

	// 提前查询好每一篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		zap.L().Error("get votaData from redis failed", zap.Error(err))
		return
	}

	// 初始化切片
	data = make([]*models.ApiPostDetail, 0, len(posts))

	// 将帖子的作者以及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(pid) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(pid) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx], // 顺序是一致的
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}

	return
}

// GetPostListNew 将两个查询帖子的接口合二为一
func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 根据请求参数的不同，执行不同的逻辑
	if p.CommunityID == 0 {
		// 查所有
		data, err = GetPostList2(p)
	} else {
		// 根据社区id查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("logic.GetPostListNew failed", zap.Error(err))
		return
	}
	return
}
