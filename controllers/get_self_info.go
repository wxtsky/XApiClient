package controllers

import (
	"TwitterApi/config"
	"TwitterApi/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func GetSelfInfo(c *gin.Context) {
	reqUrl := config.V1API + "/account/verify_credentials.json"
	headers := utils.GetHeaders(c.GetString("auth_token"))
	client := resty.New()
	fmt.Println(c.GetString("proxy"))
	if proxy := c.GetString("proxy"); proxy != "" {
		client.SetProxy(proxy)
	}
	resp, err := client.R().SetHeaders(headers).Get(reqUrl)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	var jsonData interface{}
	err = json.Unmarshal(resp.Body(), &jsonData)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"success": true,
		"data":    jsonData,
	})
}
