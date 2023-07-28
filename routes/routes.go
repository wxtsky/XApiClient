package routes

import (
	"TwitterApi/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/self_info", controllers.GetSelfInfo)
	router.POST("/delete_tweet", controllers.DeleteTweet)
	router.POST("/like_tweet", controllers.LikeTweet)
	router.POST("/follow_user", controllers.FollowUser)
	router.POST("/retweet_tweet", controllers.RetweetTweet)
	router.POST("/reply_tweet", controllers.ReplyTweet)
}
