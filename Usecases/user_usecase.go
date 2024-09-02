package usecases

import (
	"strings"

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

func (uc *userUseCase) GetAllUsers(ctx *gin.Context) ([]*dtos.UserResponseAll, *models.ErrorResponse) {
	return uc.userRepo.GetAllUsers(ctx)
}

func (uc *userUseCase) GetUserById(id string, ctx *gin.Context) (*dtos.UserResponseSingle, *models.ErrorResponse) {
	return uc.userRepo.GetUserById(id, ctx)
}

func (uc *userUseCase) GetUsersGroup(id string, ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	return uc.userRepo.GetUsersGroups(id, ctx)
}

func (uc *userUseCase) SearchUsers(searchFields dtos.SearchFields, ctx *gin.Context) ([]*dtos.UserResponseAll, *models.ErrorResponse) {
	return uc.userRepo.SearchUsers(searchFields, ctx)
}

func (uc *userUseCase) CreateUser(user dtos.UserCreateRequest, ctx *gin.Context) (*dtos.UserResponseSingle, *models.ErrorResponse) {
	if _, err := uc.CheckEmailExists(user.Email, ctx); err != nil {
		return nil, err
	}

	if err := uc.ValidateEmail(user.Email); err != nil {
		return nil, err
	}

	newUser, nErr := uc.userRepo.CreateUser(user, ctx)

	if nErr != nil {
		return nil, nErr
	}

	return uc.userRepo.GetUserById(newUser.UID, ctx)
}

func (uc *userUseCase) UpdateUser(id string, user dtos.UserUpdateRequest, ctx *gin.Context) (*dtos.UserResponseSingle, *models.ErrorResponse) {
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

	if user.RoleId != "" {
		err := uc.AddUserToRole(dtos.AddUserToRoleRequest{
			UserUID: userToUpdate.UID,
			RoleId:  user.RoleId,
		}, ctx)
		if err != nil {
			return nil, err
		}
	}

	updateUs := &dtos.UserUpdateRequest{
		Name:    userToUpdate.Name,
		Email:   userToUpdate.Email,
		Status:  userToUpdate.Status,
		UserUID: userToUpdate.UID,
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

func (uc *userUseCase) AddUserToGroup(req dtos.AddUserToGroupRequest, ctx *gin.Context) (*models.ErrorResponse, string) {
	// Check if the user exists
	_, err := uc.userRepo.GetUserById(req.UserUID, ctx)
	if err != nil {
		return models.NotFound("User not found"), ""
	}

	// Check which groups exist and filter out non-existent ones
	var validGroups []string
	for _, groupId := range req.GroupIds {
		_, err = uc.groupRepo.GetGroupById(groupId, ctx)
		if err == nil {
			validGroups = append(validGroups, groupId)
		}
	}
	if len(validGroups) == 0 {
		return models.NotFound("None of the specified groups were found"), ""
	}

	// Get the user's current groups
	groups, tErr := uc.userRepo.GetUsersGroups(req.UserUID, ctx)
	if tErr != nil {
		return tErr, ""
	}

	// Filter out groups where the user is already a member
	var newGroups []string
	var existingGroups []string
	for _, groupId := range validGroups {
		isMember := false
		for _, group := range groups {
			if group.UID == groupId {
				existingGroups = append(existingGroups, groupId)
				isMember = true
				break
			}
		}
		if !isMember {
			newGroups = append(newGroups, groupId)
		}
	}

	// If there are new groups to add the user to
	successMessage := ""
	if len(newGroups) > 0 {
		req.GroupIds = newGroups
		addErr := uc.userRepo.AddUserToGroup(req, ctx)
		if addErr != nil {
			return addErr, ""
		}

		successMessage += "User has been added to the given groups: "
	}

	// Create a success message that includes both new and existing group info
	if len(existingGroups) > 0 {
		successMessage += " The user was already a member of the following groups: " + strings.Join(existingGroups, ", ")
	}

	return nil, successMessage
}

func (uc *userUseCase) AddUserToRole(req dtos.AddUserToRoleRequest, ctx *gin.Context) *models.ErrorResponse {

	_, err := uc.roleRepo.GetRoleById(req.RoleId, ctx)
	if err != nil {
		return models.NotFound("Role not found")
	}
	return uc.userRepo.AddUserToRole(req, ctx)
}

func (uc *userUseCase) RemoveUserFromGroup(req dtos.RemoveUserFromGroupRequest, ctx *gin.Context) (string, *models.ErrorResponse) {
	_, err := uc.userRepo.GetUserById(req.UserUID, ctx)
	if err != nil {
		return "", models.NotFound("User not found")
	}

	var validGroupUIDs []string
	var nonMemberGroups []string

	for _, groupId := range req.GroupIds {
		groupResponse, err := uc.groupRepo.GetGroupById(groupId, ctx)
		if err == nil {
			validGroupUIDs = append(validGroupUIDs, groupResponse.UID)
		} else {
			nonMemberGroups = append(nonMemberGroups, groupId)
		}
	}

	if len(validGroupUIDs) == 0 {
		return "", models.NotFound("None of the specified groups were found")
	}

	userGroups, tErr := uc.userRepo.GetUsersGroups(req.UserUID, ctx)
	if tErr != nil {
		return "", tErr
	}

	var groupsToRemove []string
	var nonExistingGroups []string

	userGroupUIDs := make(map[string]struct{})
	for _, userGroup := range userGroups {
		userGroupUIDs[userGroup.UID] = struct{}{}
	}

	for _, groupUID := range validGroupUIDs {
		if _, isMember := userGroupUIDs[groupUID]; isMember {
			groupsToRemove = append(groupsToRemove, groupUID)
		} else {
			nonExistingGroups = append(nonExistingGroups, groupUID)
		}
	}

	if len(groupsToRemove) == 0 {
		return "", models.BadRequest("User is not a member of any of the specified groups")
	}

	if err := uc.userRepo.RemoveUserFromGroups(req.UserUID, groupsToRemove, ctx); err != nil {
		return "", models.InternalServerError(err.Error())
	}

	successMessage := "User has been removed from the following groups: " + strings.Join(groupsToRemove, ", ")
	if len(nonExistingGroups) > 0 {
		successMessage += ". The user was not a member of the following groups: " + strings.Join(nonExistingGroups, ", ")
	}

	return successMessage, nil
}
