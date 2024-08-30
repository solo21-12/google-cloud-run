package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"

	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) interfaces.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetAllRoles(ctx context.Context) ([]*models.Role, *models.ErrorResponse) {
	var roles []*models.Role
	if err := r.db.WithContext(ctx).Preload("User").Find(&roles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*models.Role{}, nil
		}
		return nil, models.InternalServerError(err.Error())
	}
	return roles, nil
}

func (r *roleRepository) GetRoleById(id string, ctx context.Context) (*models.Role, *models.ErrorResponse) {
	roleId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var role models.Role
	if err := r.db.WithContext(ctx).Preload("User").First(&role, roleId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Role not found")
		}
		return nil, models.InternalServerError(err.Error())
	}
	return &role, nil
}

func (r *roleRepository) CreateRole(role dtos.RoleCreateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse) {
	newRole := models.Role{
		RID:    uuid.New(),
		Name:   role.Name,
		Rights: role.Rights,
	}
	if err := r.db.WithContext(ctx).Create(&newRole).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return &dtos.RoleResponse{
		RID:    newRole.RID.String(),
		Name:   newRole.Name,
		Rights: newRole.Rights,
	}, nil
}

func (r *roleRepository) UpdateRole(id string, role dtos.RoleUpdateRequest, ctx context.Context) (*dtos.RoleResponse, *models.ErrorResponse) {
	roleId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var existingRole models.Role
	if err := r.db.WithContext(ctx).First(&existingRole, roleId).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	existingRole.Name = role.Name
	existingRole.Rights = role.Rights

	if err := r.db.WithContext(ctx).Save(&existingRole).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.RoleResponse{
		RID:    existingRole.RID.String(),
		Name:   existingRole.Name,
		Rights: existingRole.Rights,
	}, nil
}

func (r *roleRepository) DeleteRole(id string, ctx context.Context) *models.ErrorResponse {
	roleId, err := uuid.Parse(id)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	if err := r.db.WithContext(ctx).Delete(&models.Role{}, roleId).Error; err != nil {
		return models.InternalServerError(err.Error())
	}
	return nil
}


func (r *roleRepository) GetRoleUsers(role *models.Role, ctx context.Context) ([]*models.User, *models.ErrorResponse) {
	var users []*models.User
	if err := r.db.WithContext(ctx).Model(&role).Association("User").Find(&users); err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return users, nil
}