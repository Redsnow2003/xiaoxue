package module

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Role represents a user role
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// @Summary Create a new role
// @Description Create a new role
// @Tags roles
// @Accept json
// @Produce json
// @Param role body Role true "Role"
// @Success 200 {object} Role
// @Router /roles [post]
func CreateRole(c *gin.Context) {
	var role Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add role creation logic here
	c.JSON(http.StatusOK, role)
}

// @Summary Get a role by ID
// @Description Get a role by ID
// @Tags roles
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} Role
// @Router /roles/{id} [get]
func GetRole(c *gin.Context) {
	//id := c.Param("id")
	// Add role retrieval logic here
	role := Role{ID: 1, Name: "Admin"} // Example role
	c.JSON(http.StatusOK, role)
}

// @Summary Update a role
// @Description Update a role
// @Tags roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param role body Role true "Role"
// @Success 200 {object} Role
// @Router /roles/{id} [put]
func UpdateRole(c *gin.Context) {
	//id := c.Param("id")
	var role Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add role update logic here
	role.ID = 1 // Example role ID
	c.JSON(http.StatusOK, role)
}

// @Summary Delete a role
// @Description Delete a role
// @Tags roles
// @Param id path int true "Role ID"
// @Success 200 {string} string "ok"
// @Router /roles/{id} [delete]
func DeleteRole(c *gin.Context) {
	//id := c.Param("id")
	// Add role deletion logic here
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}