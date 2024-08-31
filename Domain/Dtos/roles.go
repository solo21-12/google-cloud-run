package dtos

import models "github.com/google-run-code/Domain/Models"

type RoleCreateRequest struct {
	Name   string        `json:"name"`
	Rights models.Rights `json:"rights"`
}

type RoleUpdateRequest struct {
	Name   string        `json:"name"`
	Rights models.Rights `json:"rights"`
}

type RoleResponse struct {
	RID     string          `json:"rid"`
	Name   string        `json:"name"`
	Rights models.Rights `json:"rights"`
}
