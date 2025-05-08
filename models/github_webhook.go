package models

type GitHubWebhook struct {
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	HeadCommit struct {
		ID      string `json:"id"`
		Message string `json:"message"`
		Author  struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		} `json:"author"`
	} `json:"head_commit"`
}
