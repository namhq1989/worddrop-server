package scraper

type CategoryType string

const (
	CategoryPolitics      CategoryType = "politics"
	CategoryTechnology    CategoryType = "technology"
	CategoryBusiness      CategoryType = "business"
	CategoryScience       CategoryType = "science"
	CategoryHealth        CategoryType = "health"
	CategorySports        CategoryType = "sports"
	CategoryEntertainment CategoryType = "entertainment"
	CategoryWorldNews     CategoryType = "world"
	CategoryEducation     CategoryType = "education"
)

func AllCategories() []CategoryType {
	return []CategoryType{
		CategoryPolitics,
		CategoryTechnology,
		CategoryBusiness,
		CategoryScience,
		CategoryHealth,
		CategorySports,
		CategoryEntertainment,
		CategoryWorldNews,
		CategoryEducation,
	}
}
