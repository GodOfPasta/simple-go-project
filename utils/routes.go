package utils

import (
	"net/http"
	"simple-go-project/db"
	"simple-go-project/models"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	var tasks []models.Task
	var dbase, err = db.OpenDBReadConnection()

	rows, err := dbase.Query("SELECT * FROM todo_list.tasks;")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task

		if err := rows.Scan(&task.HashKey, &task.Name, &task.Description, &task.Created, &task.Updated, &task.Deadline, &task.Closed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
	db.CloseDBConnection(dbase)
}
