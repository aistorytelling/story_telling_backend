package service

import (
	"fmt"
	"github.com/huqiyii/facility/conve"
	jsoniter "github.com/json-iterator/go"
	"story_telling_backend/biz/db"
	"story_telling_backend/biz/model/db_model"
	"story_telling_backend/biz/model/story_telling_backend"
	"story_telling_backend/biz/service/common"
)

func SearchNovel(req *story_telling_backend.SearchNovelReq) (*story_telling_backend.SearchNovelData, error) {
	dbClient := db.GetDBClient()
	query := dbClient.Model(&db_model.BookTable{})
	if req.CustomValue != nil {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *req.CustomValue))
	}
	if len(req.Tags) > 0 {
		subConditions := dbClient.Model(&db_model.BookTable{})
		for _, tag := range req.Tags {
			subConditions = subConditions.Or("JSON_CONTAINS(tags, ?)", tag)
		}
		query = query.Where(subConditions)
	}
	var novels []db_model.BookTable
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	pageNo, pageSize := common.GetPagination(req.GetPagination())
	if err := query.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&novels).Error; err != nil {
		return nil, err
	}
	data := &story_telling_backend.SearchNovelData{
		Total: total,
		Items: make([]*story_telling_backend.SearchNovelItem, 0),
	}
	for _, novel := range novels {
		item := &story_telling_backend.SearchNovelItem{
			ID:         conve.Int64Default(novel.ID, -1),
			CoverUrl:   conve.StringDefault(novel.CoverURI, ""),
			NovelName:  conve.StringDefault(novel.Name, ""),
			AuthorName: conve.StringDefault(novel.AuthorName, ""),
			Tags:       make([]string, 0),
			Describes:  make([]string, 0),
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
