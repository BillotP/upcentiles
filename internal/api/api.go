package api

import (
	"fmt"
	"strings"
	"time"
)

// VERSION is set at build time.
var VERSION = "dev"

// AnalysisDimension is either 'likes' , 'comments', 'favorites' or 'retweets'.
type AnalysisDimension string

const (
	Likes     AnalysisDimension = AnalysisDimension("likes")
	Comments  AnalysisDimension = AnalysisDimension("comments")
	Favorites AnalysisDimension = AnalysisDimension("favorites")
	Retweets  AnalysisDimension = AnalysisDimension("retweets")
)

// String return an AnalysisDimension underlying string data.
func (ad AnalysisDimension) String() string { return string(ad) }

var AllAnalysisDimensions = []string{
	Likes.String(),
	Comments.String(),
	Favorites.String(),
	Retweets.String(),
}

// Valid check if the string match one of the const values.
func (ad AnalysisDimension) Valid() bool {
	switch ad {
	case Likes,
		Comments,
		Favorites,
		Retweets:
		return true
	}

	return false
}

// AnalysisParam is the query parameters for /analysis handler.
type AnalysisParam struct {
	Duration  string            `query:"duration"`
	Dimension AnalysisDimension `query:"dimension"`
}

// Validate if both key of AnalysisParam struct are correctly defined.
func (ap AnalysisParam) Validate() error {
	_, err := time.ParseDuration(ap.Duration)
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w", err)
	}

	if !ap.Dimension.Valid() {
		return fmt.Errorf(`invalid dimension "%s", valid dimensions are : "%s"`,
			ap.Dimension.String(), strings.Join(AllAnalysisDimensions, ","))
	}

	return nil
}

// LikeDimensionPercentiles is percentiles values for dimension 'likes'.
type LikeDimensionPercentiles struct {
	LikesP50 int `json:"likes_p50"`
	LikesP90 int `json:"likes_p90"`
	LikesP99 int `json:"likes_p99"`
}

// CommentDimensionPercentiles is percentiles values for dimension 'comments'.
type CommentDimensionPercentiles struct {
	CommentsP50 int `json:"comments_p50"`
	CommentsP90 int `json:"comments_p90"`
	CommentsP99 int `json:"comments_p99"`
}

// FavoriteDimensionPercentiles is percentiles values for dimension 'favorites'.
type FavoriteDimensionPercentiles struct {
	FavoritesP50 int `json:"favorites_p50"`
	FavoritesP90 int `json:"favorites_p90"`
	FavoritesP99 int `json:"favorites_p99"`
}

// RetweetDimensionPercentiles is percentiles values for dimension 'retweets'.
type RetweetDimensionPercentiles struct {
	RetweetsP50 int `json:"retweets_p50"`
	RetweetsP90 int `json:"retweets_p90"`
	RetweetsP99 int `json:"retweets_p99"`
}

// AnalysisResponse is the format of /analysis JSON response.
type AnalysisResponse struct {
	TotalPosts                    int `json:"total_posts"`
	MinTimestamp                  int `json:"minimum_timestamp"`
	MaxTimestamp                  int `json:"maximum_timestamp"`
	*LikeDimensionPercentiles     `json:",omitempty"`
	*CommentDimensionPercentiles  `json:",omitempty"`
	*FavoriteDimensionPercentiles `json:",omitempty"`
	*RetweetDimensionPercentiles  `json:",omitempty"`
}

// FillPercentiles set the required response fields for a particular dimension.
func (ar *AnalysisResponse) FillPercentiles(dimension AnalysisDimension, datas [3]int) {
	switch dimension {
	case Likes:
		ar.LikeDimensionPercentiles = &LikeDimensionPercentiles{
			datas[0],
			datas[1],
			datas[2],
		}
	case Comments:
		ar.CommentDimensionPercentiles = &CommentDimensionPercentiles{
			datas[0],
			datas[1],
			datas[2],
		}
	case Favorites:
		ar.FavoriteDimensionPercentiles = &FavoriteDimensionPercentiles{
			datas[0],
			datas[1],
			datas[2],
		}
	case Retweets:
		ar.RetweetDimensionPercentiles = &RetweetDimensionPercentiles{
			datas[0],
			datas[1],
			datas[2],
		}
	}
}
