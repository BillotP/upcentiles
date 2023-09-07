package internal

import (
	"fmt"
	"time"
)

type AnalysisDimension string

const (
	Likes     AnalysisDimension = AnalysisDimension("likes")
	Comments  AnalysisDimension = AnalysisDimension("comments")
	Favorites AnalysisDimension = AnalysisDimension("favorites")
	Retweets  AnalysisDimension = AnalysisDimension("retweets")
)

func (ad AnalysisDimension) Valid() bool {
	switch ad {
	case Likes, Comments, Favorites, Retweets:
		return true
	}
	return false
}

type AnalysisParam struct {
	Duration  string            `query:"duration"`
	Dimension AnalysisDimension `query:"dimension"`
}

func (ap AnalysisParam) Validate() error {
	_, err := time.ParseDuration(ap.Duration)
	if err != nil {
		return err
	}
	if !ap.Dimension.Valid() {
		return fmt.Errorf(`invalid dimension "%s"`, string(ap.Dimension))
	}

	return nil
}

type LikeDimensionPercentiles struct {
	LikesP50 uint64 `json:"likes_p50"`
	LikesP90 uint64 `json:"likes_p90"`
	LikesP99 uint64 `json:"likes_p99"`
}

type CommentDimensionPercentiles struct {
	CommentsP50 uint64 `json:"comments_p50"`
	CommentsP90 uint64 `json:"comments_p90"`
	CommentsP99 uint64 `json:"comments_p99"`
}

type FavoriteDimensionPercentiles struct {
	FavoritesP50 uint64 `json:"favorites_p50"`
	FavoritesP90 uint64 `json:"favorites_p90"`
	FavoritesP99 uint64 `json:"favorites_p99"`
}

type RetweetDimensionPercentiles struct {
	RetweetsP50 uint64 `json:"retweets_p50"`
	RetweetsP90 uint64 `json:"retweets_p90"`
	RetweetsP99 uint64 `json:"retweets_p99"`
}

type AnalysisResponse struct {
	TotalPosts                    uint64 `json:"total_posts"`
	MinTimestamp                  uint64 `json:"minimum_timestamp"`
	MaxTimestamp                  uint64 `json:"maximum_timestamp"`
	*LikeDimensionPercentiles     `json:",omitempty"`
	*CommentDimensionPercentiles  `json:",omitempty"`
	*FavoriteDimensionPercentiles `json:",omitempty"`
	*RetweetDimensionPercentiles  `json:",omitempty"`
}

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
