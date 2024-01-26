package report

import (
	"fmt"
	"sort"

	"github.com/taylormonacelli/darksheep/data"
)

func Report1() error {
	err := data.LoadData()
	if err != nil {
		return fmt.Errorf("error running loadData func: %v", err)
	}

	copiedRepos := make([]data.RepositoryInfo, len(data.Repositories))
	copy(copiedRepos, data.Repositories)

	// sort by subpath
	sort.Slice(copiedRepos, func(i, j int) bool {
		return copiedRepos[i].Subpath < copiedRepos[j].Subpath
	})

	for _, repo := range copiedRepos {
		fmt.Println(repo.Subpath, repo.GitURL, repo.GitURL)
	}

	return nil
}
