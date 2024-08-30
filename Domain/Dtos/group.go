package dtos

type GroupCreateRequest struct {
	Name string `json:"name"`
}

type GroupUpdateRequest struct {
	Name string `json:"name"`
}

type GroupResponse struct {
	ID   uint    `json:"id"`
	Name string `json:"name"`
}
