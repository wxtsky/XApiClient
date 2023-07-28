package utils

import (
	"github.com/gin-gonic/gin"
)

func ScraperRun(key string, qid string, name string, queries interface{}, c *gin.Context) interface{} {
	//// 创建一个键值对，键为key，值为queries
	//headers := GetHeaders(c.GetString("auth_token"))
	//_queries := map[string]interface{}{
	//	key: queries,
	//}
	////        params = {
	////            'variables': Operation.default_variables | keys | kwargs,
	////            'features': Operation.default_features,
	////        }
	//variables := config.DefaultVariables
	//variables[key] = queries
	//params := map[string]interface{}{
	//	"variables": _queries,
	//	"features":  config.DefaultVariables,
	//}
	//client := resty.New()
	//if c.GetString("proxy") != "" {
	//	client.SetProxy(c.GetString("proxy"))
	//}
	//url := "https://twitter.com/i/api/graphql/" + qid + "/" + name
	return nil
}
