package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Repositories []RepositoryInfo

func LoadCache() error {
	file, err := os.Open(Cache)
	if err != nil {
		return fmt.Errorf("error opening %s: %v", Cache, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading %s: %v", Cache, err)
	}

	err = json.Unmarshal(content, &Repositories)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}
	return nil
}

func LoadData() error {
	err := Fetch()
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	err = LoadCache()
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	return nil
}
