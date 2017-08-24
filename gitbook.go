package gitbook

import (
	"github.com/mholt/archiver"
	"github.com/revel/revel"
	"os"
	"os/exec"
)

func New(repo string, user string, extension string) Book {
	return Book{
		UserName:     user,
		Repo:         repo,
		RepoPath:     repoPath(),
		OutputFolder: outputFile(repo, ""),
		OutputFile:   outputFile(repo, extension),
	}
}

// Clone GitBook source from Github
func (b Book) CloneRepo() error {
	token := revel.Config.StringDefault("github.token", "")

	privateGithubAccessUrl := "https://" + token + "@github.com/" + b.Repo + ".git"

	_, err := exec.Command("git", "clone", privateGithubAccessUrl, b.RepoPath).Output()

	if err != nil {
		return err
	}

	return nil
}

// Remove repository directory and pdf/zip
func (b Book) RmAll() error {
	if err := os.RemoveAll(b.RepoPath); err != nil {
		return err
	}

	if err := os.Remove(b.OutputFile); err != nil {
		return err
	}

	if _, err := os.Stat(b.OutputFolder); err == nil {
		os.RemoveAll(b.OutputFolder)
	}

	return nil
}

// Exec gitbook command
func (b Book) GeneratePdf() error {
	_, err := exec.Command("gitbook", "pdf", b.RepoPath, b.OutputFile).Output()
	if err != nil {
		return err
	}

	return nil
}

// Exec gitbook command
func (b Book) Build(format string) error {
	_, err := exec.Command("gitbook", "build", "--format "+format, b.RepoPath, b.OutputFolder).Output()
	if err != nil {
		return err
	}

	return nil
}

// Zip target path
func (b Book) Zip() error {
	return archiver.Zip.Make(b.OutputFile, []string{b.OutputFolder})
}
