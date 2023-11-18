package product

type Domain struct {
	resource ResourceItf
}

type Resource struct {
}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	Rating      string `json:"rating"`
	Price       string `json:"price"`
	StoreName   string `json:"store_name"`
}

type TokopediaSearchParams struct {
	Query     string
	Page      int
	SortOrder string
}
