package gitbook

import (
	"github.com/nlopes/slack"
	"github.com/revel/revel"
)

// Post file to Slack
func (b Book) SendSlack() {
	api := slack.New(b.Slack.Token)

	param := slack.FileUploadParameters{
		Title:          trimRepoOwner(b.Repo),
		File:           b.OutputFile,
		Channels:       []string{b.Slack.Channel},
		InitialComment: b.Slack.Message,
	}

	file, err := api.UploadFile(param)
	if err != nil {
		revel.ERROR.Println(err.Error())
	}

	revel.INFO.Println(file.URLPrivate)
}

// Slack Message
func (b *Book) Message(text string) *Book {
	b.Slack.Message = "@" + b.UserName + " " + text
	return b
}
