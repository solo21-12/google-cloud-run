package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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

func (uc *userController) isSearch(srarchField dtos.SearchFields) bool {
	return srarchField.Name != "" || srarchField.OrderBy != "" || srarchField.Limit != 10
}

func (uc *userController) GetUsers(c *gin.Context) {
	searchFields := dtos.SearchFields{
		Name:    c.Query("name"),
		Limit:   10,
		OrderBy: c.Query("orderby"),
	}

	if limitParam := c.Query("limit"); limitParam != "" {
		if limit, err := strconv.Atoi(limitParam); err == nil {
			if limit > 0 {
				searchFields.Limit = limit
			} else {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Limit must be greater than zero"})
				return
			}
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
	}

	var users []*dtos.UserResponseAll
	var errResp *models.ErrorResponse

	if !uc.isSearch(searchFields) {
		users, errResp = uc.usecase.GetAllUsers(c)
		if users == nil {
		} else {
			users, errResp = uc.usecase.SearchUsers(searchFields, c)
		}

		if errResp != nil {
			c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
			return
		}

	}

	if len(users) == 0 {
		c.IndentedJSON(http.StatusOK, []string{})

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

	if len(users) == 0 {
		c.IndentedJSON(http.StatusOK, []string{})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func (uc *userController) handleCreateOrUpdateUser(c *gin.Context, isUpdate bool) {
	var userRequest interface{}
	id := c.Param("id")

	if isUpdate {
		userRequest = &dtos.UserUpdateRequest{}
		userRequest.(*dtos.UserUpdateRequest).UserUID = id
	} else {
		userRequest = &dtos.UserCreateRequest{}
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "one or more fields are missing"})
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

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uc *userController) AddUserToGroup(c *gin.Context) {
	var req dtos.AddUserToGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "one or more fields are missing"})
		return
	}

	id := c.Param("id")
	req.UserUID = id
	err, msg := uc.usecase.AddUserToGroup(req, c)

	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": msg})
}

func (uc *userController) DeletetUserFromGroup(c *gin.Context) {
	var req dtos.RemoveUserFromGroupRequest

	id := c.Param("id")

	if err := c.ShouldBind(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validator := validator.New()
	if err := validator.Struct(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserUID = id

	msg, nErr := uc.usecase.RemoveUserFromGroup(req, c)

	if nErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": nErr.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Message": msg})
}
