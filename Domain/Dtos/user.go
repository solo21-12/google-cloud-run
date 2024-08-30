package dtos

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

type UserUpdateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

type AddUserToGroupRequest struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
}

type AddUserToRoleRequest struct {
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}

type UserResponse struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status int    `json:"status"`
}

type SearchFields struct {
	Search  string `json:"search"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"orderBy"`
}
