package gitbook

import (
	"github.com/nlopes/slack"
)

// Post file to Slack
func (b Book) SendSlack() error {
	api := slack.New(b.Slack.Token)

	param := slack.FileUploadParameters{
		Title:          trimRepoOwner(b.Repo),
		File:           b.OutputFile,
		Channels:       []string{b.Slack.Channel},
		InitialComment: b.Slack.Message,
	}

	_, err := api.UploadFile(param)
	if err != nil {
		return err
	}

	return nil
}

// Slack Message
func (b *Book) Message(text string) *Book {
	b.Slack.Message = "@" + b.UserName + " " + text
	return b
}
