package catalog

type Cake struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Price       float64  `json:"price"`
	Weight      float64  `json:"weight"` // Weight in grams
	Images      []string `json:"images"` // Slice of image URLs
}

var Cakes = []Cake{
	{
		ID:          "1",
		Title:       "Chocolate Forest",
		Category:    "Birthday",
		Description: "Filled with chocolate",
		Weight:      500,
		Price:       530.0,
		Images:      []string{"http://cdn.images.cake1.png"},
	},
	{
		ID:          "2",
		Title:       "Chocolate Vanilla",
		Category:    "Birthday",
		Description: "Filled with vanilla chocolate",
		Weight:      500,
		Price:       420.0,
		Images:      []string{"http://cdn.images.cake2.png"},
	},
	{
		ID:          "3",
		Title:       "Red Velvet",
		Category:    "Anniversary",
		Description: "Filled with red velvet",
		Weight:      500,
		Price:       600.0,
		Images:      []string{"http://cdn.images.cake3.png"},
	},
}
