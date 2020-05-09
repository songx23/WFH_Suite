package slack

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"randomchats/pkg/httpclient"
)

type Client struct {
	httpClient *http.Client
	baseURL    url.URL
	oauthToken string
}

func NewClient(httpClient *http.Client, baseURL url.URL, token string) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		oauthToken: token,
	}
}

func (c *Client) GetUserIDsInChannel(channelID string) ([]string, error) {
	reqURL := c.baseURL
	queryString := fmt.Sprintf("?token=%s&channel=%s", url.PathEscape(c.oauthToken), url.PathEscape(channelID))
	reqURL.Path = path.Join(reqURL.Path, "/conversations.members", queryString)
	req, _ := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var memberRes channelMembersResponse
	if err := httpclient.DoRequest(c.httpClient, req, &memberRes); err != nil {
		return []string{}, err
	}
	return memberRes.Members, nil
}

func (c *Client) GetUserName(userID string) (string, error) {
	reqURL := c.baseURL
	queryString := fmt.Sprintf("?token=%s&user=%s", url.PathEscape(c.oauthToken), url.PathEscape(userID))
	reqURL.Path = path.Join(reqURL.Path, "/users.info", queryString)
	req, _ := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var userRes userRespsone
	if err := httpclient.DoRequest(c.httpClient, req, &userRes); err != nil {
		return "", err
	}
	return userRes.User.name, nil
}

func (c *Client) PostMessage(channelID string, message string) error {
	reqURL := c.baseURL
	reqURL.Path = path.Join(reqURL.Path, "/chat.postMessage")
	form := url.Values{}
	form.Add("channel", channelID)
	form.Add("text", message)
	req, _ := http.NewRequest(http.MethodPost, reqURL.String(), strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.oauthToken))

	var msgRes postMessageResponse
	if err := httpclient.DoRequest(c.httpClient, req, &msgRes); err != nil {
		return err
	}

	if !msgRes.OK {
		return fmt.Errorf("Unknown error happened when posting message")
	}
	return nil
}
