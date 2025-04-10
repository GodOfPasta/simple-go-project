package utils

import (
	"net/http"
	"simple-go-project/db"
	"simple-go-project/models"
	"simple-go-project/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetTasks(c *gin.Context) {
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
	db.CloseDBConnection(dbase)
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	var dbase, err = db.OpenDBReadConnection()
	var id = c.Param("id")
	var task models.Task

	err = dbase.QueryRow("SELECT * FROM todo_list.tasks WHERE hash_key = $1", id).
		Scan(&task.HashKey, &task.Name, &task.Description, &task.Created, &task.Updated, &task.Deadline, &task.Closed)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		} else {
			panic(err)
		}
	}

	db.CloseDBConnection(dbase)
	c.JSON(http.StatusOK, task)
}

func AddTask(c *gin.Context) {
	var task models.Task
	var dbase, err = db.OpenDBWriteConnection()
	var errors []string

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	} else {
		val := validator.New()
		val.RegisterValidation("deadlineValidator", validators.DeadlineValidator)

		err := val.Struct(task)
		if err != nil {
			if e, ok := err.(validator.ValidationErrors); ok {
				for _, fieldError := range e {
					// TODO: Write more clear error text maybe
					var errorText = "Field: " + fieldError.Field() + ", Error: " + fieldError.Tag()
					errors = append(errors, errorText)
				}
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors})
			return
		}
	}

	_, err = dbase.Exec("INSERT INTO todo_list.tasks (name, description, deadline, closed) VALUES ($1, $2, $3, $4)",
		task.Name, task.Description, task.Deadline, task.Closed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = dbase.QueryRow("SELECT * FROM todo_list.tasks ORDER BY created DESC LIMIT 1").
		Scan(&task.HashKey, &task.Name, &task.Description, &task.Created, &task.Updated, &task.Deadline, &task.Closed)

	if err != nil {
		panic(err)
	}

	db.CloseDBConnection(dbase)
	c.JSON(http.StatusOK, task)
}
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	var dbase, err = db.OpenDBWriteConnection()
	var errors []string

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	} else {
		val := validator.New()
		val.RegisterValidation("deadlineValidator", validators.DeadlineValidator)

		err := val.Struct(task)
		if err != nil {
			if e, ok := err.(validator.ValidationErrors); ok {
				for _, fieldError := range e {
					// TODO: Write more clear error text maybe
					var errorText = "Field: " + fieldError.Field() + ", Error: " + fieldError.Tag()
					errors = append(errors, errorText)
				}
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors})
			return
		}
	}

	_, err = dbase.Exec("UPDATE todo_list.tasks SET name=$1, description=$2, deadline=$3, closed=$4, updated=now() WHERE hash_key=$5",
		task.Name, task.Description, task.Deadline, task.Closed, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = dbase.QueryRow("SELECT * FROM todo_list.tasks ORDER BY created DESC LIMIT 1").
		Scan(&task.HashKey, &task.Name, &task.Description, &task.Created, &task.Updated, &task.Deadline, &task.Closed)

	if err != nil {
		panic(err)
	}

	db.CloseDBConnection(dbase)
	c.JSON(http.StatusOK, task)
}
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var dbase, err = db.OpenDBWriteConnection()
	_, err = dbase.Exec("DELETE FROM todo_list.tasks WHERE hash_key=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	db.CloseDBConnection(dbase)
	c.Status(http.StatusNoContent)
}
