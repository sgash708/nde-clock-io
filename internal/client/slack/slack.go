package slack

import (
	"os"

	"github.com/slack-go/slack"
)

type SlackInterface interface {
	UploadFile(string) error
}

type Slack struct {
	Client  *slack.Client
	Channel string
}

func NewSlack(token, channel string) *Slack {
	return &Slack{
		Client:  slack.New(token),
		Channel: channel,
	}
}

func (s Slack) UploadFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	_, err = s.Client.UploadFile(
		slack.FileUploadParameters{
			Reader:   f,
			Filename: "inout_screenshot",
			Channels: []string{s.Channel},
		},
	)

	return err
}
