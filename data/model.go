package data

type RepositoryInfo struct {
	BrowseURL string `json:"browse_url"`
	CreatedAt string `json:"created_at"`
	GitCommit string `json:"git_commit"`
	GitURL    string `json:"git_url"`
	IndexedAt string `json:"indexed_at"`
	Path      string `json:"path"`
	Release   string `json:"release"`
	Subpath   string `json:"subpath"`
	Version   string `json:"version"`
}
