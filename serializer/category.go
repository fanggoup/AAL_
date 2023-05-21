package serializer

import "AAL_time/modle"

type Category struct {
	ID uint `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func BuildCategory(item modle.Category) Category {
	return Category{
		ID: item.ID,
		Name: item.Name,
		Description: item.Description,
	}
}

func BuildCategorys(items []modle.Category) (categorys []Category){
	for _,item := range items{
		category := BuildCategory(item)
		categorys = append(categorys,category)
	}
	return categorys
}
