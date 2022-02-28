package hostedgraphite

import (
	"github.com/dghubble/sling"
	"net/http"
)

type Dashboard struct {
	Dashboard DashboardDefinition `json:"dashboard"`
}

type DashboardDefinition struct {
	Style         string        `json:"style"`
	Templating    interface{}   `json:"templating"`
	Links         []interface{} `json:"links"`
	GraphTooltip  int           `json:"graphTooltip"`
	AlertPanelMap interface{}   `json:"alertPanelMap"`
	Editable      bool          `json:"editable"`
	Annotations   interface{}   `json:"annotations"`
	GnetId        interface{}   `json:"gnetId"`
	Timepicker    interface{}   `json:"timepicker"`
	Title         string        `json:"title"`
	Version       int           `json:"version"`
	Time          interface{}   `json:"time"`
	Timezone      string        `json:"timezone"`
	SchemaVersion int           `json:"schemaVersion"`
	Panels        []interface{} `json:"panels"`
}

type DashboardService struct {
	sling *sling.Sling
}

func newDashboardService(sling *sling.Sling) *DashboardService {
	return &DashboardService{sling: sling.Path("v2/grafana/dashboards")}
}

func (d *DashboardService) Get(slug string) (*Dashboard, *http.Response, error) {
	dashboard := new(Dashboard)
	apiError := new(APIError)

	resp, err := d.sling.New().Get(slug).Receive(dashboard, apiError)

	return dashboard, resp, handleError(err, *apiError)
}

func (d *DashboardService) Create(dashboard *Dashboard) (*http.Response, error) {
	apiError := new(APIError)

	resp, err := d.sling.New().Post("").BodyJSON(dashboard).Receive(nil, apiError)

	return resp, handleError(err, *apiError)
}

func (d *DashboardService) Update(dashboard *Dashboard) (*http.Response, error) {
	apiError := new(APIError)

	resp, err := d.sling.New().Put("").BodyJSON(dashboard).Receive(nil, apiError)

	return resp, handleError(err, *apiError)
}

func (d *DashboardService) Delete(dashboard *Dashboard) (*http.Response, error) {
	apiError := new(APIError)

	resp, err := d.sling.New().Post("").BodyJSON(dashboard).Receive(nil, apiError)

	return resp, handleError(err, *apiError)
}

func (d *DashboardService) Exist(slug string) bool {
	_, resp, _ := d.Get(slug)

	return resp.StatusCode != 404
}
