package random_chats

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"letschat/pkg/chat"
	slack "letschat/pkg/client"
)

func LetsChat(w http.ResponseWriter, r *http.Request) {
	httpClient := http.DefaultClient
	// get from env var
	oauthToken := os.Getenv("OAUTH_KEY")
	channelID := os.Getenv("CHANNEL_ID")
	personPerGroup, err := strconv.Atoi(os.Getenv("PER_GROUP"))
	if err != nil {
		fmt.Printf("Invalid person per group")
		return
	}
	slackUrl, _ := url.Parse("https://slack.com/api")
	slackClient := slack.NewClient(httpClient, *slackUrl, oauthToken)
	service := chat.NewService(slackClient)
	if err := service.Process(channelID, personPerGroup); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Message sent!")
}
