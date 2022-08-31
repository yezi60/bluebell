package mysql

import (
	"bluebell/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post 
    (post_id, title, content, author_id, community_id)
	values (?,?,?,?,?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据id查询单个帖子
func GetPostById(pid int64) (post *models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id = ?
	`
	post = new(models.Post)
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表函数 改进之后可以通过 score的排名与创建时间的排名进行帖子的获取
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
    post_id,title,content,author_id,community_id,create_time
	from post
	ORDER BY create_time
	DESC 
	limit ?,?
	`
	posts = make([]*models.Post, 0, size)

	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id,?)
	`
	// 根据在ids中出现的顺序进行排序 ids:[001,003,009,004]	sqlx应该是生成对应多数量的'?'
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}

	// 重新更正查询语句
	query = db.Rebind(query)

	// 开辟内存
	postList = []*models.Post{}

	// 传指针
	err = db.Select(&postList, query, args...) //!!!!!!!!

	return
}
