package product

type CategoryFormatter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func FormatCategory(category Category) CategoryFormatter {
	categoryFormatter := CategoryFormatter{}
	categoryFormatter.Id = category.Id
	categoryFormatter.Name = category.Name

	return categoryFormatter
}

func FormatCategories(categories []Category) []CategoryFormatter {
	categoriesFormatter := []CategoryFormatter{}

	for _, category := range categories {
		categoryFormatter := FormatCategory((category))
		categoriesFormatter = append(categoriesFormatter, categoryFormatter)
	}

	return categoriesFormatter
}
