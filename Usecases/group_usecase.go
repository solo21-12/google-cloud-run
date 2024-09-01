package usecases

import (
	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type groupUseCase struct {
	groupRepo interfaces.GroupRepository
}

func NewGroupUseCase(groupRepo interfaces.GroupRepository) interfaces.GroupUseCase {
	return &groupUseCase{
		groupRepo: groupRepo,
	}
}

func (uc *groupUseCase) checkGroupExists(id string, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	group, err := uc.groupRepo.GetGroupById(id, ctx)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (uc *groupUseCase) GetAllGroups(ctx *gin.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	return uc.groupRepo.GetAllGroups(ctx)
}

func (uc *groupUseCase) GetGroupById(id string, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	return uc.checkGroupExists(id, ctx)
}

func (uc *groupUseCase) GetGroupUsers(id string, ctx *gin.Context) ([]dtos.UserResponse, *models.ErrorResponse) {
	_, err := uc.checkGroupExists(id, ctx)
	if err != nil {
		return nil, err
	}

	users, err := uc.groupRepo.GetGroupUsers(id, ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *groupUseCase) CreateGroup(group dtos.GroupCreateRequest, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	if group, err := uc.groupRepo.GetGroupByName(group.Name, ctx); err == nil && group != nil {
		return nil, models.BadRequest("Group with the given name already exists")
	}
	return uc.groupRepo.CreateGroup(group, ctx)
}

func (uc *groupUseCase) UpdateGroup(id string, group dtos.GroupUpdateRequest, ctx *gin.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	if _, err := uc.checkGroupExists(id, ctx); err != nil {
		return nil, err
	}
	return uc.groupRepo.UpdateGroup(id, group, ctx)
}

func (uc *groupUseCase) DeleteGroup(id string, ctx *gin.Context) *models.ErrorResponse {
	if _, err := uc.checkGroupExists(id, ctx); err != nil {
		return err
	}

	return uc.groupRepo.DeleteGroup(id, ctx)
}
