package slack

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/songx23/letschat/pkg/httpclient"
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
	queryString := fmt.Sprintf("token=%s&channel=%s", url.PathEscape(c.oauthToken), url.PathEscape(channelID))
	reqURL.Path = path.Join(reqURL.Path, "/conversations.members")
	reqURL.RawQuery = queryString
	request, _ := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var memberRes channelMembersResponse
	if err := httpclient.DoRequest(c.httpClient, request, &memberRes); err != nil {
		return []string{}, err
	}
	if !memberRes.OK {
		return []string{}, fmt.Errorf("slack error: %s", memberRes.Err)
	}
	return memberRes.Members, nil
}

func (c *Client) GetUserName(userID string) (string, error) {
	reqURL := c.baseURL
	queryString := fmt.Sprintf("token=%s&user=%s", url.PathEscape(c.oauthToken), url.PathEscape(userID))
	reqURL.Path = path.Join(reqURL.Path, "/users.info")
	reqURL.RawQuery = queryString
	req, _ := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var userRes userRespsone
	if err := httpclient.DoRequest(c.httpClient, req, &userRes); err != nil {
		return "", err
	}
	if !userRes.OK {
		return "", fmt.Errorf("slack error: %s", userRes.Err)
	}
	if userRes.User.IsBot {
		return "", nil
	}
	return userRes.User.Name, nil
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
		return fmt.Errorf("slack error: %s", msgRes.Err)
	}
	return nil
}
