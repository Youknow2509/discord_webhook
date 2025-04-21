package model

import "time"

// embed type
const (
	EmbedTypeRich       = "rich"
	EmbedTypeImage      = "image"
	EmbedTypeVideo      = "video"
	EmbedTypeGifv       = "gifv"
	EmbedTypeArticle    = "article"
	EmbedTypeLink       = "link"
	EmbedTypePollResult = "poll_result"
)

// embed struct
type Embed struct {
	Title       string         `json:"title,omitempty"`
	Type        string         `json:"type,omitempty"` // always "rich" for webhook embeds
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Timestamp   time.Time      `json:"timestamp,omitempty"` // ISO8601 timestamp
	Color       int            `json:"color,omitempty"`
	Footer      EmbedFooter    `json:"embed_footer,omitempty"`
	Image       EmbedImage     `json:"embed_image,omitempty"`
	Thumbnail   EmbedThumbnail `json:"embed_thumbnail,omitempty"`
	Video       EmbedVideo     `json:"embed_video,omitempty"`
	Provider    EmbedProvider  `json:"embed_provider,omitempty"`
	Author      EmbedAuthor    `json:"embed_author,omitempty"`
	Fields      []EmbedField   `json:"fields,omitempty"` // max of 25
}

// EmbedFooter struct
type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedImage struct
type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// EmbedThumbnail struct
type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// EmbedVideo struct
type EmbedVideo struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// EmbedProvider struct
type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// EmbedAuthor struct
type EmbedAuthor struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedField struct
type EmbedField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// Poll Result Embed Fields
type PollResultEmbedField struct {
	PollQuestionText          string `json:"poll_question_text,omitempty"`
	VictorAnswerVotes         int    `json:"victor_answer_votes,omitempty"`
	TotalVotes                int    `json:"total_votes,omitempty"`
	VictorAnswerId            string `json:"victor_answer_id,omitempty"`
	VictorAnswerText          string `json:"victor_answer_text,omitempty"`
	VictorAnswerEmojiId       string `json:"victor_answer_emoji_id,omitempty"`
	VictorAnswerEmojiName     string `json:"victor_answer_emoji_name,omitempty"`
	VictorAnswerEmojiAnimated string `json:"victor_answer_emoji_animated,omitempty"`
}
