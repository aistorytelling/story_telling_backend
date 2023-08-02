namespace go story_telling_backend

struct GetNovelTagsReq {
    1: optional i64 NovelID (api.path="novel_id") // 不传则获取所有的tags枚举
}

struct GetNovelTagsData {
    1: required list<string> Tags (api.path="tags")
}

struct GetNovelTagsResp {
    1: required i32 Code (api.body="code")
    2: required string Message (api.body="message")
    3: optional GetNovelTagsData Data (api.body="data")
}

struct Pagination {
    1: required i32 PageSize (api.body="page_size") // 大于0，默认值10
    2: required i32 PageNo (api.body="page_no") // 大于0，默认值1
}

struct SearchNovelReq {
    1: optional string CustomValue (api.body="custom_value") // 非必填，不传不用于筛选
    2: optional list<string> Tags (api.body="tags") // 非必填，不传或长度为空不用于筛选
    3: optional Pagination pagination (api.body="pagination")
}

struct SearchNovelItem {
    1: required i64 ID (api.body="id") // 小说id
    2: required string CoverUrl (api.body="cover_url") // 封面图url, 这里用编码后的图片文件还是url呢？
    3: required string NovelName (api.body="novel_name") // 小说名字
    4: required string AuthorName (api.body="author_name") // 作者名字
    5: required list<string> Tags (api.body="tags") // 小说标签
    6: required list<string> Describes (api.body="describes") // 一些七七八八的描述
}

struct SearchNovelData {
    1: required list<SearchNovelItem> Items (api.body="items")
    2: required i64 Total (api.body="total")
}

struct SearchNovelResp {
    1: required i32 Code (api.body="code")
    2: required string Message (api.body="message")
    3: optional SearchNovelData Data (api.body="data")
}

struct GetNovelDetailReq {
}

struct GetNovelDetailResp {
}

struct GetNovelChapterTitleListReq {
    1: required i64 NovelID (api.path="novel_id")
    2: required i32 PageSize (api.query="page_size") // 大于0，默认值10
    3: required i32 PageNo (api.query="page_no") // 大于0，默认值1
}

struct GetNovelChapterTitleItem {
    1: required i64 ID (api.body="id") // 章节id
    2: required string Title (api.body="title") // 章节标题
    3: required i16 Ind (api.body="ind") // 章节号
}

struct GetNovelChapterTitleListData {
    1: required list<GetNovelChapterTitleItem> Items (api.body="items")
    2: required i64 Total (api.body="total")
}

struct GetNovelChapterTitleListResp {
    1: required i32 Code (api.body="code")
    2: required string Message (api.body="message")
    3: optional GetNovelChapterTitleListData Data (api.body="data")
}

struct GetNovelTellingConfigReq {
    1: required i64 NovelID (api.path="novel_id")
}

struct GetNovelTellingConfigData { // 这里应该枚举，枚举值待定
    1: required list<string> Style (api.body="style") // 风格
    2: required list<string> Timbre (api.body="timbre") // 音色
}

struct GetNovelTellingConfigResp {
    1: required i32 Code (api.body="code")
    2: required string Message (api.body="message")
    3: optional GetNovelTellingConfigData Data (api.body="data")
}

struct GetChapterDetailReq {
    1: required i64 NovelID (api.path="novel_id")
    2: required i64 ChapterID (api.path="chapter_id")
    3: optional string Style (api.quest="style") // 风格, 不传或错误值使用默认值
    4: optional string Timbre (api.quest="timbre") // 音色， 不传或错误值使用默认值
}

struct GetChapterDetailData {
    1: required string Title (api.body="title") // 章节标题
    2: required string FrontendUri (api.body="frontend_uri")
    3: required string AudioUri (api.body="audio_uri") // 音频，mp3格式，需要decode
    4: required i64 AudioDuration (api.body="audio_duration") // 音频时长
    5: required string TextUri (api.body="text_uri") // 文本地址
}

struct GetChapterDetailResp {
    1: required i32 Code (api.body="code")
    2: required string Message (api.body="message")
    3: optional GetChapterDetailData Data (api.body="data")
}

service NovelBackendService {
    // novel
    GetNovelTagsResp GetNovelTags(1: GetNovelTagsReq request) (api.get="/story_telling/api/v1/novel/tags");
    SearchNovelResp SearchNovel(1: SearchNovelReq request) (api.post="/story_telling/api/v1/novel/search");
    GetNovelDetailResp GetNovelDetail(1: GetNovelDetailReq request) (api.get="/story_telling/api/v1/novel/detail/:novel_id");
    GetNovelTellingConfigResp GetNovelTellingConfig(1: GetNovelTellingConfigReq request) (api.get="/story_telling/api/v1/telling/config/:novel_id");
    // chapter
    GetNovelChapterTitleListResp GetNovelChapterTitle(1: GetNovelChapterTitleListReq request) (api.get="/story_telling/api/v1/chapter/title_list/:novel_id");
    GetChapterDetailResp GetChapterDetail(1: GetChapterDetailReq request) (api.get="/story_telling/api/v1/chapter/detail/:novel_id/:chapter_id");
}