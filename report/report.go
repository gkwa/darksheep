package report

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/taylormonacelli/darksheep/fetch"
)

var repositories []fetch.RepositoryInfo

func LoadCache() error {
	file, err := os.Open(fetch.Cache)
	if err != nil {
		return fmt.Errorf("error opening data.json: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading data.json: %v", err)
	}

	err = json.Unmarshal(content, &repositories)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %v", err)
	}
	return nil
}

func Report1() error {
	err := fetch.Run()
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	err = LoadCache()
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	for _, repo := range repositories {
		fmt.Println(repo.Subpath, repo.GitURL, repo.GitURL)
	}

	return nil
}
