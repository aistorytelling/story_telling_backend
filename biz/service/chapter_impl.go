package service

import (
	"github.com/huqiyii/facility/conve"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"story_telling_backend/biz/db"
	"story_telling_backend/biz/model/db_model"
	"story_telling_backend/biz/model/story_telling_backend"
	"story_telling_backend/biz/service/common"
)

func GetChapterTitles(req *story_telling_backend.GetNovelChapterTitleListReq) (*story_telling_backend.GetNovelChapterTitleListData, error) {
	dbClient := db.GetDBClient()
	query := dbClient.Model(&db_model.ChapterTable{})
	query = query.Where("book_id = ?", req.NovelID)

	var chapters = make([]*db_model.ChapterTable, 0)
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	pageNo, pageSize := common.GetPagination(&story_telling_backend.Pagination{PageNo: req.PageNo, PageSize: req.PageSize})
	if err := query.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&chapters).Error; err != nil {
		return nil, err
	}

	data := &story_telling_backend.GetNovelChapterTitleListData{
		Total: total,
		Items: make([]*story_telling_backend.GetNovelChapterTitleItem, 0),
	}
	for _, chapter := range chapters {
		item := &story_telling_backend.GetNovelChapterTitleItem{
			ID:    conve.Int64Default(chapter.ID, 0),
			Title: conve.StringDefault(chapter.ChapterTitle, ""),
			Ind:   conve.Int16Default(chapter.ChapterInd, 0),
		}
		data.Items = append(data.Items, item)
	}

	return data, nil
}

func GetChapterDetail(req *story_telling_backend.GetChapterDetailReq) (*story_telling_backend.GetChapterDetailData, error) {
	dbClient := db.GetDBClient()
	query := dbClient.Model(&db_model.ChapterTable{})
	query = query.Where("book_id = ? and id = ?", req.NovelID, req.ChapterID)

	var chapter db_model.ChapterTable
	if err := query.First(&chapter).Error; err != nil {
		return nil, err
	}
	var err error
	data := &story_telling_backend.GetChapterDetailData{
		Title:   conve.StringDefault(chapter.ChapterTitle, ""),
		TextUri: viper.GetString("chapter.file_host") + conve.StringDefault(chapter.TxtURI, ""),
	}
	if data.AudioUri, err = json2StringSlice(conve.StringDefault(chapter.AudioMaleURI, ""), viper.GetString("chapter.file_host")); err != nil {
		return nil, err
	}
	if data.FrontendUri, err = json2StringSlice(conve.StringDefault(chapter.AudioMaleFronted, ""), viper.GetString("chapter.file_host")); err != nil {
		return nil, err
	}
	if data.AudioDuration, err = json2Int64Slice(conve.StringDefault(chapter.AudioMaleLength, "")); err != nil {
		return nil, err
	}

	if conve.StringDefault(req.Timbre, "male") == "female" {
		if data.AudioUri, err = json2StringSlice(conve.StringDefault(chapter.AudioFemaleURI, ""), viper.GetString("chapter.file_host")); err != nil {
			return nil, err
		}
		if data.FrontendUri, err = json2StringSlice(conve.StringDefault(chapter.AudioFemaleFronted, ""), viper.GetString("chapter.file_host")); err != nil {
			return nil, err
		}
		if data.AudioDuration, err = json2Int64Slice(conve.StringDefault(chapter.AudioFemaleLength, "")); err != nil {
			return nil, err
		}
	}
	return data, nil
}

func json2StringSlice(s string, prefix string) ([]string, error) {
	var result []string
	if err := jsoniter.UnmarshalFromString(s, &result); err != nil {
		return nil, err
	}
	for i, _ := range result {
		result[i] = prefix + result[i]
	}
	return result, nil
}

func json2Int64Slice(s string) ([]int64, error) {
	var result []int64
	if err := jsoniter.UnmarshalFromString(s, &result); err != nil {
		return nil, err
	}
	return result, nil
}
