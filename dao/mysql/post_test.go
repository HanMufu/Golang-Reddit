package mysql

import (
	"go-web-app/models"
	"go-web-app/settings"
	"testing"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "35.201.223.76",
		User:         "root",
		Password:     "926443",
		DB:           "bluebell",
		Port:         23333,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		PostID:      10,
		AuthorId:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
