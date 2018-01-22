package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGithubClientReturnNonNil(t *testing.T) {
	cli := NewGithubClient(context.Background())
	assert.NotNil(t, cli)
}
