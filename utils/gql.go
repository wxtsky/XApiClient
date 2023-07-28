package utils

import (
	"TwitterApi/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func encodeParams(params map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range params {
		data, err := json.Marshal(v)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		result[k] = string(data)
	}
	return result
}
func Gql(method string, operation string, variables map[string]interface{}, c *gin.Context) (map[string]interface{}, error) {
	qid := config.AccountOperations[operation]
	op := operation
	params := map[string]interface{}{
		"queryId":   qid,
		"features":  config.DefaultFeatures,
		"variables": variables,
	}
	for k, v := range config.DefaultVariables {
		params["features"].(map[string]interface{})[k] = v
	}
	url := config.GqlAPI + "/" + qid + "/" + op
	headers := GetHeaders(c.GetHeader("auth_token"))
	client := resty.New()
	if c.GetString("proxy") != "" {
		client.SetProxy(c.GetString("proxy"))
	}
	if method == "POST" {
		post, err := client.R().SetHeaders(headers).SetBody(params).Post(url)
		if err != nil {
			return nil, err
		}
		var jsonData map[string]interface{}
		err = json.Unmarshal(post.Body(), &jsonData)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	} else if method == "GET" {
		get, err := client.R().SetHeaders(headers).SetQueryParams(encodeParams(params)).Get(url)
		if err != nil {
			return nil, err
		}
		var jsonData map[string]interface{}
		err = json.Unmarshal(get.Body(), &jsonData)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	} else {
		return nil, nil
	}
}
