package controllers

import (
	"TwitterApi/config"
	"TwitterApi/models"
	"TwitterApi/utils"
	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	var request models.FollowUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Please provide a valid user id",
		})
		return
	}
	setting := config.FollowSettings
	setting["user_id"] = request.UserId
	resp, err := utils.V1("friendships/create.json", setting, c)
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
