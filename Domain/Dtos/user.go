package dtos

type UserCreateRequest struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Status int    `json:"status"`
	RoleId string `json:"role_id"`
}

type UserUpdateRequest struct {
	Name    *string `json:"name,omitempty"`
	Email   *string `json:"email,omitempty"`
	Status  *int    `json:"status,omitempty"`
	RoleId  *string `json:"role_id,omitempty"`
	UserUID string  `json:"UserUID"`
}

type AddUserToGroupRequest struct {
	UserUID  string   `json:"UserUID"`
	GroupIds []string `json:"group_ids" binding:"required"`
}

type AddUserToRoleRequest struct {
	UserUID string `json:"UserUID"`
	RoleId  string `json:"role_id" binding:"required"`
}

type UserResponse struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status int    `json:"status"`
}

type UserResponseSingle struct {
	UID    string          `json:"uid"`
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Status int             `json:"status"`
	Groups []GroupResponse `json:"groups"`
	Role   *RoleResponse   `json:"role"`
}

type UserResponseAll struct {
	UID    string               `json:"uid"`
	Name   string               `json:"name"`
	Email  string               `json:"email"`
	Status int                  `json:"status"`
	Role   *RoleResponseNoRight `json:"role"`
}

type SearchFields struct {
	Name    string `json:"Name"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"orderBy"`
}

type RemoveUserFromGroupRequest struct {
	UserUID  string   `json:"UserUID"`
	GroupIds []string `json:"group_ids" binding:"required"`
}
