package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-basic-crud/task"
	"gorm.io/gorm"
	"net/http"
)

type Response struct {
	Data  interface{} `json:data`
	Error string      `json:omitempty`
}

type taskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *taskHandler {
	return &taskHandler{taskService}
}

func (handler *taskHandler) Store(c *gin.Context) {
	var input task.InputTask
	err := c.ShouldBindJSON(&input)

	if err != nil {

		response := Response{
			Error: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)

		return
	}

	newTask, err := handler.taskService.Store(input)

	if err != nil {
		response := Response{
			Error: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)

		return
	}

	response := Response{
		Data: newTask,
	}

	c.JSON(http.StatusOK, response)
}

func (handler *taskHandler) Index(c *gin.Context) {
	var input task.InputTask
	err := c.ShouldBindJSON(&input)
	tasks, err := handler.taskService.Index()

	if err != nil {
		response := Response{
			Error: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)

	}

	response := Response{
		Data: tasks,
	}

	c.JSON(http.StatusOK, response)
}


func (handler *taskHandler) Delete(c *gin.Context) {
	input := task.DeleteTask{Id: c.Param("id")}

	err := handler.taskService.Delete(input)

	if err != nil {
		response := Response{
			Error:    err.Error(),
		}

		c.JSON(http.StatusConflict, response)

		if errors.Is(err, gorm.ErrRecordNotFound) {

			c.AbortWithStatus(http.StatusNotFound)
		}
	}


	c.JSON(http.StatusNoContent, struct {}{})
}
