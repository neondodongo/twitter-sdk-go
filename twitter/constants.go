package twitter

const (

	// Twitter API Client Config Defaults
	defaultBaseURL       string = "https://api.twitter.com"
	defaultHTTPTimeoutMS int64  = 10000

	// Twitter API v2 Endpoints
	endpointV2TweetLookup                string = "/2/tweets"
	endpointV2TweetsSearchRecent         string = "/2/tweets/search/recent"
	endpointV2UserLookup                 string = "/2/users"
	endpointV2UserLookupByUsername       string = "/2/users/by"
	endpointV2UserLookupByUsernameSingle string = "/2/users/by/username"
)
