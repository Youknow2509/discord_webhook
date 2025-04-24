package model

type Message struct {
	Content         string            `json:"content"`
	UserName        string            `json:"username"`
	AvatarUrl       string            `json:"avatar_url"`
	Tts             bool              `json:"tts"`
	Embeds          []Embed           `json:"embeds"`
	AllowedMentions AllowedMentions   `json:"allowed_mentions"`
	Components      []ComponentObject `json:"components"`
	// Files 		[]File            `json:"files"`
	PayloadJson string             `json:"payload_json"`
	Attachments []AttachmentObject `json:"attachments"`
	Flags       int64              `json:"flags"`
	ThreadName  string             `json:"thread_name"`
	AppliedTags []string           `json:"applied_tags"`
	Poll        PollObject         `json:"poll"`
}
