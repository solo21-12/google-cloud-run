package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtos "github.com/google-run-code/Domain/Dtos"
	interfaces "github.com/google-run-code/Domain/Interfaces"
)

type roleController struct {
	roleUsecase interfaces.RoleUseCase
}

func NewRoleController(roleUsecase interfaces.RoleUseCase) interfaces.RoleController {
	return &roleController{
		roleUsecase: roleUsecase,
	}
}

func (rc *roleController) GetAllRoles(c *gin.Context) {
	roles, errResp := rc.roleUsecase.GetAllRoles(c.Request.Context())

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, roles)
}

func (rc *roleController) GetRoleById(c *gin.Context) {
	id := c.Param("id")
	role, errResp := rc.roleUsecase.GetRoleById(id, c.Request.Context())

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, role)
}

func (rc *roleController) CreateRole(c *gin.Context) {
	var role dtos.RoleCreateRequest

	if err := c.ShouldBindJSON(&role); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdRole, errResp := rc.roleUsecase.CreateRole(role, c.Request.Context())

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusCreated, createdRole)

}

func (rc *roleController) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role dtos.RoleUpdateRequest

	if err := c.ShouldBindJSON(&role); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRole, errResp := rc.roleUsecase.UpdateRole(id, role, c.Request.Context())

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedRole)
}

func (rc *roleController) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	errResp := rc.roleUsecase.DeleteRole(id, c.Request.Context())

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Role deleted successfully"})
}

func (rc *roleController) GetRoleUsers(c *gin.Context) {
	id := c.Param("id")

	users, errResp := rc.roleUsecase.GetRoleUsers(id, c.Request.Context())

	if errResp != nil {
		c.IndentedJSON(errResp.Code, gin.H{"error": errResp.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
