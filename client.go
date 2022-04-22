package hgApi

import (
	"encoding/base64"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/hashicorp/go-cleanhttp"
	"net/http"
)

const hgAPI = "https://api.hostedgraphite.com/api/"

type Client struct {
	baseUrl   string
	sling     *sling.Sling
	Dashboard *DashboardService
}

func NewDefault(token string) *Client {
	return New(token, "", nil)
}

func New(token string, baseUrl string, client *http.Client) *Client {
	b64Token := base64.StdEncoding.EncodeToString([]byte(token))

	if baseUrl == "" {
		baseUrl = hgAPI
	}

	if client == nil {
		client = cleanhttp.DefaultClient()
	}

	base := sling.New().Client(client).Base(baseUrl).Set("Authorization", fmt.Sprintf("Basic %s", b64Token))

	return &Client{
		sling:     base,
		Dashboard: newDashboardService(base.New()),
	}
}
