package model

// Allowed Mentions Structure
type AllowedMentions struct {
	Parse       []string `json:"parse,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Users       []string `json:"users,omitempty"`
	RepliedUser bool     `json:"replied_user,omitempty"`
}

// Mention Types
const (
	MentionTypeRole     = "roles"
	MentionTypeUser     = "users"
	MentionTypeEveryone = "everyone"
)
