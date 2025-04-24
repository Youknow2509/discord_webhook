package model

// type Emoji
type EmojiObject struct {
	Id            string      `json:"id,omitempty"`
	Name          string      `json:"name,omitempty"`
	Roles         []string    `json:"roles,omitempty"`
	User          *UserObject `json:"user,omitempty"`
	RequireColons bool        `json:"require_colons,omitempty"`
	Managed       bool        `json:"managed,omitempty"`
	Animated      bool        `json:"animated,omitempty"`
	Available     bool        `json:"available,omitempty"`
}
