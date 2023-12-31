// Code generated by hertz generator. DO NOT EDIT.

package story_telling_backend

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	story_telling_backend "story_telling_backend/biz/handler/story_telling_backend"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_story_telling := root.Group("/story_telling", _story_tellingMw()...)
		{
			_api := _story_telling.Group("/api", _apiMw()...)
			{
				_v1 := _api.Group("/v1", _v1Mw()...)
				{
					_chapter := _v1.Group("/chapter", _chapterMw()...)
					{
						_detail := _chapter.Group("/detail", _detailMw()...)
						{
							_novel_id := _detail.Group("/:novel_id", _novel_idMw()...)
							_novel_id.GET("/:chapter_id", append(_getchapterdetailMw(), story_telling_backend.GetChapterDetail)...)
						}
					}
					{
						_title_list := _chapter.Group("/title_list", _title_listMw()...)
						_title_list.GET("/:novel_id", append(_getnovelchaptertitleMw(), story_telling_backend.GetNovelChapterTitle)...)
					}
				}
				{
					_novel := _v1.Group("/novel", _novelMw()...)
					_novel.POST("/search", append(_searchnovelMw(), story_telling_backend.SearchNovel)...)
					_novel.GET("/tags", append(_getnoveltagsMw(), story_telling_backend.GetNovelTags)...)
					{
						_detail0 := _novel.Group("/detail", _detail0Mw()...)
						_detail0.GET("/:novel_id", append(_getnoveldetailMw(), story_telling_backend.GetNovelDetail)...)
					}
				}
				{
					_telling := _v1.Group("/telling", _tellingMw()...)
					{
						_config := _telling.Group("/config", _configMw()...)
						_config.GET("/:novel_id", append(_getnoveltellingconfigMw(), story_telling_backend.GetNovelTellingConfig)...)
					}
				}
			}
		}
	}
}
