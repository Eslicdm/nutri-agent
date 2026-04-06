package nutrition

import (
	"encoding/json"
	"os"
)

type Product struct {
	Name         string  `json:"name"`
	Portion      string  `json:"portion"`
	ProteinG     float64 `json:"protein_g"`
	CarbG        float64 `json:"carb_g"`
	FatG         float64 `json:"fat_g"`
	AveragePrice float64 `json:"average_price"`
}

func LoadCatalog(path string) ([]Product, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var catalog []Product
	err = json.Unmarshal(file, &catalog)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}
