package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"

	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) interfaces.GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) GetAllGroups(ctx context.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	var groups []*models.Group
	if err := r.db.WithContext(ctx).Preload("Users").Find(&groups).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*dtos.GroupResponse{}, nil
		}
		return nil, models.InternalServerError(err.Error())
	}

	var result []*dtos.GroupResponse

	for _, group := range groups {
		result = append(result, &dtos.GroupResponse{
			GID:  group.GID.String(),
			Name: group.Name,
		})

	}
	return result, nil
}

func (r *GroupRepository) GetGroupById(id string, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	gId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var group models.Group
	if err := r.db.WithContext(ctx).Preload("Users").First(&group, gId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Group not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	result := dtos.GroupResponse{
		GID:  group.GID.String(),
		Name: group.Name,
	}
	return &result, nil
}

func (r *GroupRepository) GetGroupUsers(id string, ctx context.Context) ([]dtos.UserResponse, *models.ErrorResponse) {
	var group models.Group
	gId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}

	if err := r.db.WithContext(ctx).Preload("Users").First(&group, gId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Group not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	var users []dtos.UserResponse

	for _, user := range group.Users {
		users = append(users, dtos.UserResponse{
			UID:    user.UID.String(),
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
		})
	}

	return users, nil
}

func (r *GroupRepository) GetGroupByName(name string, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	var group models.Group
	if err := r.db.WithContext(ctx).Where("name = ?", name).Preload("Users").First(&group).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Group not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	result := dtos.GroupResponse{
		GID:  group.GID.String(),
		Name: group.Name,
	}
	return &result, nil
}



func (r *GroupRepository) CreateGroup(group dtos.GroupCreateRequest, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	newGroup := models.Group{
		GID:  uuid.New(),
		Name: group.Name,
	}
	if err := r.db.WithContext(ctx).Create(&newGroup).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return &dtos.GroupResponse{
		GID:  newGroup.GID.String(),
		Name: newGroup.Name,
	}, nil
}

func (r *GroupRepository) UpdateGroup(id string, group dtos.GroupUpdateRequest, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	gId, err := uuid.Parse(id)
	if err != nil {
		return nil, models.InternalServerError("Invalid UUID format")
	}
	var existingGroup models.Group
	if err := r.db.WithContext(ctx).First(&existingGroup, gId).Error; err != nil {

		return nil, models.InternalServerError(err.Error())
	}

	existingGroup.Name = group.Name

	if err := r.db.WithContext(ctx).Save(&existingGroup).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.GroupResponse{GID: existingGroup.GID.String(), Name: existingGroup.Name}, nil
}

func (r *GroupRepository) DeleteGroup(id string, ctx context.Context) *models.ErrorResponse {
	gId, err := uuid.Parse(id)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	var group models.Group
	if err := r.db.WithContext(ctx).First(&group, gId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("Group not found")
		}
		return models.InternalServerError(err.Error())
	}

	if err := r.db.WithContext(ctx).Model(&group).Association("Users").Clear(); err != nil {
		return models.InternalServerError("Failed to clear group associations")
	}

	if err := r.db.WithContext(ctx).Delete(&group).Error; err != nil {
		return models.InternalServerError("Failed to delete group")
	}

	return nil
}
