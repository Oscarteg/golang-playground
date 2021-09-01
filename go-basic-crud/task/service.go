package task

type Service interface {
	Store(input InputTask) (Task, error)
	Index() ([]Task, error)
	Delete(input DeleteTask) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Store(input InputTask) (Task, error) {
	task := Task{
		Name:        input.Name,
		Description: input.Description,
	}

	newTask, error := s.repository.Insert(task)

	if error != nil {
		return task, error
	}

	return newTask, nil
}

func (s *service) Index() ([]Task, error) {
	return s.repository.Index()
}

func (s *service) Delete(input DeleteTask) error {
	return s.repository.Delete(input.Id)
}