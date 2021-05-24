package user

type User struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Username        string            `json:"username"`
	CreatedAt       string            `json:"created_at,omitempty"`
	Description     string            `json:"description,omitempty"`
	Entities        Entities          `json:"entities,omitempty"`
	Location        string            `json:"location,omitempty"`
	PinnedTweetID   string            `json:"pinned_tweet_id,omitempty"`
	ProfileImageURL string            `json:"profile_image_url,omitempty"`
	Protected       bool              `json:"protected,omitempty"`
	PublicMetrics   UserPublicMetrics `json:"public_metrics,omitempty"`
	URL             string            `json:"url,omitempty"`
	Verified        bool              `json:"verified,omitempty"`
	Withheld        UserWithheld      `json:"withheld,omitempty"`
}

type Entities struct {
	URL         EntityURL         `json:"url"`
	Description EntityDescription `json:"description"`
}

type EntityDescription struct {
	URLs     []URL     `json:"urls"`
	Hashtags []Hashtag `json:"hashtags"`
	Mentions []Mention `json:"mentions"`
	Cashtags []Cashtag `json:"cashtags"`
}

type EntityURL struct {
	URLs []URL `json:"urls"`
}

type Cashtag struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Cashtag string `json:"cashtag"`
}

type Hashtag struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Hashtag string `json:"hashtag"`
}

type Mention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Username string `json:"username"`
}

type URL struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
}

type UserPublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type UserWithheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}
