package main

import (
	"fmt"
	"net/http"
	"net/url"

	"randomchats/pkg/chat"
	slack "randomchats/pkg/client"
)

func main() {
	httpClient := http.DefaultClient
	// get from env var
	oauthToken := ""
	channelID := ""
	personPerGroup := 3
	slackUrl, _ := url.Parse("https://slack.com/api")
	slackClient := slack.NewClient(httpClient, *slackUrl, oauthToken)
	service := chat.NewService(slackClient)
	if err := service.Process(channelID, personPerGroup); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Message sent!")
}
