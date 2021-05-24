package twitter

import (
	"github.com/neondodongo/twitter-sdk-go/twitter/dto/media"
	"github.com/neondodongo/twitter-sdk-go/twitter/dto/tweet"
	"github.com/neondodongo/twitter-sdk-go/twitter/dto/user"
)

// GetTweetByIDResponse JSON payload returned by the Twitter API v2 - /2/tweets/:id
type GetTweetByIDResponse struct {
	Tweet    tweet.Tweet   `json:"data"`
	Includes TweetIncludes `json:"includes,omitempty"`
	Error    Error         `json:"error,omitempty"`
}

// GetTweetsResponse JSON payload returned by the Twitter API v2 - /2/tweets
type GetTweetsResponse struct {
	Tweets   []tweet.Tweet `json:"data"`
	Includes TweetIncludes `json:"includes,omitempty"`
	Error    Error         `json:"error,omitempty"`
}

// GetUserByUsernameResponse JSON payload returned by the Twitter API v2 - /2/users/by/username/:username
type GetUserByUsernameResponse struct {
	UserData user.User    `json:"data"`
	Includes UserIncludes `json:"includes"`
	Error    Error        `json:"error,omitempty"`
}

// SearchTweetsResponse Twitter API v2 tweet search response
type SearchTweetsResponse struct {
	Data []tweet.SearchTweetsData `json:"data"`
}

// TweetIncludes partial JSON payload returned upon providing the optional "tweet.fields" query parameter
type TweetIncludes struct {
	Media []media.Media `json:"media"`
}

// UserIncludes partial JSON payload returned upon providing the optional "user.fields" query parameter
type UserIncludes struct {
	Tweets   []tweet.Tweet `json:"tweets"`
	Includes TweetIncludes `json:"includes,omitempty"`
}

// Error partial JSON payload returned when the Twitter API detects an issue with fulfilling the request
type Error struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
