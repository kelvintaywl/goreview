package service

import (
	"context"
	"testing"

	"github.com/kelvintaywl/goreview/domain"
	"github.com/stretchr/testify/assert"
)

func TestRandomReviewersSuccess(t *testing.T) {
	cfg := domain.ReviewConfig{
		NumReviewers: 3,
		Reviewers:    []string{"richard.hendricks", "erlich.bachman", "gilfoyle", "dinesh"},
		WebhookURL:   "",
	}
	exclude := "richard.hendricks"
	reviewers, err := RandomReviewers(context.Background(), cfg, exclude)
	assert.Nil(t, err)
	assert.Len(t, reviewers, 3)
}

func TestRandomReviewersFailInvalidNum(t *testing.T) {
	cfg := domain.ReviewConfig{
		NumReviewers: 0,
		Reviewers:    []string{"a", "b", "c"},
		WebhookURL:   "",
	}
	_, err := RandomReviewers(context.Background(), cfg, "")
	assert.NotNil(t, err)
}

func TestRandomReviewersSuccessEvenIfNotEnoughReviewers(t *testing.T) {
	cfg := domain.ReviewConfig{
		NumReviewers: 5,
		Reviewers:    []string{"a", "b", "c"},
		WebhookURL:   "",
	}
	reviewers, err := RandomReviewers(context.Background(), cfg, "a")
	assert.Nil(t, err)
	assert.True(t, len(reviewers) <= 5)
}
