package product

type Service interface {
	NewCategory(input CreateCategoryInput) (Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) NewCategory(input CreateCategoryInput) (Category, error) {
	category := Category{}
	category.Name = input.Name

	newCategory, err := s.repository.CreateCategory(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}
