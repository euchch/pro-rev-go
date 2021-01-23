package lib

import "os"

type Product struct {
	Name       string   `json:"name"`
	Link       string   `json:"link"`
	Status     string   `json:"status"`
	Price      float64  `json:"price"`
	Tax        float64  `json:"tax"`
	Fees       float64  `json:"fees"`
	Commission float64  `json:"commission"`
	Total      []string `json:"total"`
}

func (p *Product) ListAll(options map[string]interface{}) []Product {
	var ret []Product
	path, err := os.Executable()
}
