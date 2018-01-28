package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/kelvintaywl/goreview/domain"
	er "github.com/kelvintaywl/goreview/domain/errors"
)

// FetchConfig ...
func FetchConfig(ctx context.Context, data domain.PullRequestEventPayload) (*domain.ReviewConfig, error) {
	// TODO: use redis to store fetched config so we dont need to always fetch it
	gc := NewGithubClient(ctx)
	cfgJSON, err := gc.GetContents(
		ctx,
		data.Repository.Owner.Name,
		data.Repository.Name,
		"goreview.json",
		data.PullRequest.Head.Branch,
	)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return nil, er.NewGitHubError(err)
	}
	fmt.Printf("cfgJSON: %s", cfgJSON)

	var cfg domain.ReviewConfig
	if err = json.Unmarshal([]byte(cfgJSON), &cfg); err != nil {
		fmt.Printf("JSON parse err: %s", err.Error())
		return nil, er.ErrJSONParseFailed
	}
	return &cfg, nil
}

// RandomReviewers ...
func RandomReviewers(ctx context.Context, cfg domain.ReviewConfig, exclude string) ([]string, error) {
	potential := cfg.Reviewers
	candidates := make([]string, 0, len(potential))
	for _, candidate := range potential {
		if candidate != exclude {
			candidates = append(candidates, candidate)
		}
	}
	size := cfg.NumReviewers
	if size < 1 {
		return nil, er.NewInvalidArgument(
			fmt.Sprintf("invalid num of reviewers requested: %d", size),
		)
	}

	if len(candidates) <= size {
		return candidates, nil
	}

	suggested := make([]string, size)
	perm := rand.Perm(len(candidates))
	for i := 0; i <= size; i++ {
		pos := perm[i]
		suggested = append(suggested, candidates[pos])
	}
	return suggested, nil
}

// PostWebhook ...
func PostWebhook(ctx context.Context, data domain.PullRequestEventPayload, url string, reviewers []string) error {
	p := domain.AssignReviewersResponsePayload{
		URL:       data.PullRequest.URL,
		Number:    data.PullRequest.Number,
		Repo:      data.Repository.FullName,
		Reviewers: reviewers,
	}
	bmsg, err := json.Marshal(p)
	if err != nil {
		return errors.New("unable to parse response into JSON")
	}
	cli := &http.Client{Timeout: time.Second * 10}
	resp, err := cli.Post(url, "application/json", bytes.NewBuffer(bmsg))
	if err != nil || 400 <= resp.StatusCode {
		if err != nil {
			fmt.Printf("Error posting HTTP: %v", err)
		}
		return fmt.Errorf("failed to send payload to webhook url: %s", url)
	}
	return nil
}

// AssignReviewers ...
func AssignReviewers(ctx context.Context, data domain.PullRequestEventPayload, cfg domain.ReviewConfig) error {
	gc := NewGithubClient(ctx)
	reviewers, err := RandomReviewers(ctx, cfg, data.User.Name)
	if err != nil {
		return err
	}

	err = gc.AssignReviewers(
		ctx,
		data.Repository.Owner.Name,
		data.Repository.Name,
		int(data.PullRequest.Number),
		reviewers,
	)
	if err != nil {
		return err
	}
	if len(cfg.WebhookURL) > 0 {
		fmt.Printf("Sending Webhooks: %s, for reviewers: %v", cfg.WebhookURL, reviewers)
		return PostWebhook(ctx, data, cfg.WebhookURL, reviewers)
	}
	return nil
}
