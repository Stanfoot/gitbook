package gitbook

import (
	"fmt"
	"strings"
	"time"
)

func (b Book) temporaryPath(addPath string) string {
	return b.BasePath + addPath
}

func (b Book) repoPath() string {
	return b.temporaryPath(fmt.Sprintf("%v", time.Now().Unix()))
}

func (b Book) outputFile(repo string, extension string) string {
	return b.temporaryPath(trimRepoOwner(repo) + extension)
}

func trimRepoOwner(repo string) string {
	var repoName string
	slice := strings.Split(repo, "/")

	if len(slice) < 1 {
		repoName = slice[0]
	} else {
		repoName = slice[1]
	}

	return repoName
}
