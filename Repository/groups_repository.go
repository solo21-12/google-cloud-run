package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
	"github.com/google-run-code/config"
)

type groupRepository struct {
	dbConfig *config.PostgresConfig
}

func NewGroupRepository(dbConfig *config.PostgresConfig) interfaces.GroupRepository {
	return &groupRepository{
		dbConfig: dbConfig,
	}
}

func (r *groupRepository) getDB(ctx *gin.Context) (*gorm.DB, error) {
	dbName := ctx.GetString("dbName")
	db, ok := r.dbConfig.GetDB(dbName)
	if ok != nil {
		return nil, models.InternalServerError("Failed to get database connection")
	}

	return db, nil
}

func (r *groupRepository) GetAllGroups(ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	var groups []*models.Group
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).Preload("Users").Find(&groups).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*dtos.GroupResponse{}, nil
		}
		return nil, models.InternalServerError(err.Error())
	}

	var result []*dtos.GroupResponse

	for _, group := range groups {
		result = append(result, &dtos.GroupResponse{
			UID:  group.UID.String(),
			Name: group.Name,
		})

	}
	return result, nil
}

func (r *groupRepository) GetGroupById(UID string, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	var group models.Group
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).Preload("Users").
		Where("uid = ?", UID).
		First(&group).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Group not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	result := dtos.GroupResponse{
		UID:  group.UID.String(),
		Name: group.Name,
	}
	return &result, nil
}

func (r *groupRepository) GetGroupUsers(UID string, ctx *gin.Context) ([]dtos.UserResponse, *models.ErrorResponse) {
	var group models.Group
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).
		Preload("Users").
		Where("uid = ?", UID).
		First(&group).Error; err != nil {
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

func (r *groupRepository) GetGroupByName(name string, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	var group models.Group
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).Where("name = ?", name).Preload("Users").First(&group).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.NotFound("Group not found")
		}
		return nil, models.InternalServerError(err.Error())
	}

	result := dtos.GroupResponse{
		UID:  group.UID.String(),
		Name: group.Name,
	}
	return &result, nil
}

func (r *groupRepository) CreateGroup(group dtos.GroupCreateRequest, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	newGroup := models.Group{
		UID:  uuid.New(),
		Name: group.Name,
	}
	if err := db.WithContext(ctx).Create(&newGroup).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}
	return &dtos.GroupResponse{
		UID:  newGroup.UID.String(),
		Name: newGroup.Name,
	}, nil
}

func (r *groupRepository) UpdateGroup(UID string, group dtos.GroupUpdateRequest, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	var existingGroup models.Group
	db, err := r.getDB(ctx)

	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	if err := db.WithContext(ctx).
		Where("uid = ?", UID).
		First(&existingGroup).Error; err != nil {

		return nil, models.InternalServerError(err.Error())
	}

	existingGroup.Name = group.Name

	if err := db.WithContext(ctx).Save(&existingGroup).Error; err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return &dtos.GroupResponse{UID: existingGroup.UID.String(), Name: existingGroup.Name}, nil
}
func (r *groupRepository) DeleteGroup(id string, ctx *gin.Context) *models.ErrorResponse {
	// Parse the UUID from the string
	db, err := r.getDB(ctx)

	if err != nil {
		return models.InternalServerError(err.Error())
	}

	UID, err := uuid.Parse(id)
	if err != nil {
		return models.InternalServerError("Invalid UUID format")
	}

	// Fetch the group using uid
	var group models.Group
	if err := db.WithContext(ctx).
		Where("uid = ?", UID). // Use uid to fetch the group
		First(&group).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.NotFound("Group not found")
		}
		return models.InternalServerError("Failed to fetch group: " + err.Error())
	}

	// Dissociate users from the group
	if err := db.WithContext(ctx).Model(&group).Association("Users").Clear(); err != nil {
		return models.InternalServerError("Failed to clear group associations: " + err.Error())
	}

	// Delete the group
	if err := db.WithContext(ctx).
		Where("uid = ?", UID). // Use uid for deletion
		Delete(&group).Error; err != nil {
		return models.InternalServerError("Failed to delete group: " + err.Error())
	}

	return nil
}
