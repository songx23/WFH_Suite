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
	slackClient := slack.NewClient(httpClient, url.URL{Host: "https://slack.com/api"}, oauthToken)
	service := chat.NewService(slackClient)
	if err := service.Process(channelID, personPerGroup); err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Message sent!")
}
