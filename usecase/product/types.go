package ucproduct

type GetProductParam struct {
	Category string `json:"category"`
	Limit    int    `json:"limit"`
	Worker   int    `json:"worker"`
}
