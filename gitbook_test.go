package gitbook

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	repo        = "Stanfoot/owned_service_spec.doc"
	userName    = "sawadashota"
	extension   = ".zip"
	basePath    = "/Users/sawadashota/Downloads"
	githubToken = "tb4xruty34rv34tuxb346rytv3fbxty2"
)

func TestNew(t *testing.T) {
	book := New(repo, userName, extension, basePath)

	assert.IsType(t, &Book{}, book)
	assert.Equal(t, book.BasePath, basePath)
	assert.Equal(t, book.UserName, userName)
	assert.Equal(t, book.Repo, repo)
}

func TestGithub(t *testing.T) {
	book := New(repo, userName, extension, basePath)
	assert.Empty(t, book.GithubToken)
	book.Github(githubToken)
	assert.Equal(t, book.GithubToken, githubToken)
}
