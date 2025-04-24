package model

type ComponentType int

const (
	ActionRow ComponentType = iota + 1
	Button
	StringSelect
	TextInput
	UserSelect
	RoleSelect
	MentionableSelect
	ChannelSelect
)

// component item
type ComponentItem struct {
	Type       ComponentType     `json:"type"`
	Components []ComponentObject `json:"components"`
	Lable      string            `json:"label,omitempty"`
	Style      int               `json:"style,omitempty"`
	CustomId   string            `json:"custom_id,omitempty"`
}

// component object
type ComponentObject struct {
	Content    string          `json:"content,omitempty"`
	Components []ComponentItem `json:"components,omitempty"`
}


