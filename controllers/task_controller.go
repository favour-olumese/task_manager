package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// Get all tasks.
func GetAllTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Tasks)
}

// Get specific task based on ID.
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, task := range data.Tasks {
		if task.ID == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found."})
}

// Update existing task.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task

	// Bind the request data to the variable created.
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range data.Tasks {
		if task.ID == id {

			// Check for changes.
			// Title
			if updatedTask.Title != "" {
				data.Tasks[i].Title = updatedTask.Title
			}

			// Description
			if updatedTask.Description != "" {
				data.Tasks[i].Description = updatedTask.Description
			}

			// Status
			if updatedTask.Status != "" {
				data.Tasks[i].Status = updatedTask.Status
			}

			// Due Date
			if !updatedTask.DueDate.IsZero() {
				data.Tasks[i].DueDate = updatedTask.DueDate
			}

			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated."})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found."})
}

// Delete exiting task.
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range data.Tasks {
		if task.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Removed."})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found."})
}

// Create new task.
func NewTask(c *gin.Context) {
	var newTask models.Task

	// Bind request to variable
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the last ID in currently stored.
	memLastIndex := 0
	if len(data.Tasks) > 0 {
		memLastIndex = len(data.Tasks) - 1
	}

	lastID, _ := strconv.Atoi(data.Tasks[memLastIndex].ID)

	// Create a new ID and assign to the task.
	newTask.ID = strconv.Itoa(lastID + 1)

	data.Tasks = append(data.Tasks, newTask)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task Created."})
}
