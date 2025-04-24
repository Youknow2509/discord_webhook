package model

import (
	"time"
)

// Poll object
type PollObject struct {
	Question         PollMediaObject    `json:"question"`
	Answers          []PollAnswerObject `json:"answers"`
	Expiry           *time.Time         `json:"expiry,omitempty"`
	AllowMultiselect bool               `json:"allow_multiselect"`
	LayoutType       int                `json:"layout_type"`
	Results          *PollResultsObject `json:"results,omitempty"`
}

// Poll Media Object
type PollMediaObject struct {
	Text  string        `json:"text,omitempty"`
	Emoji *PartialEmoji `json:"emoji,omitempty"`
}

// Partial Emoji Object
type PartialEmoji struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}

// Poll Answer Object
type PollAnswerObject struct {
	AnswerID  int             `json:"answer_id"`
	PollMedia PollMediaObject `json:"poll_media"`
}

// Poll Results Object
type PollResultsObject struct {
	IsFinalized  bool                    `json:"is_finalized"`
	AnswerCounts []PollAnswerCountObject `json:"answer_counts"`
}

// Poll Answer Count Object
type PollAnswerCountObject struct {
	AnswerID int  `json:"answer_id"`
	Count    int  `json:"count"`
	MeVoted  bool `json:"me_voted,omitempty"`
}

// Poll Layout Type
type PollLayoutType int

const (
	PollLayoutDefault PollLayoutType = iota + 1
	PollLayoutLeaderboard
)
