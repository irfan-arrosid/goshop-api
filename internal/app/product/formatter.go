package product

type CategoryFormatter struct {
	Name string `json:"name"`
}

func FormatCategory(category Category) CategoryFormatter {
	categoryFormatter := CategoryFormatter{}
	categoryFormatter.Name = category.Name

	return categoryFormatter
}
