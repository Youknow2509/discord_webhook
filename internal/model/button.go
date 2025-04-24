package model

// button struct
type ButtonObject struct {
	Type       ComponentType     `json:"type"`
	Style      int               `json:"style,omitempty"`
	Lable      string            `json:"label,omitempty"`
	Emoji      EmojiObject       `json:"emoji,omitempty"`
	Components []ComponentObject `json:"components"`
	CustomId   string            `json:"custom_id,omitempty"`
	SkuId      string            `json:"sku_id,omitempty"`
	Url        string            `json:"url,omitempty"`
	Disabled   bool              `json:"disabled,omitempty"`
}

// button style enum
type ButtonStyle int

const (
	ButtonStylePrimary ButtonStyle = iota + 1
	ButtonStyleSecondary
	ButtonStyleSuccess
	ButtonStyleDanger
	ButtonStyleLink
	ButtonStylePremium
)
