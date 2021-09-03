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

	c.JSON(http.StatusCreated, response)
}

// ListTask godoc
// @Summary List tasks
// @Description get tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} task.Task
// @Router /tasks [get]
func (handler *taskHandler) Index(c *gin.Context) {
	var input task.InputTask
	err := c.ShouldBindJSON(&input)
	tasks, err := handler.taskService.Index()

	if err != nil {
		response := Response{
			Error: err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := Response{
		Data: tasks,
	}

	c.JSON(http.StatusOK, response)
}

func (handler *taskHandler) Find(c *gin.Context) {
	input := task.FindTask{Id: c.Param("id")}

	newTask, err := handler.taskService.Find(input)

	if err != nil {
		response := Response{
			Error:    err.Error(),
		}

		c.JSON(http.StatusNotFound, response)
		return

	}

	response := Response{
		Data: newTask,
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

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response)
			return
		}

		c.JSON(http.StatusConflict,  response)
		return
	}


	// No JSON here because we need an empty response
	c.Status(http.StatusNoContent)
}


func (handler *taskHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input task.UpdateTask
	err := c.ShouldBindJSON(&input)

	if err != nil {

		response := Response{
			Error: err.Error(),
		}

		c.JSON(http.StatusNotFound, response)

		return
	}

	newTask, err := handler.taskService.Update(id, input)

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


