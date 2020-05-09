package chat

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
	// randomise user name slice

	// split user name slice based on ppg (persons per group)

	// compose a message
	message := ""
	if err := s.slack.PostMessage(channelID, message); err != nil {
		return err
	}
	return nil
}
