package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
	"time"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "172.30.198.210",
		User:         "root",
		Password:     "Mysql:12138",
		DbName:       "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	err := CreatePost(&models.Post{
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
		CreateTime:  time.Now(),
	})

	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed,err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
