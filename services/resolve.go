package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var DappCategoryMap map[string]string

func LoadCategoryMap() error {
	path := filepath.Join("internal", "data", "category.json")

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open category.json: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&DappCategoryMap)
	if err != nil {
		return fmt.Errorf("failed to decode category.json: %v", err)
	}

	return nil
}
