package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

// 把每一步数据库操作封装成函数
// 等待logic层根据业务需求调用

const secret = "soleaf.xyz"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(userName string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`

	var count int

	if err := db.Get(&count, sqlStr, userName); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("用户已存在")
	}

	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encrytPassword(user.Password)
	// 执行sql语句入库
	sqlStr := `insert into user(user_id,username,password)values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	sqlStr := `select user_id,username from user where user_id = ? `
	user = new(models.User)
	err = db.Get(user, sqlStr, uid)
	return
}

func Login(user *models.User) (err error) {

	oPassword := user.Password // 用户登陆的密码

	sqlStr := `select user_id,username,password from user where username = ? `

	err = db.Get(user, sqlStr, user.Username)

	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}

	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encrytPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return
}

func encrytPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	//转成16进制的字符串
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
