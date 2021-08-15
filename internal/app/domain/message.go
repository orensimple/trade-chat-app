package domain

// CreateMessageRequest struct for create message request.
type CreateMessageRequest struct {
	ChatID string `form:"chat,omitempty" json:"chat,omitempty"`
	UserID string `form:"user,omitempty" json:"user,omitempty"`
	Body   string `form:"body,omitempty" json:"body,omitempty"`
}

// SearchMessageRequest struct for search message request.
type SearchMessageRequest struct {
	UserID string `form:"user,omitempty" json:"user,omitempty"`
	ChatID string `form:"chat,omitempty" json:"chat,omitempty"`
}
