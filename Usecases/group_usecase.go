package usecases

import (
	"context"

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

func (uc *groupUseCase) checkGroupExists(id string, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	group, err := uc.groupRepo.GetGroupById(id, ctx)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (uc *groupUseCase) GetAllGroups(ctx context.Context) ([]*dtos.GroupResponse, *models.ErrorResponse) {
	return uc.groupRepo.GetAllGroups(ctx)
}

func (uc *groupUseCase) GetGroupById(id string, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	return uc.checkGroupExists(id, ctx)
}

func (uc *groupUseCase) GetGroupUsers(id string, ctx context.Context) ([]dtos.UserResponse, *models.ErrorResponse) {
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

func (uc *groupUseCase) CreateGroup(group dtos.GroupCreateRequest, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	return uc.groupRepo.CreateGroup(group, ctx)
}

func (uc *groupUseCase) UpdateGroup(id string, group dtos.GroupUpdateRequest, ctx context.Context) (*dtos.GroupResponse, *models.ErrorResponse) {
	if _, err := uc.checkGroupExists(id, ctx); err != nil {
		return nil, err
	}
	return uc.groupRepo.UpdateGroup(id, group, ctx)
}

func (uc *groupUseCase) DeleteGroup(id string, ctx context.Context) *models.ErrorResponse {
	if _, err := uc.checkGroupExists(id, ctx); err != nil {
		return err
	}

	return uc.groupRepo.DeleteGroup(id, ctx)
}
