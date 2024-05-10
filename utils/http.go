package Utils

import (
	"context"
	"encoding/json"
	"errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"net/url"
	"time"
)

func RequestByPost(ctx context.Context, url string, data url.Values) (result map[string]interface{}, err error) {
	client := http.Client{Timeout: time.Second * 10}
	res, err := client.PostForm(url, data)
	if err != nil {
		return nil, errors.New("创建http客户端失败" + err.Error())
	}
	defer res.Body.Close()
	res_json, err := io.ReadAll(res.Body)
	if err = json.Unmarshal(res_json, &result); err != nil {
		return nil, errors.New("响应数据反序列化失败" + err.Error())
	}
	var req_json []byte
	if req_json, err = json.Marshal(data); err != nil {
		return nil, err
	}
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("RequestParams", string(res_json)),
		attribute.String("ResponseParams", string(req_json)),
		attribute.String("Protocol", res.Request.Proto),
		attribute.String("URL", url),
		attribute.String("Method", res.Request.Method),
	)
	return result, nil
}
