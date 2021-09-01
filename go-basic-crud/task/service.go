package task

type Service interface {
	Store(input InputTask) (Task, error)
	Find(input FindTask) (Task, error)
	Update(id string, input UpdateTask) (Task, error)
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

	return s.repository.Insert(task)
}

func (s *service) Index() ([]Task, error) {
	return s.repository.Index()
}

func (s *service) Delete(input DeleteTask) error {
	return s.repository.Delete(input.Id)
}

func (s *service) Find(input FindTask) (Task, error ){
	return s.repository.Find(input.Id)
}

func (s *service) Update(id string, input UpdateTask) (Task, error ){
	task := Task{
		Name:        input.Name,
		Description: input.Description,
	}

	return s.repository.Update(id, task)
}

