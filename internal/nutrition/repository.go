package nutrition

import (
	"encoding/json"
	"os"
)

type Product struct {
	Name         string  `json:"name"`
	TotalWeightG float64 `json:"total_weight_g"`
	PortionG     float64 `json:"portion_g"`
	CarbG        float64 `json:"carb_g"`
	ProteinG     float64 `json:"protein_g"`
	FatG         float64 `json:"fat_g"`
	FiberG       float64 `json:"fiber_g"`
	Calories     float64 `json:"calories"`
	AveragePrice float64 `json:"average_price"`
	Category     string  `json:"category"`
	Strength     string  `json:"strength,omitempty"`
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
