package Utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

func RequestByPost(url string, data url.Values) (result map[string]interface{}, err error) {
	client := http.Client{Timeout: time.Second * 5}
	res, err := client.PostForm(url, data)
	if err != nil {
		return nil, errors.New("创建http客户端失败" + err.Error())
	}
	res_json, err := io.ReadAll(res.Body)
	if err = json.Unmarshal(res_json, &result); err != nil {
		return nil, errors.New("响应数据反序列化失败" + err.Error())
	}
	return result, nil
}
