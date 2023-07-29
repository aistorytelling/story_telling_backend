package service

import "story_telling_backend/biz/model/story_telling_backend"

func GetTellingConfig(req *story_telling_backend.GetNovelTellingConfigReq) (*story_telling_backend.GetNovelTellingConfigData, error) {
	data := &story_telling_backend.GetNovelTellingConfigData{
		Style:  []string{"激烈", "昂扬", "轻柔"},
		Timbre: []string{"男", "女"},
	}
	return data, nil
}
