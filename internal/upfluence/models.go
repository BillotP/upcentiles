package upfluence

type StreamEventV2 struct {
	Timestamp int  `json:"timestamp"`
	Likes     *int `json:"likes"`
	Comments  *int `json:"comments"`
	Retweets  *int `json:"retweets"`
	Favorites *int `json:"favorites"`
}

type AnalysisValue struct {
	Timestamp      int `json:"timestamp"`
	DimensionValue int `json:"dimension_value"`
}

// StreamEvent is the data field from Upfluence SSE stream payload.
type StreamEvent struct {
	InstagramMedia *struct {
		ID                int      `json:"id"`
		Text              string   `json:"text"`
		Link              string   `json:"link"`
		Type              string   `json:"type"`
		LocationName      string   `json:"location_name"`
		Likes             int      `json:"likes"`
		Comments          int      `json:"comments"`
		Timestamp         int      `json:"timestamp"`
		PostID            string   `json:"post_id"`
		Views             int      `json:"views"`
		Mtype             int      `json:"mtype"`
		TaggedProfiles    []string `json:"tagged_profiles"`
		MentionedProfiles []string `json:"mentioned_profiles"`
		ThumbnailURL      string   `json:"thumbnail_url"`
		HiddenLikes       bool     `json:"hidden_likes"`
		Plays             int      `json:"plays"`
	} `json:"instagram_media,omitempty"`

	TiktokVideo *struct {
		ID           int      `json:"id"`
		Name         string   `json:"name"`
		ThumbnailURL string   `json:"thumbnail_url"`
		Link         string   `json:"link"`
		Comments     int      `json:"comments"`
		Likes        int      `json:"likes"`
		Timestamp    int      `json:"timestamp"`
		PostID       string   `json:"post_id"`
		Plays        int      `json:"plays"`
		Shares       int      `json:"shares"`
		Tags         []string `json:"tags"`
	} `json:"tiktok_video,omitempty"`

	Pin *struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Links       string `json:"links"`
		Likes       int    `json:"likes"`
		Comments    int    `json:"comments"`
		Saves       int    `json:"saves"`
		Repins      int    `json:"repins"`
		Timestamp   int    `json:"timestamp"`
		PostID      string `json:"post_id"`
	} `json:"pin,omitempty"`

	YoutubeVideo *struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Views       int    `json:"views"`
		Comments    int    `json:"comments"`
		Likes       int    `json:"likes"`
		Dislikes    int    `json:"dislikes"`
		Timestamp   int    `json:"timestamp"`
		PostID      string `json:"post_id"`
	} `json:"youtube_video,omitempty"`

	Article *struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Timestamp   int    `json:"timestamp"`
		URL         string `json:"url"`
		Content     string `json:"content"`
	} `json:"article,omitempty"`
	// Tweet need checks
	Tweet *struct {
		ID        int `json:"id"`
		Likes     int `json:"likes"`
		Retweets  int `json:"retweets"`
		Favorites int `json:"favorites"`
		Comments  int `json:"comments"`
		Timestamp int `json:"timestamp"`
	} `json:"tweet,omitempty"`
	// FacebookStatus need checks too
	FacebookStatus *struct {
		ID        int `json:"id"`
		Timestamp int `json:"timestamp"`
		Comments  int `json:"comments"`
		Likes     int `json:"likes"`
	} `json:"facebook_status,omitempty"`
}
