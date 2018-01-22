package errors

import (
	"errors"
	"fmt"
)

type (
	// InvalidArgument is thrown when caller issues invalid arguments.
	InvalidArgument struct {
		msg string
	}
	// ContentNotFound is thrown when requested resource's content is missing or empty
	ContentNotFound struct {
		resource string
	}

	// GitHubError is thrown when Github API request failed (wrapped).
	GitHubError struct {
		gherr error
	}
)

var (
	// ErrJSONParseFailed is thrown when requested content cannot be converted to valid JSON.
	ErrJSONParseFailed = errors.New("unable to parse to JSON")
)

// NewInvalidArgument returns a new InvalidArgument.
func NewInvalidArgument(msg string) *InvalidArgument {
	return &InvalidArgument{
		msg: msg,
	}
}

func (e *InvalidArgument) Error() string {
	return e.msg
}

// NewContentNotFound returns a new ContentNotFound.
func NewContentNotFound(resource string) *ContentNotFound {
	return &ContentNotFound{
		resource: resource,
	}
}

func (e *ContentNotFound) Error() string {
	return fmt.Sprintf(
		"empty or missing content; unable to get contents from resource: %s", e.resource,
	)
}

// NewGitHubError returns a new GiHubError.
func NewGitHubError(err error) *GitHubError {
	return &GitHubError{
		gherr: err,
	}
}

func (e *GitHubError) Error() string {
	return e.gherr.Error()
}
