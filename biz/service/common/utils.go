package common

import (
	"math"
	"story_telling_backend/biz/model/story_telling_backend"
)

func GetPagination(pagination *story_telling_backend.Pagination) (int, int) {
	if pagination == nil {
		return 0, 10
	}
	pageNo := math.Max(float64(pagination.PageNo), 1)
	pageSize := math.Max(float64(pagination.PageSize), 10)
	return int(pageNo), int(pageSize)
}
