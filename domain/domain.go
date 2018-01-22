package domain

const (
	GitHubTokenEnv string = "GITHUB_ACCESS_TOKEN"
	GitHubURL      string = "https://github.com/kelvintaywl/goreview"
)

type (
	UserPayload struct {
		Name string `json:"login"`
	}

	RepoPayload struct {
		Name     string      `json:"name"`
		FullName string      `json:"full_name"`
		Owner    UserPayload `json:"owner"`
	}

	HeadPayload struct {
		Branch string `json:"ref"`
	}

	PullRequestPayload struct {
		ID     int64       `json:"id"`
		URL    string      `json:"url"`
		Number int64       `json:"number"`
		State  string      `json:"state"`
		Head   HeadPayload `json:"head"`
	}

	PullRequestEventPayload struct {
		Action      string             `json:"action"`
		Repository  RepoPayload        `json:"repository"`
		User        UserPayload        `json:"sender"`
		PullRequest PullRequestPayload `json:"pull_request"`
	}

	ReviewConfig struct {
		NumReviewers int      `json:"num_reviewers" default:"2"`
		Reviewers    []string `json:"reviewers"`
		WebhookURL   string   `json:"webhook_url,omitempty"`
	}
)
