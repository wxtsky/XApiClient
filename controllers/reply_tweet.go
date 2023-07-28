package controllers

import (
	"TwitterApi/models"
	"TwitterApi/utils"
	"github.com/gin-gonic/gin"
)

func ReplyTweet(c *gin.Context) {
	var request models.ReplyTweet
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}
	resp, err := utils.Gql("POST", "CreateTweet", map[string]interface{}{
		"tweet_id":                request.TweetId,
		"tweet_text":              request.TweetText,
		"batch_compose":           "BatchSubsequent",
		"semantic_annotation_ids": []string{},
		"reply": map[string]interface{}{
			"in_reply_to_tweet_id":   request.TweetId,
			"exclude_reply_user_ids": []string{},
		},
		"media": map[string]interface{}{
			"media_entities":     []string{},
			"possibly_sensitive": false,
		},
	},
		c)
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
