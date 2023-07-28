package controllers

import (
	"TwitterApi/models"
	"TwitterApi/utils"
	"github.com/gin-gonic/gin"
)

func LikeTweet(c *gin.Context) {
	var request models.LikeTweet
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}
	resp, err := utils.Gql("POST", "FavoriteTweet", map[string]interface{}{"tweet_id": request.TweetId}, c)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    resp,
	})
}
