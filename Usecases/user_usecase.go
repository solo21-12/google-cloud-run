package usecases

import (
	"context"

	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type userUseCase struct {
	userRepo        interfaces.UserRepository
	roleRepo        interfaces.RoleRepository
	groupRepo       interfaces.GroupRepository
	passwordService interfaces.PasswordService
	emailService    interfaces.EmailService
}

func NewUserUseCase(
	userRepo interfaces.UserRepository,
	passwordService interfaces.PasswordService,
	emailService interfaces.EmailService,
	roleRepo interfaces.RoleRepository,
	groupRepo interfaces.GroupRepository,
) interfaces.UserUseCase {
	return &userUseCase{
		userRepo:        userRepo,
		passwordService: passwordService,
		emailService:    emailService,
		roleRepo:        roleRepo,
		groupRepo:       groupRepo,
	}
}

func (uc *userUseCase) checkEmailExists(email string, ctx context.Context) (*models.User, *models.ErrorResponse) {
	existingUser, err := uc.userRepo.GetUserByEmail(email, ctx)
	if existingUser != nil && err == nil {
		return existingUser, models.BadRequest("User already exists")
	}
	return nil, nil
}

func (uc *userUseCase) processPassword(password string) (string, *models.ErrorResponse) {
	if err := uc.passwordService.ValidatePasswordStrength(password); err != nil {
		return "", err
	}

	hashedPassword, err := uc.passwordService.EncryptPassword(password)
	if err != nil {
		return "", models.InternalServerError(err.Error())
	}

	return hashedPassword, nil
}

func (uc *userUseCase) validateEmail(email string) *models.ErrorResponse {
	if valid := uc.emailService.IsValidEmail(email); !valid {
		return models.BadRequest("Invalid email")
	}
	return nil
}

func (uc *userUseCase) GetAllUsers(ctx context.Context) ([]*models.User, *models.ErrorResponse) {
	return uc.userRepo.GetAllUsers(ctx)
}

func (uc *userUseCase) GetUserById(id string, ctx context.Context) (*models.User, *models.ErrorResponse) {
	return uc.userRepo.GetUserById(id, ctx)
}

func (uc *userUseCase) GetUsersGroup(id string, ctx context.Context) ([]*models.Group, *models.ErrorResponse) {
	return uc.userRepo.GetUsersGroups(id, ctx)
}

func (uc *userUseCase) SearchUsers(searchFields dtos.SearchFields, ctx context.Context) ([]*models.User, *models.ErrorResponse) {
	return uc.userRepo.SearchUsers(searchFields, ctx)
}

func (uc *userUseCase) CreateUser(user dtos.UserCreateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	if _, err := uc.checkEmailExists(user.Email, ctx); err != nil {
		return nil, err
	}

	if err := uc.validateEmail(user.Email); err != nil {
		return nil, err
	}

	hashedPassword, err := uc.processPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	return uc.userRepo.CreateUser(user, ctx)
}

func (uc *userUseCase) UpdateUser(id string, user dtos.UserUpdateRequest, ctx context.Context) (*dtos.UserResponse, *models.ErrorResponse) {
	userToUpdate, err := uc.userRepo.GetUserById(id, ctx)
	if userToUpdate == nil && err != nil {
		return nil, models.NotFound("User not found")
	}

	if user.Name != "" {
		userToUpdate.Name = user.Name
	}

	if user.Email != "" {
		if err := uc.validateEmail(user.Email); err != nil {
			return nil, err
		}
		userToUpdate.Email = user.Email
	}

	if user.Password != "" {
		hashedPassword, err := uc.processPassword(user.Password)
		if err != nil {
			return nil, err
		}
		userToUpdate.Password = hashedPassword
	}

	if user.Status != 0 {
		userToUpdate.Status = user.Status
	}

	return uc.userRepo.UpdateUser(id, userToUpdate, ctx)
}

func (uc *userUseCase) DeleteUser(id string, ctx context.Context) *models.ErrorResponse {
	_, err := uc.userRepo.GetUserById(id, ctx)
	if err != nil {
		return models.NotFound("User not found")
	}

	return uc.userRepo.DeleteUser(id, ctx)
}

func (uc *userUseCase) AddUserToGroup(req dtos.AddUserToGroupRequest, ctx context.Context) *models.ErrorResponse {
	_, err := uc.userRepo.GetUserById(req.UserId, ctx)
	if err != nil {
		return models.NotFound("User not found")
	}

	_, err = uc.groupRepo.GetGroupById(req.GroupId, ctx)
	if err != nil {
		return models.NotFound("Group not found")
	}

	return uc.userRepo.AddUserToGroup(req, ctx)
}

func (uc *userUseCase) AddUserToRole(req dtos.AddUserToRoleRequest, ctx context.Context) *models.ErrorResponse {
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
