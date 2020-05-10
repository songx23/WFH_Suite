package chat

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type slackClient interface {
	GetUserIDsInChannel(channelID string) ([]string, error)
	GetUserName(userID string) (string, error)
	PostMessage(channelID string, message string) error
}

type Service struct {
	slack slackClient
}

func NewService(client slackClient) *Service {
	return &Service{slack: client}
}

func (s *Service) Process(channelID string, ppg int) error {
	userIDs, err := s.slack.GetUserIDsInChannel(channelID)
	if err != nil {
		return err
	}
	var userNames []string
	for _, v := range userIDs {
		name, err := s.slack.GetUserName(v)
		if err != nil {
			return err
		}
		userNames = append(userNames, name)
	}
	// remove bot users
	userNames = deleteEmptyStrings(userNames)
	// randomise user name slice
	userNames = shuffle(userNames)
	// split user name slice based on ppg (persons per group)
	groups := splitGroups(userNames, ppg)
	// compose a message
	message := composeMessage(groups)
	// post message
	if err := s.slack.PostMessage(channelID, message); err != nil {
		return err
	}
	return nil
}

func deleteEmptyStrings(slice []string) []string {
	var r []string
	for _, str := range slice {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func shuffle(slice []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(slice))
	perm := r.Perm(len(slice))
	for i, randIndex := range perm {
		ret[i] = slice[randIndex]
	}
	return ret
}

func splitGroups(slice []string, ppg int) [][]string {
	if ppg < 0 || ppg > len(slice) {
		return [][]string{slice}
	}
	var result [][]string
	idx := 0
	for idx < len(slice) {
		rightBound := idx + ppg
		if rightBound > len(slice) {
			rightBound = len(slice)
		}
		result = append(result, slice[idx:rightBound])
		idx = rightBound
	}
	// adjust the last group
	if result != nil && len(result) > 1 {
		lastGroup := result[len(result)-1]
		if float32(len(lastGroup)) <= float32(ppg/2) {
			// append to the second last group
			secondLastGroup := &result[len(result)-2]
			*secondLastGroup = append(*secondLastGroup, lastGroup...)
			result = result[0 : len(result)-1]
		}
	}
	return result
}

func composeMessage(groups [][]string) string {
	msgs := []string{"Good day team:roller_coaster:. The random chat roster of this week :scroll::"}
	for _, v := range groups {
		for i := range v {
			v[i] = fmt.Sprintf("<@%s>", v[i])
		}
		msgLine := strings.Join(v, " :blob-wine-gif: ")
		msgs = append(msgs, msgLine)
	}
	return strings.Join(msgs, "\n")
}
