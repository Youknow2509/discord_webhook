package model

// type attachment object
type AttachmentObject struct {
	Question         PollMediaObject   `json:"question"`
	Answers          PollAnswerObject  `json:"answers"`
	Expiry           string            `json:"expiry,omitempty"`
	AllowMultiselect bool              `json:"allow_multiselect,omitempty"`
	LayoutType       int               `json:"layout_type,omitempty"`
	Results          PollResultsObject `json:"results,omitempty"`
}
