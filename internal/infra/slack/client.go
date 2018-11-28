package slack

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin/json"
)

var postMessageEndpoint = "https://slack.com/api/chat.postMessage"

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token}
}

func (c *Client) SendMessage(channel, text string) error {
	msg := MessageChannelRequest{Channel: channel, Text: text}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, postMessageEndpoint, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("request to Slack API failed with status %d", res.StatusCode)
	}
	return nil
}
