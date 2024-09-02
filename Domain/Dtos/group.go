package dtos

type GroupCreateRequest struct {
	Name string `json:"name"`
}

type GroupUpdateRequest struct {
	Name string `json:"name"`
}

type GroupResponse struct {
	UID   string    `json:"uid"`
	Name string `json:"name"`
}
