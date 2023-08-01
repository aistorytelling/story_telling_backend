package service

import "story_telling_backend/biz/model/story_telling_backend"

func GetTellingConfig(req *story_telling_backend.GetNovelTellingConfigReq) (*story_telling_backend.GetNovelTellingConfigData, error) {
	data := &story_telling_backend.GetNovelTellingConfigData{
		Style:  []string{"激烈", "昂扬", "轻柔"},
		Timbre: []string{"male", "female"},
	}
	return data, nil
}
