package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	Cache    = "daggerverse.json"
	CacheAge = 5 * time.Hour
)

func Fetch() error {
	url := "https://daggerverse.dev/api/refs"

	var repositories []RepositoryInfo
	if !isFileOlderThan(Cache, CacheAge) {
		slog.Debug("cache is not old. Skipping fetch.", "min age", CacheAge)
	} else {
		data, err := fetchData(url)
		if err != nil {
			return fmt.Errorf("error fetching data: %v", err)
		}

		err = savePrettyPrintedJSON(Cache, data)
		if err != nil {
			return fmt.Errorf("error writing to %s: %v", Cache, err)
		}

		err = json.Unmarshal(data, &repositories)
		if err != nil {
			return fmt.Errorf("error unmarshaling data: %v", err)
		}

		fmt.Println("Data saved successfully.")
	}

	for _, repo := range repositories {
		fmt.Println(repo.Subpath, repo.GitURL, repo.GitURL)
	}

	return nil
}

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func isFileOlderThan(filename string, age time.Duration) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return true // Treat as older if there is an error getting file information
	}

	fileAge := time.Since(fileInfo.ModTime())

	return fileAge > age
}

func savePrettyPrintedJSON(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(prettyJSON)
	if err != nil {
		return err
	}

	return nil
}
