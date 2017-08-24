package gitbook

import (
	"fmt"
	"github.com/revel/revel"
	"strings"
	"time"
)

func tmpPath(addPath string) string {
	return revel.Config.StringDefault("repo.path", "") + addPath
}

func repoPath() string {
	return tmpPath(fmt.Sprintf("%v", time.Now().Unix()))
}

func outputFile(repo string, extension string) string {
	return tmpPath(trimRepoOwner(repo) + extension)
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
