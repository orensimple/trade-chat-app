package domain

// CreateChatRequest struct for create chat request.
type CreateChatRequest struct {
	UserIDs []string `form:"users,omitempty" json:"users,omitempty"`
	Label   string   `form:"label,omitempty" json:"label,omitempty"`
}

// SearchChatRequest struct for search chat request.
type SearchChatRequest struct {
	UserIDs []string `form:"users,omitempty" json:"users,omitempty"`
}
