package upfluence

// StreamEvent is the data field from Upfluence SSE stream payload.
type StreamEvent struct {
	InstagramMedia *struct {
		ID                int      `json:"id"`
		Text              string   `json:"text"`
		Link              string   `json:"link"`
		Type              string   `json:"type"`
		LocationName      string   `json:"location_name"`
		Likes             uint64   `json:"likes"`
		Comments          uint64   `json:"comments"`
		Timestamp         uint64   `json:"timestamp"`
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
		Comments     uint64   `json:"comments"`
		Likes        uint64   `json:"likes"`
		Timestamp    uint64   `json:"timestamp"`
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
		Likes       uint64 `json:"likes"`
		Comments    uint64 `json:"comments"`
		Saves       int    `json:"saves"`
		Repins      int    `json:"repins"`
		Timestamp   uint64 `json:"timestamp"`
		PostID      string `json:"post_id"`
	} `json:"pin,omitempty"`

	YoutubeVideo *struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Views       int    `json:"views"`
		Comments    uint64 `json:"comments"`
		Likes       uint64 `json:"likes"`
		Dislikes    int    `json:"dislikes"`
		Timestamp   uint64 `json:"timestamp"`
		PostID      string `json:"post_id"`
	} `json:"youtube_video,omitempty"`

	Article *struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Timestamp   uint64 `json:"timestamp"`
		URL         string `json:"url"`
		Content     string `json:"content"`
	} `json:"article,omitempty"`
	// Tweet need checks
	Tweet *struct {
		ID        int    `json:"id"`
		Likes     uint64 `json:"likes"`
		Retweets  uint64 `json:"retweets"`
		Favorites uint64 `json:"favorites"`
		Comments  uint64 `json:"comments"`
		Timestamp uint64 `json:"timestamp"`
	} `json:"tweet,omitempty"`
	// FacebookStatus need checks too
	FacebookStatus *struct {
		ID        int    `json:"id"`
		Timestamp uint64 `json:"timestamp"`
		Comments  uint64 `json:"comments"`
		Likes     uint64 `json:"likes"`
	} `json:"facebook_status,omitempty"`
}

func (s *StreamEvent) Timestamp() uint64 {
	switch {
	case s.Article != nil:
		return s.Article.Timestamp
	case s.InstagramMedia != nil:
		return s.InstagramMedia.Timestamp
	case s.Pin != nil:
		return s.Pin.Timestamp
	case s.TiktokVideo != nil:
		return s.TiktokVideo.Timestamp
	case s.YoutubeVideo != nil:
		return s.YoutubeVideo.Timestamp
	case s.Tweet != nil:
		return s.Tweet.Timestamp
	case s.FacebookStatus != nil:
		return s.FacebookStatus.Timestamp
	default:
		return 0
	}
}

func (s *StreamEvent) Likes() uint64 {
	switch {
	case s.Article != nil:
		return 0
	case s.InstagramMedia != nil:
		return s.InstagramMedia.Likes
	case s.Pin != nil:
		return s.Pin.Likes
	case s.TiktokVideo != nil:
		return s.TiktokVideo.Likes
	case s.YoutubeVideo != nil:
		return s.YoutubeVideo.Likes
	case s.Tweet != nil:
		return s.Tweet.Likes
	case s.FacebookStatus != nil:
		return s.FacebookStatus.Likes
	default:
		return 0
	}
}

func (s *StreamEvent) Comments() uint64 {
	switch {
	case s.Article != nil:
		return 0
	case s.InstagramMedia != nil:
		return s.InstagramMedia.Comments
	case s.Pin != nil:
		return s.Pin.Comments
	case s.TiktokVideo != nil:
		return s.TiktokVideo.Comments
	case s.YoutubeVideo != nil:
		return s.YoutubeVideo.Comments
	case s.Tweet != nil:
		return s.Tweet.Comments
	case s.FacebookStatus != nil:
		return s.FacebookStatus.Comments
	default:
		return 0
	}
}

func (s *StreamEvent) Favorites() uint64 {
	switch {
	case s.Article != nil:
		return 0
	case s.InstagramMedia != nil:
		return 0
	case s.Pin != nil:
		return 0
	case s.TiktokVideo != nil:
		return 0
	case s.YoutubeVideo != nil:
		return 0
	case s.Tweet != nil:
		return s.Tweet.Favorites
	case s.FacebookStatus != nil:
		return 0
	default:
		return 0
	}
}

func (s *StreamEvent) Retweets() uint64 {
	switch {
	case s.Article != nil:
		return 0
	case s.InstagramMedia != nil:
		return 0
	case s.Pin != nil:
		return 0
	case s.TiktokVideo != nil:
		return 0
	case s.YoutubeVideo != nil:
		return 0
	case s.Tweet != nil:
		return s.Tweet.Retweets
	case s.FacebookStatus != nil:
		return 0
	default:
		return 0
	}
}
