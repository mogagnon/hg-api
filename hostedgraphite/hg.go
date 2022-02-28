package hostedgraphite

import (
	"encoding/base64"
	"fmt"
	"github.com/dghubble/sling"
)

const hgAPI = "https://api.hostedgraphite.com/api/"

type Client struct {
	sling     *sling.Sling
	Dashboard *DashboardService
}

func NewClient(token string) *Client {
	b64Token := base64.StdEncoding.EncodeToString([]byte(token))

	base := sling.New().Base(hgAPI).Set("Authorization", fmt.Sprintf("Basic %s", b64Token))

	return &Client{
		sling:     base,
		Dashboard: newDashboardService(base.New()),
	}
}
