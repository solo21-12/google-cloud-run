package dtos

import "encoding/json"

type RoleCreateRequest struct {
	Name   string          `json:"name" binding:"required"`
	Rights json.RawMessage `json:"rights" binding:"required"`
}

type RoleUpdateRequest struct {
	Name   string          `json:"name"`
	Rights json.RawMessage `json:"rights"`
}

type RoleResponse struct {
	UID    string          `json:"uid"`
	Name   string          `json:"name"`
	Rights json.RawMessage `json:"rights"`
}


type RoleResponseNoRight struct {
	UID    string          `json:"uid"`
	Name   string          `json:"name"`
}
