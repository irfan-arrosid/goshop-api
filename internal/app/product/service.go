package product

type Service interface {
	NewCategory(input CreateCategoryInput) (Category, error)
	GetCategories() ([]Category, error)
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

func (s *service) GetCategories() ([]Category, error) {
	categories, err := s.repository.FindAllCategory()
	if err != nil {
		return categories, err
	}

	return categories, nil
}
