package media

type Media struct {
	MediaKey         string           `json:"media_key"`
	Type             string           `json:"type"`
	DurationMS       int              `json:"duration_ms,omitempty"`
	Height           int              `json:"height,omitempty"`
	NonPublicMetrics NonPublicMetrics `json:"non_public_metrics,omitempty"`
	OrganicMetrics   OrganicMetrics   `json:"organic_metrics,omitempty"`
	PreviewImageURL  string           `json:"preview_image_url,omitempty"`
	PromotedMetrics  PromotedMetrics  `json:"PromotedMetrics,omitempty"`
	PublicMetrics    PublicMetrics    `json:"public_metrics,omitempty"`
	Width            int              `json:"width,omitempty"`
}

type NonPublicMetrics struct {
	Playback0Count   int `json:"playback_0_count"`
	Playback100Count int `json:"playback_100_count"`
	Playback25Count  int `json:"playback_25_count"`
	Playback50Count  int `json:"playback_50_count"`
	Playback75Count  int `json:"playback_75_count"`
}

type OrganicMetrics struct {
	Playback0Count   int `json:"playback_0_count"`
	Playback100Count int `json:"playback_100_count"`
	Playback25Count  int `json:"playback_25_count"`
	Playback50Count  int `json:"playback_50_count"`
	Playback75Count  int `json:"playback_75_count"`
	ViewCount        int `json:"view_count"`
}

type PromotedMetrics struct {
	Playback0Count   int `json:"playback_0_count"`
	Playback100Count int `json:"playback_100_count"`
	Playback25Count  int `json:"playback_25_count"`
	Playback50Count  int `json:"playback_50_count"`
	Playback75Count  int `json:"playback_75_count"`
	ViewCount        int `json:"view_count"`
}

type PublicMetrics struct {
	ViewCount int `json:"view_count"`
}
