package models

type Tweet struct {
	Tweet string `json:"tweet" required:"true"`
}

type UnTweet struct {
	TweetId int64 `json:"tweetId" required:"true"`
}

type LikeTweet struct {
	TweetId int64 `json:"tweetId" required:"true"`
}

type FollowUser struct {
	UserId int64 `json:"userId" required:"true"`
}

type Retweet struct {
	TweetId int64 `json:"tweetId" required:"true"`
}

type ReplyTweet struct {
	TweetId   int64  `json:"tweetId" required:"true"`
	TweetText string `json:"tweetText" required:"true"`
}

type Test struct {
	Text string `json:"text" required:"true"`
}
