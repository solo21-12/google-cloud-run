package usecases

import (
	"context"

	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type roleUseCase struct {
	roleRepository interfaces.RoleRepository
}

func NewRoleUseCase(roleRepository interfaces.RoleRepository) interfaces.RoleUseCase {
	return &roleUseCase{
		roleRepository: roleRepository,
	}
}

func (uc *roleUseCase) checkRoleExists(id string, ctx context.Context) (*models.Role, *models.ErrorResponse) {
	role, err := uc.roleRepository.GetRoleById(id, ctx)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (uc *roleUseCase) GetAllRoles(ctx context.Context) ([]*models.Role, *models.ErrorResponse) {
	return uc.roleRepository.GetAllRoles(ctx)
}

func (uc *roleUseCase) GetRoleById(id string, ctx context.Context) (*models.Role, *models.ErrorResponse) {
	return uc.checkRoleExists(id, ctx)
}

func (uc *roleUseCase) CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse) {
	return uc.roleRepository.CreateRole(role, ctx)
}

func (uc *roleUseCase) UpdateRole(id string, role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse) {
	if _, err := uc.checkRoleExists(id, ctx); err != nil {
		return nil, err
	}
	return uc.roleRepository.UpdateRole(id, role, ctx)
}

func (uc *roleUseCase) DeleteRole(id string, ctx context.Context) *models.ErrorResponse {
	if _, err := uc.checkRoleExists(id, ctx); err != nil {
		return err
	}
	return uc.roleRepository.DeleteRole(id, ctx)
}


func (uc *roleUseCase) GetRoleUsers(id string, ctx context.Context) ([]*models.User, *models.ErrorResponse) {
	role, err := uc.checkRoleExists(id, ctx)
	if err != nil {
		return nil, err
	}
	return uc.roleRepository.GetRoleUsers(role, ctx)
}