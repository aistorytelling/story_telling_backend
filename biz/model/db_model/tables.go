package db_model

import "time"

type BookTable struct {
	ID         *uint64    `gorm:"primaryKey" json:"id"`
	Name       *string    `json:"name"`
	CoverURI   *string    `json:"cover_uri"`
	AuthorName *string    `json:"author_name"`
	NovelAbs   *string    `json:"novel_abs"`
	Tags       *string    `json:"tags"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	CreatedBy  *string    `json:"created_by"`
	Status     *int8      `json:"status"`
}

func (BookTable) TableName() string {
	return "booktable"
}

type ChapterTable struct {
	ID                 *uint64    `gorm:"primaryKey" json:"id"`
	ChapterInd         *int16     `json:"chapter_ind"`
	ChapterTitle       *string    `json:"chapter_title"`
	BookID             *int64     `json:"book_id"`
	TxtURI             *string    `json:"txt_uri"`
	TxtStatus          *int8      `json:"txt_status"`
	AbsURI             *string    `json:"abs_uri"`
	AbsStatus          *int8      `json:"abs_status"`
	AudioMaleURI       *string    `json:"audio_male_uri"`
	AudioMaleLength    *string    `json:"audio_male_length"`
	AudioMaleFronted   *string    `json:"audio_male_fronted"`
	AudioFemaleURI     *string    `json:"audio_female_uri"`
	AudioFemaleLength  *string    `json:"audio_female_length"`
	AudioFemaleFronted *string    `json:"audio_female_fronted"`
	AudioStatus        *int8      `json:"audio_status"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
}

func (ChapterTable) TableName() string {
	return "chaptertable"
}
