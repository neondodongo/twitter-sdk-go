package tweet

// SearchTweetsData Twitter API v2 individual tweet data
type SearchTweetsData struct {
	// ID unique identifier of a Tweet (default field)
	ID string `json:"id"`
	// Text content of a Tweet (default field)
	Text string `json:"text"`

	// Optional fields applied if present in the request's query parameters

	// AuthorID unique identifier of user who created the Tweet
	AuthorID string `json:"author_id,omitempty"`
	// ConversationID id of the original Tweet of the conversation
	ConversationID string `json:"conversation_id,omitempty"`
	// CreatedAt creation time of the Tweet
	CreatedAt string `json:"created_at,omitempty"`
	// InReplyToUserID indicates the user ID of the parent Tweet's author if this Tweet is a reply
	InReplyToUserID string `json:"in_reply_to_user_id,omitempty"`
	// ReferencedTweets a list of Tweets this Tweet refers to
	ReferencedTweets string `json:"type,omitempty"`
}
