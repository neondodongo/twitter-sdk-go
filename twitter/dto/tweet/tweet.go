package tweet

/*
 * Twitter API v2 Object Models
 *
 * https://developer.twitter.com/en/docs/twitter-api/data-dictionary/introduction
 */

// Tweet the basic building block of all things Twitter! :D
//
// Used in Tweet lookup operations
//
// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/tweet
type Tweet struct {
	// ID the unique identifier of the requested Tweet (default value)
	ID string `json:"id"`
	// Text the actual UTF-8 text of the Tweet (default value)
	Text string `json:"text"`
	// Attachments specifies the type of attachments present in this Tweet (if any)
	Attachments Attachments `json:"attachments,omitempty"`
	// AuthorID unique identifier of the user who posted this Tweet
	AuthorID string `json:"author_id,omitempty"`
	// ContextAnnotations contains context annotations for the Tweet
	ContextAnnotations []ContextAnnotation `json:"context_annotations,omitempty"`
	// ConversationID used to reconstruct the converstion from a Tweet
	ConversationID string `json:"conversation_id,omitempty"`
	// CreatedAt creation time of a Tweet
	CreatedAt string `json:"created_at,omitempty"`
	// Entities represents entities which have been parsed out of the text of a Tweet
	Entities Entities `json:"entities,omitempty"`
	// Geo contains details about the location tagged by the user/author of a tweet
	Geo Geo `json:"geo,omitempty"`
	// InReplyToUserID contains the original Tweet's author ID if a represented Tweet is a reply
	InReplyToUserID string `json:"in_reply_to_user_id,omitempty"`
	// Language of a Tweet if detected by Twitter (BCP47 launguage tag)
	Language string `json:"lang,omitempty"`
	// NonPublicMetrics non-public engagement metrics for a Tweet at the time of the request (requires user context authentication)
	NonPublicMetrics NonPublicMetrics `json:"non_public_metrics,omitempty"`
	// OrganicMetrics engagement metrics, tracked in an organic context, for a Tweet at the time of the request (requires user context authentication)
	OrganicMetrics EngagementMetrics `json:"organic_metrics,omitempty"`
	// PossiblySensitive an indicator that the URL contained in a Tweet may contain sensitive content or media
	PossiblySensitive bool `json:"possibly_sensistive,omitempty"`
	// PromotedMetrics engagement metrics, tracked in a promoted context, for a Tweet at the time of the request (requires user context authentication)
	PromotedMetrics EngagementMetrics `json:"promoted_metrics,omitempty"`
	// PublicMetrics public engagement metrics for a Tweet at the time of the request
	PublicMetrics PublicMetrics `json:"public_metrics,omitempty"`
	// ReferencedTweets a lit of Tweets a particular Tweet refers to
	ReferencedTweets []ReferencedTweet `json:"referenced_tweets,omitempty"`
	// ReplySettings shows who can reply to a Tweet
	ReplySettings string `json:"reply_settings,omitempty"`
	// Source the name of the app the user Tweeted from
	Source string `json:"source,omitempty"`
	// Withheld contains withholding details for withheld content
	Withheld Withheld `json:"withheld,omitempty"`
}

type Annotation struct {
	Start          int     `json:"start"`
	End            int     `json:"end"`
	Probability    float32 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

type Attachments struct {
	PollIDs   []string `json:"poll_ids"`
	MediaKeys []string `json:"media_keys"`
}

type Coordinates struct {
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}
type ContextAnnotation struct {
	Domain ContextEntity `json:"domain"`
	Entity ContextEntity `json:"entity"`
}
type ContextEntity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type EngagementMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	LikeCount         int `json:"like_count"`
	ReplyCount        int `json:"reply_count"`
	RetweetCount      int `json:"retweet_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
}

type Entities struct {
	Annotations []Annotation `json:"annotations"`
	CashTags    []Tag        `json:"cashtags"`
	HashTags    []Tag        `json:"hashtags"`
	Mentions    []Tag        `json:"mentions"`
	URLs        []URL        `json:"urls"`
}

type Geo struct {
	Coordinates Coordinates `json:"coordinates"`
	PlaceID     string      `json:"place_id"`
}

type NonPublicMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
}

type PublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type ReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Tag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type URL struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UnwoundURL  string `json:"unwound_url"`
}

type Withheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
}
