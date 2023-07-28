package utils

import (
	"TwitterApi/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/url"
)

func V1(path string, params map[string]interface{}, c *gin.Context) (map[string]interface{}, error) {
	headers := GetHeaders(c.GetHeader("auth_token"))
	headers["content-type"] = "application/x-www-form-urlencoded"
	client := resty.New()
	queryParams := url.Values{}
	for k, v := range params {
		queryParams.Set(k, fmt.Sprintf("%v", v))
	}
	if c.GetString("proxy") != "" {
		client.SetProxy(c.GetString("proxy"))
	}
	fmt.Println(config.V1API + "/" + path + "?" + queryParams.Encode())
	r, err := client.R().SetHeaders(headers).Post(config.V1API + "/" + path + "?" + queryParams.Encode())
	if err != nil {
		return nil, err
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(r.Body(), &jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
