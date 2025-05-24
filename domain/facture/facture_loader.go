package facture

import (
	"encoding/json"
	"os"
)

func LoadFactureFromJSON(path string) (*Facture, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var f Facture
	if err := json.Unmarshal(file, &f); err != nil {
		return nil, err
	}

	return &f, nil
}
