package github_releases

import "time"

// RepositoryReleasesModel is the database model for repository releases
// objects.
type RepositoryReleasesModel struct {
	TotalCount int32
	Author     struct {
		Email string
		Login string
	}
	CreatedAt       time.Time
	Description     string
	DescriptionHtml string
	IsDraft         bool
	IsPrerelease    bool
	Name            string
	PublishedAt     time.Time
	Url             string
	TagName         string
	IsLatest        bool
}
