package slack

type channelMembersResponse struct {
	OK      bool     `json:"ok"`
	Members []string `json:"members"`
}

type userRespsone struct {
	OK   bool      `json:"ok"`
	User slackUser `json:"user"`
}

type slackUser struct {
	name string `json:"name"`
}

type postMessageResponse struct {
	OK bool `json:"ok"`
}
