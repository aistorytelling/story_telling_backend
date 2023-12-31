// Code generated by hertz generator.

package story_telling_backend

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	log "github.com/sirupsen/logrus"
	"story_telling_backend/biz/model/story_telling_backend"
	"story_telling_backend/biz/service"
)

// SearchNovel .
// @router /story_telling/api/v1/novel/search [POST]
func SearchNovel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req story_telling_backend.SearchNovelReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(story_telling_backend.SearchNovelResp)
	if resp.Data, err = service.SearchNovel(&req); err != nil {
		log.WithFields(log.Fields{
			"interface": "SearchNovel",
			"err":       err.Error(),
		}).Error("请求失败")
		c.String(consts.StatusOK, "请求失败")
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetNovelDetail .
// @router /story_telling/api/v1/novel/detailvel_id [GET]
func GetNovelDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req story_telling_backend.GetNovelDetailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(story_telling_backend.GetNovelDetailResp)

	c.JSON(consts.StatusOK, resp)
}

// GetNovelTellingConfig .
// @router /story_telling/api/v1/telling/configvel_id [GET]
func GetNovelTellingConfig(ctx context.Context, c *app.RequestContext) {
	var err error
	var req story_telling_backend.GetNovelTellingConfigReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(story_telling_backend.GetNovelTellingConfigResp)

	if resp.Data, err = service.GetTellingConfig(&req); err != nil {
		log.WithFields(log.Fields{
			"interface": "GetNovelTellingConfig",
			"err":       err.Error(),
		}).Error("请求失败")
		c.String(consts.StatusOK, "请求失败")
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetNovelChapterTitle .
// @router /story_telling/api/v1/chapter/title_listvel_id [GET]
func GetNovelChapterTitle(ctx context.Context, c *app.RequestContext) {
	var err error
	var req story_telling_backend.GetNovelChapterTitleListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(story_telling_backend.GetNovelChapterTitleListResp)

	if resp.Data, err = service.GetChapterTitles(&req); err != nil {
		log.WithFields(log.Fields{
			"interface": "GetNovelChapterTitle",
			"err":       err.Error(),
		}).Error("请求失败")
		c.String(consts.StatusOK, "请求失败")
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetChapterDetail .
// @router /story_telling/api/v1/chapter/detailvel_id/:chapter_id [GET]
func GetChapterDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req story_telling_backend.GetChapterDetailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(story_telling_backend.GetChapterDetailResp)

	if resp.Data, err = service.GetChapterDetail(&req); err != nil {
		log.WithFields(log.Fields{
			"interface": "GetChapterDetail",
			"err":       err.Error(),
		}).Error("请求失败")
		c.String(consts.StatusOK, "请求失败")
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetNovelTags .
// @router /story_telling/api/v1/novel/tags [POST]
func GetNovelTags(ctx context.Context, c *app.RequestContext) {
	var err error
	var req story_telling_backend.GetNovelTagsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(story_telling_backend.GetNovelTagsResp)

	if resp.Data, err = service.GetNovelTags(&req); err != nil {
		log.WithFields(log.Fields{
			"interface": "GetNovelTags",
			"err":       err.Error(),
		}).Error("请求失败")
		c.String(consts.StatusOK, "请求失败")
		return
	}

	c.JSON(consts.StatusOK, resp)
}
