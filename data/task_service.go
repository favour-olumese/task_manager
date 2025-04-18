package data

import (
	"task_manager/models"
	"time"
)

// In memory data
var Tasks = []models.Task{
	{ID: "1", Title: "Bill Board", Description: "To design a new bill board.", DueDate: time.Now().AddDate(0, 0, 3), Status: "In Progress"},
	{ID: "2", Title: "Bill Board", Description: "To design a new bill board.", DueDate: time.Now().AddDate(0, 0, 4), Status: "In Progress"},
	{ID: "3", Title: "Bill Board", Description: "To design a new bill board.", DueDate: time.Now().AddDate(0, 0, 1), Status: "Pending"},
	{ID: "4", Title: "New Flyer", Description: "To design a new flyer.", DueDate: time.Now().AddDate(0, 0, 7), Status: "Completed"},
}
