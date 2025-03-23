package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MSaaS-Framework/MSaaS/pkg/base"
	"MSaaS-Framework/MSaaS/pkg/core"
	"MSaaS-Framework/MSaaS/pkg/object"
)

// CreateProject creates a new Project
func CreateProject(c *gin.Context) {

	// project 객체 생성
	project := object.Project{}

	// get project data from request
	c.BindJSON(&project)

	// create project in host path
	core.CreateProjectInHostPath(&project)

	// Save project to database by using Ent
	// Context에서 DBClient를 가져옴
	client, err := base.GetDBClientFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get database client"})
		return
	}

	tx, err := client.Tx(c) // 트랜잭션 시작
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to start transaction"})
		return
	}

	// // GeneralSpec을 생성하기 위해 ent.GeneralSpec 객체 생성
	// generalSpec, err := tx.GeneralSpec.Create().
	// 	SetName(project.General.Name).
	// 	SetType(project.General.Type).
	// 	SetDescription(project.General.Description).
	// 	Save(c)
	// if err != nil {
	// 	tx.Rollback() // 실패 시 롤백
	// 	c.JSON(500, gin.H{"error": "Failed to create GeneralSpec"})
	// 	return
	// }

	// // Project를 생성하기 위해 ent.Project 객체 생성
	// _, err = tx.Project.Create().
	// 	SetGeneralspec(generalSpec).
	// 	Save(c)
	// if err != nil {
	// 	tx.Rollback() // 실패 시 롤백
	// 	c.JSON(500, gin.H{"error": "Failed to create Project"})
	// 	return
	// }

	if err := tx.Commit(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.String(http.StatusOK, "Create Project")
}

// GetProject gets a Project by ID
func GetProject(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add get logic
	c.String(http.StatusOK, "Get Project with ID: %s", id)
}

// UpdateProject updates an existing Project by ID
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add update logic
	c.String(http.StatusOK, "Update Project with ID: %s", id)
}

// DeleteProject deletes an existing Project by ID
func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	// // change id type as string to uuid
	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to parse ID"})
	// 	return
	// }

	// client, err := base.GetDBClientFromContext(c)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to get database client"})
	// 	return
	// }

	// // get project object from database
	// project, err := client.Project.Query().
	// 	Where(project.ID(uuid)).
	// 	WithGeneralspec(). // GeneralSpec을 함께 로드
	// 	Only(c)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to get Project with GeneralSpec"})
	// 	return
	// }

	// // delete host path first
	// err = core.DeleteProjectInHostPath(project.Edges.Generalspec.Name)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to delete Project in host path"})
	// 	return
	// }

	// // 트랜잭션 시작
	// tx, err := client.Tx(c)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to start transaction"})
	// 	return
	// }

	// // Delete project in database
	// err = tx.Project.DeleteOneID(uuid).Exec(c)
	// if err != nil {
	// 	tx.Rollback()
	// 	c.JSON(500, gin.H{"error": "Failed to delete Project"})
	// 	return
	// }

	// // Delete associated GeneralSpec
	// err = tx.GeneralSpec.DeleteOneID(project.Edges.Generalspec.ID).Exec(c)
	// if err != nil {
	// 	tx.Rollback()
	// 	c.JSON(500, gin.H{"error": "Failed to delete GeneralSpec"})
	// 	return
	// }

	// // 트랜잭션 커밋
	// if err := tx.Commit(); err != nil {
	// 	c.JSON(500, gin.H{"error": "Failed to commit transaction"})
	// 	return
	// }

	c.String(http.StatusOK, "Delete Project with ID: %s", id)
}

// RegisterProjectRoutes registers the CRUD routes for Project
func RegisterProjectRoutes(router *gin.Engine) {
	router.POST("/project", CreateProject)
	router.GET("/project/:id", GetProject)
	router.PUT("/project/:id", UpdateProject)
	router.DELETE("/project/:id", DeleteProject)
}
