package dtos

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserToGroupRequest struct {
	UserId  uint `json:"user_id"`
	GroupId uint `json:"group_id"`
}

type AddUserToRoleRequest struct {
	UserId uint `json:"user_id"`
	RoleId uint `json:"role_id"`
}

type UserResponse struct {
	UID   uint   `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
