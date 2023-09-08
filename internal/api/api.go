package api

import (
	"fmt"
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

// Valid check if the string match one of the const values.
func (ad AnalysisDimension) Valid() bool {
	switch ad {
	case Likes, Comments, Favorites, Retweets:
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
		return fmt.Errorf("failed to parse duration : %w", err)
	}

	if !ap.Dimension.Valid() {
		return fmt.Errorf(`invalid dimension "%s"`, string(ap.Dimension))
	}

	return nil
}

// LikeDimensionPercentiles is percentiles values for dimension 'likes'.
type LikeDimensionPercentiles struct {
	LikesP50 uint64 `json:"likes_p50"`
	LikesP90 uint64 `json:"likes_p90"`
	LikesP99 uint64 `json:"likes_p99"`
}

// CommentDimensionPercentiles is percentiles values for dimension 'comments'.
type CommentDimensionPercentiles struct {
	CommentsP50 uint64 `json:"comments_p50"`
	CommentsP90 uint64 `json:"comments_p90"`
	CommentsP99 uint64 `json:"comments_p99"`
}

// FavoriteDimensionPercentiles is percentiles values for dimension 'favorites'.
type FavoriteDimensionPercentiles struct {
	FavoritesP50 uint64 `json:"favorites_p50"`
	FavoritesP90 uint64 `json:"favorites_p90"`
	FavoritesP99 uint64 `json:"favorites_p99"`
}

// RetweetDimensionPercentiles is percentiles values for dimension 'retweets'.
type RetweetDimensionPercentiles struct {
	RetweetsP50 uint64 `json:"retweets_p50"`
	RetweetsP90 uint64 `json:"retweets_p90"`
	RetweetsP99 uint64 `json:"retweets_p99"`
}

// AnalysisResponse is the format of /analysis JSON response.
type AnalysisResponse struct {
	TotalPosts                    uint64 `json:"total_posts"`
	MinTimestamp                  uint64 `json:"minimum_timestamp"`
	MaxTimestamp                  uint64 `json:"maximum_timestamp"`
	*LikeDimensionPercentiles     `json:",omitempty"`
	*CommentDimensionPercentiles  `json:",omitempty"`
	*FavoriteDimensionPercentiles `json:",omitempty"`
	*RetweetDimensionPercentiles  `json:",omitempty"`
}

// Fill set the required response fields for a particular dimension.
func (ar *AnalysisResponse) Fill(dimension AnalysisDimension, datas [3]uint64) {
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
