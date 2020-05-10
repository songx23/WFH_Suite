package slack

type channelMembersResponse struct {
	OK      bool     `json:"ok"`
	Members []string `json:"members"`
	Err     string   `json:"error,omitempty"`
}

type userRespsone struct {
	OK   bool      `json:"ok"`
	User slackUser `json:"user"`
	Err  string    `json:"error,omitempty"`
}

type slackUser struct {
	Name  string `json:"name"`
	IsBot bool   `json:"is_bot,omitempty"`
}

type postMessageResponse struct {
	OK  bool   `json:"ok"`
	Err string `json:"error,omitempty"`
}

type attachment struct {
	Text string `json:"text"`
}

type messageAttachments struct {
	Attachments []attachment `json:"attachments"`
}

func mapAttachemnts(msgs []string) messageAttachments {
	var result []attachment
	for _, v := range msgs {
		result = append(result, attachment{Text: v})
	}
	return messageAttachments{Attachments: result}
}
