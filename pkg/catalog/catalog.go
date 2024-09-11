package catalog

type Cake struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"` // Weight in gms
	Image       string  `json:"image"`  // URL of the image
}

var Cakes = []Cake{
	{
		ID:          "1",
		Title:       "Chocolate Forest",
		Category:    "Birthday",
		Description: "Filled with chocolate",
		Weight:      500,
		Price:       530.0,
		Image:       "http://cdn.images.cake1.png",
	},
	{
		ID:          "2",
		Title:       "Chocolate Vanila",
		Category:    "Birthday",
		Description: "Filled with vanila chocolate",
		Weight:      500,
		Price:       420.0,
		Image:       "http://cdn.images.cake2.png",
	},
	{
		ID:          "3",
		Title:       "Red Velvet",
		Category:    "Annivesary",
		Description: "Filled with red velvet",
		Weight:      500,
		Price:       600.0,
		Image:       "http://cdn.images.cake3.png",
	},
}
