package service

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/huqiyii/facility/conve"
	jsoniter "github.com/json-iterator/go"
	"story_telling_backend/biz/db"
	"story_telling_backend/biz/model/db_model"
	"story_telling_backend/biz/model/story_telling_backend"
	"story_telling_backend/biz/service/common"
)

func SearchNovel(req *story_telling_backend.SearchNovelReq) (*story_telling_backend.SearchNovelData, error) {
	dbClient := db.GetDBClient().Debug()
	query := dbClient.Model(&db_model.BookTable{}).Where("status = 1")
	if req.CustomValue != nil {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *req.CustomValue))
	}
	if len(req.Tags) > 0 {
		query = query.Where("JSON_CONTAINS(tags, JSON_Array(?))", req.Tags)
	}
	var novels []db_model.BookTable
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	pageNo, pageSize := common.GetPagination(req.GetPagination())
	if err := query.Order("id").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&novels).Error; err != nil {
		return nil, err
	}
	data := &story_telling_backend.SearchNovelData{
		Total: total,
		Items: make([]*story_telling_backend.SearchNovelItem, 0),
	}
	for _, novel := range novels {
		item := &story_telling_backend.SearchNovelItem{
			ID:          conve.Int64Default(novel.ID, -1),
			CoverUrl:    conve.StringDefault(novel.CoverURI, ""),
			NovelName:   conve.StringDefault(novel.Name, ""),
			AuthorName:  conve.StringDefault(novel.AuthorName, ""),
			NovelStatus: conve.Int16Default(novel.Status, 0),
			Tags:        make([]string, 0),
			Describes:   make([]string, 0),
		}
		if err := jsoniter.UnmarshalFromString(conve.StringDefault(novel.Tags, ""), &item.Tags); err != nil {
			return nil, err
		}
		if err := jsoniter.UnmarshalFromString(conve.StringDefault(novel.NovelAbs, ""), &item.Describes); err != nil {
			return nil, err
		}
		data.Items = append(data.Items, item)
	}
	return data, nil
}

func GetNovelTags(req *story_telling_backend.GetNovelTagsReq) (*story_telling_backend.GetNovelTagsData, error) {
	dbClient := db.GetDBClient()
	query := dbClient.Model(&db_model.BookTable{})
	if req.NovelID != nil {
		query = query.Where("id LIKE ?", conve.Int64Default(req.NovelID, 0))
	}
	var novels []db_model.BookTable
	if err := query.Find(&novels).Error; err != nil {
		return nil, err
	}

	tagSet := mapset.NewSet[string]()
	for _, novel := range novels {
		tags := make([]string, 0)
		if err := jsoniter.UnmarshalFromString(conve.StringDefault(novel.Tags, ""), &tags); err != nil {
			return nil, err
		}
		tagSet.Append(tags...)
	}

	data := &story_telling_backend.GetNovelTagsData{
		Tags: tagSet.ToSlice(),
	}
	return data, nil
}
