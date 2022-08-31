package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2. 生成UID
	userID := snowflake.GenID()
	// 构造一个user实例
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 3. 保存进数据库
	mysql.InsertUser(u)

	return
}

func Login(p *models.ParamsLogin) (user *models.User, err error) {

	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 传递的是指针，就能拿到user.UserID
	if err = mysql.Login(user); err != nil {
		//登陆失败,err中已经存了原因
		return nil, err
	}

	// 生成JWT的token
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
