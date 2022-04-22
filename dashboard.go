package hgApi

import (
	"github.com/dghubble/sling"
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
}

type Dashboard struct {
	Meta      DashboardMeta       `json:"meta"`
	Dashboard DashboardDefinition `json:"dashboard"`
}

type DashboardDefinition struct {
	Annotations   interface{}   `json:"annotations"`
	Editable      bool          `json:"editable"`
	GnetId        interface{}   `json:"gnetId"`
	GraphTooltip  int           `json:"graphTooltip"`
	Id            int           `json:"id"`
	Iteration     int64         `json:"iteration"`
	Links         []interface{} `json:"links"`
	Panels        []interface{} `json:"panels"`
	SchemaVersion int           `json:"schemaVersion"`
	Style         string        `json:"style"`
	Tags          []interface{} `json:"tags"`
	Templating    interface{}   `json:"templating"`
	Timezone      interface{}   `json:"timezone"`
	Title         string        `json:"title"`
	Version       int           `json:"version"`
}

type DashboardMeta struct {
	Slug string `json:"slug"`
}

type DashboardSaveResponse struct {
	ID      int64  `json:"id"`
	Slug    string `json:"slug"`
	Status  string `json:"status"`
	UID     string `json:"uid"`
	Version int64  `json:"version"`
}

type DashboardService struct {
	sling *sling.Sling
}

func newDashboardService(sling *sling.Sling) *DashboardService {
	return &DashboardService{sling: sling.Path("v2/grafana/dashboards/")}
}

func (d *DashboardService) Get(slug string) (*Dashboard, *http.Response, error) {
	dashboard := new(Dashboard)
	apiError := new(APIError)

	resp, err := d.sling.New().Get(slug).Receive(dashboard, apiError)

	return dashboard, resp, handleError(err, *apiError)
}

func (d *DashboardService) Create(dashboard *Dashboard) (*DashboardSaveResponse, *http.Response, error) {
	dashboardResponse := new(DashboardSaveResponse)
	apiError := new(APIError)

	resp, err := d.sling.New().Post("").BodyJSON(dashboard.Dashboard).Receive(dashboardResponse, apiError)

	return dashboardResponse, resp, handleError(err, *apiError)
}

func (d *DashboardService) Update(dashboard *Dashboard) (*DashboardSaveResponse, *http.Response, error) {
	dashboardResponse := new(DashboardSaveResponse)
	apiError := new(APIError)

	resp, err := d.sling.New().Put("").BodyJSON(dashboard.Dashboard).Receive(dashboardResponse, apiError)

	return dashboardResponse, resp, handleError(err, *apiError)
}

func (d *DashboardService) Delete(slug string) (*http.Response, error) {
	apiError := new(APIError)

	resp, err := d.sling.New().Delete(slug).Receive(nil, apiError)

	return resp, handleError(err, *apiError)
}

func (d *DashboardService) Exist(slug string) bool {
	_, resp, _ := d.Get(slug)

	return resp.StatusCode != 404
}
