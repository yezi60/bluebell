package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

// 测试controller层的参数校验
func TestCreatePostHandler(t *testing.T) {

	r := gin.Default()

	url := "/api/v1/post"

	r.POST(url, CreatePostHandler)

	body := `{
		"community_id":1,
		"title":"test",
		"content":"just a test"
	}`

	// 构造一个新的请求
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	// 构造一个记录器
	w := httptest.NewRecorder()

	// 发送请求并将响应反序列进记录器
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	//判断响应的内容是不是按预期返回了需要登陆的错误

	// 方法1：判断响应内容中是不是包含指定的字符串
	// assert.Contains(t, w.Body.String(), "需要登陆")

	// 方法2：将结果反序列回原本的结构体，判断ErrCode是否相同
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed,err: %v", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
