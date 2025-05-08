package entity

import "time"

type Commit struct {
	RepoName       string
	CommitID       string
	CommitMessage  string
	AuthorUsername string
	AuthorEmail    string
	Payload        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
