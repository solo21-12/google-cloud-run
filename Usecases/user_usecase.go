package usecases

import (
	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type userUseCase struct {
	userRepo     interfaces.UserRepository
	roleRepo     interfaces.RoleRepository
	groupRepo    interfaces.GroupRepository
	emailService interfaces.EmailService
}

func NewUserUseCase(
	userRepo interfaces.UserRepository,
	emailService interfaces.EmailService,
	roleRepo interfaces.RoleRepository,
	groupRepo interfaces.GroupRepository,
) interfaces.UserUseCase {
	return &userUseCase{
		userRepo:     userRepo,
		emailService: emailService,
		roleRepo:     roleRepo,
		groupRepo:    groupRepo,
	}
}

func (uc *userUseCase) CheckEmailExists(email string, ctx *gin.Context) (*models.User, *models.ErrorResponse) {
	existingUser, err := uc.userRepo.GetUserByEmail(email, ctx)
	if existingUser != nil && err == nil {
		return existingUser, models.BadRequest("User already exists")
	}
	return nil, nil
}

func (uc *userUseCase) ValidateEmail(email string) *models.ErrorResponse {
	if valid := uc.emailService.IsValidEmail(email); !valid {
		return models.BadRequest("Invalid email")
	}
	return nil
}

func (uc *userUseCase) GetAllUsers(ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse) {
	return uc.userRepo.GetAllUsers(ctx)
}

func (uc *userUseCase) GetUserById(id string, ctx *gin.Context) (*dtos.UserResponseSingle, *models.ErrorResponse) {
	return uc.userRepo.GetUserById(id, ctx)
}

func (uc *userUseCase) GetUsersGroup(id string, ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	return uc.userRepo.GetUsersGroups(id, ctx)
}

func (uc *userUseCase) SearchUsers(searchFields dtos.SearchFields, ctx *gin.Context) ([]*dtos.UserResponse, *models.ErrorResponse) {
	return uc.userRepo.SearchUsers(searchFields, ctx)
}

func (uc *userUseCase) CreateUser(user dtos.UserCreateRequest, ctx *gin.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	if _, err := uc.CheckEmailExists(user.Email, ctx); err != nil {
		return nil, err
	}

	if err := uc.ValidateEmail(user.Email); err != nil {
		return nil, err
	}

	return uc.userRepo.CreateUser(user, ctx)
}

func (uc *userUseCase) UpdateUser(id string, user dtos.UserUpdateRequest, ctx *gin.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	userToUpdate, err := uc.userRepo.GetUserById(id, ctx)
	if userToUpdate == nil && err != nil {
		return nil, models.NotFound("User not found")
	}

	if user.Name != "" {
		userToUpdate.Name = user.Name
	}

	if user.Email != "" {
		if err := uc.ValidateEmail(user.Email); err != nil {
			return nil, err
		}
		userToUpdate.Email = user.Email
	}

	if user.Status != 0 {
		userToUpdate.Status = user.Status
	}

	updateUs := &dtos.UserUpdateRequest{
		Name:   userToUpdate.Name,
		Email:  userToUpdate.Email,
		Status: userToUpdate.Status,
	}
	return uc.userRepo.UpdateUser(id, updateUs, ctx)
}

func (uc *userUseCase) DeleteUser(id string, ctx *gin.Context) *models.ErrorResponse {
	_, err := uc.userRepo.GetUserById(id, ctx)
	if err != nil {
		return models.NotFound("User not found")
	}

	return uc.userRepo.DeleteUser(id, ctx)
}

func (uc *userUseCase) AddUserToGroup(req dtos.AddUserToGroupRequest, ctx *gin.Context) *models.ErrorResponse {
	// Check if the user exists
	_, err := uc.userRepo.GetUserById(req.UserId, ctx)
	if err != nil {
		return models.NotFound("User not found")
	}

	// Check if the group exists
	_, err = uc.groupRepo.GetGroupById(req.GroupId, ctx)
	if err != nil {
		return models.NotFound("Group not found")
	}

	// Get the user's current groups
	groups, tErr := uc.userRepo.GetUsersGroups(req.UserId, ctx)
	if tErr != nil {
		return tErr
	}

	// Check if the user is already a member of the group
	for _, group := range groups {
		if group.GID == req.GroupId {
			return models.BadRequest("User is already a member of the group")
		}
	}

	// Add the user to the group
	return uc.userRepo.AddUserToGroup(req, ctx)
}

func (uc *userUseCase) AddUserToRole(req dtos.AddUserToRoleRequest, ctx *gin.Context) *models.ErrorResponse {
	_, err := uc.userRepo.GetUserById(req.UserId, ctx)
	if err != nil {
		return models.NotFound("User not found")
	}

	_, err = uc.roleRepo.GetRoleById(req.RoleId, ctx)
	if err != nil {
		return models.NotFound("Role not found")
	}
	return uc.userRepo.AddUserToRole(req, ctx)
}
