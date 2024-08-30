package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
	models "github.com/google-run-code/Domain/Models"
)

type userController struct {
	usecase interfaces.UserUseCase
}

func NewUserController(usecase interfaces.UserUseCase) interfaces.UserController {
	return &userController{
		usecase: usecase,
	}
}

func (uc *userController) GetAllUsers(c *gin.Context) {
	users, err := uc.usecase.GetAllUsers(c)

	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (uc *userController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.usecase.GetUserById(id, c)

	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (uc *userController) GetUsersGroup(c *gin.Context) {
	id := c.Param("id")
	users, err := uc.usecase.GetUsersGroup(id, c)

	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (uc *userController) SearchUsers(c *gin.Context) {
	// Implement logic to search users
}

func (uc *userController) handleCreateOrUpdateUser(c *gin.Context, isUpdate bool) {
	var userRequest interface{}
	id := c.Param("id")

	if isUpdate {
		userRequest = &dtos.UserUpdateRequest{}
	} else {
		userRequest = &dtos.UserCreateRequest{}
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userResponse interface{}
	var err *models.ErrorResponse

	if isUpdate {
		userResponse, err = uc.usecase.UpdateUser(id, *(userRequest.(*dtos.UserUpdateRequest)), c)
	} else {
		userResponse, err = uc.usecase.CreateUser(*(userRequest.(*dtos.UserCreateRequest)), c)
	}

	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	status := http.StatusOK
	if !isUpdate {
		status = http.StatusCreated
	}

	c.IndentedJSON(status, userResponse)
}

func (uc *userController) CreateUser(c *gin.Context) {
	uc.handleCreateOrUpdateUser(c, false)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	uc.handleCreateOrUpdateUser(c, true)
}

func (uc *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := uc.usecase.DeleteUser(id, c)

	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
