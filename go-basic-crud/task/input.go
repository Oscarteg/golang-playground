package task

type InputTask struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTask struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type DeleteTask struct {
	Id string
}

type FindTask struct{
	Id string
}