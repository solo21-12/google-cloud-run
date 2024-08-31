package dtos

type GroupCreateRequest struct {
	Name string `json:"name"`
}

type GroupUpdateRequest struct {
	Name string `json:"name"`
}

type GroupResponse struct {
	GID   string    `json:"gid"`
	Name string `json:"name"`
}
