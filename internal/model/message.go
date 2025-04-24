package model

type Message struct {
	Content         string            `json:"content,omitempty"`
	UserName        string            `json:"username,omitempty"`
	AvatarUrl       string            `json:"avatar_url,omitempty"`
	Tts             bool              `json:"tts,omitempty"`
	Embeds          []Embed           `json:"embeds,omitempty"`
	AllowedMentions AllowedMentions   `json:"allowed_mentions,omitempty"`
	Components      []ComponentObject `json:"components,omitempty"`
	// Files 		[]File            `json:"files,omitempty"`
	PayloadJson string             `json:"payload_json,omitempty"`
	Attachments []AttachmentObject `json:"attachments,omitempty"`
	Flags       int64              `json:"flags,omitempty"`
	ThreadName  string             `json:"thread_name,omitempty"`
	AppliedTags []string           `json:"applied_tags,omitempty"`
	Poll        PollObject         `json:"poll,omitempty"`
}
