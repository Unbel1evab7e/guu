package tests

import (
	"github.com/Unbel1evab7e/guu"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecutePost(test *testing.T) {
	headers := make(map[string]string)

	headers["Content-Type"] = "application/json"

	rec, err := guu.ExecutePost[guu.TestPostResponse]("https://jsonplaceholder.typicode.com/posts",
		guu.TestPostRequest{Title: "foo", Body: "bar", UserID: 1}, nil, headers)

	assert.Equal(test, nil, err)
	assert.Condition(test, func() (success bool) {
		return rec != nil
	})
}

func TestExecuteGet(test *testing.T) {
	headers := make(map[string]string)

	headers["Content-Type"] = "application/json"

	queryParams := make(map[string]string)

	queryParams["postId"] = "1"

	rec, err := guu.ExecuteGet[[]guu.TestPostResponse]("https://jsonplaceholder.typicode.com/posts", queryParams, headers)

	assert.Equal(test, nil, err)
	assert.Condition(test, func() (success bool) {
		return rec != nil && len(*rec) > 0
	})
}
