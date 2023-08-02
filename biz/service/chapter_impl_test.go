package service

import (
	"reflect"
	"story_telling_backend/biz/model/story_telling_backend"
	"testing"
)

func TestGetChapterDetail(t *testing.T) {
	type args struct {
		req *story_telling_backend.GetChapterDetailReq
	}
	tests := []struct {
		name    string
		args    args
		want    *story_telling_backend.GetChapterDetailData
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				req: &story_telling_backend.GetChapterDetailReq{
					NovelID:   1,
					ChapterID: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChapterDetail(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChapterDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChapterDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
