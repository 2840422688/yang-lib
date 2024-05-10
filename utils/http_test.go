package Utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestRequestByPost_Success(t *testing.T) {
	// 创建一个模拟的 HTTP 服务器
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 模拟成功的响应数据
		w.Write([]byte(`{"status": "success", "data": {"message": "Hello, World!"}}`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// 使用模拟的服务器地址来测试 RequestByPost 函数
	data := url.Values{}
	data.Set("key", "value")

	result, err := RequestByPost(server.URL, data)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"status": "success", "data": map[string]interface{}{"message": "Hello, World!"}}, result)
}

func TestRequestByPost_HTTPError(t *testing.T) {
	// 使用一个无效的地址来测试 RequestByPost 函数
	invalidURL := "invalid-url"

	data := url.Values{}
	data.Set("key", "value")

	_, err := RequestByPost(invalidURL, data)

	// 验证是否返回了预期的错误
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "创建http客户端失败"))
}

func TestRequestByPost_JSONError(t *testing.T) {
	// 创建一个模拟的 HTTP 服务器
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 模拟返回一个无效的 JSON 数据
		w.Write([]byte(`invalid-json-data`))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// 使用模拟的服务器地址来测试 RequestByPost 函数
	data := url.Values{}
	data.Set("key", "value")

	_, err := RequestByPost(server.URL, data)

	// 验证是否返回了预期的错误
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "响应数据反序列化失败"))
}

func TestRequestByPost_Timeout(t *testing.T) {
	// 创建一个模拟的 HTTP 服务器
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 模拟超时
		time.Sleep(time.Second * 10)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	// 使用模拟的服务器地址来测试 RequestByPost 函数
	data := url.Values{}
	data.Set("key", "value")

	_, err := RequestByPost(server.URL, data)

	// 验证是否返回了预期的错误
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "创建http客户端失败"))
}
