package service

import (
	"context"
	"encoding/base64"
	"os"

	"github.com/google/go-github/github"
	"github.com/kelvintaywl/goreview/domain"
	er "github.com/kelvintaywl/goreview/domain/errors"
	"golang.org/x/oauth2"
)

type (
	// GithubClient is an interface that allows us to interact with some of the Github API.
	GithubClient interface {
		GetContents(ctx context.Context, owner, repo, filepath, branch string) (string, error)
		AssignReviewers(ctx context.Context, owner, repo string, num int, reviewers []string) error
	}

	proxy struct {
		client *github.Client
	}
)

func (c proxy) GetContents(ctx context.Context, owner, repo, filepath, branch string) (string, error) {
	opts := &github.RepositoryContentGetOptions{
		Ref: branch,
	}
	fc, _, _, err := c.client.Repositories.GetContents(
		ctx,
		owner,
		repo,
		filepath,
		opts,
	)
	if err != nil {
		return "", err
	}
	if fc.Content == nil {
		return "", er.NewContentNotFound(filepath)
	}

	bContent, err := base64.StdEncoding.DecodeString(*fc.Content)
	if err != nil {
		return "", er.NewContentNotFound(filepath)
	}

	return string(bContent), nil
}

// AssignReviewers assigns reviewers to the Github pull request.
func (c proxy) AssignReviewers(ctx context.Context, owner, repo string, prNum int, reviewers []string) error {
	req := github.ReviewersRequest{
		Reviewers: reviewers,
	}
	_, _, err := c.client.PullRequests.RequestReviewers(
		ctx,
		owner,
		repo,
		prNum,
		req,
	)
	if err != nil {
		return er.NewGitHubError(err)
	}
	return nil
}

// NewGithubClient returns a GithubClient.
func NewGithubClient() GithubClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv(domain.GithubTokenEnv)},
	)
	tc := oauth2.NewClient(ctx, ts)
	c := github.NewClient(tc)
	return proxy{
		client: c,
	}
}
