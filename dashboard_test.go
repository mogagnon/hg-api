package hgApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestCreateDashboard(t *testing.T) {
	client := NewDefault(os.Getenv(HgToken))

	dashboard, err := dashboardTest("create.json")
	if err != nil {
		t.Fatal(err)
	}

	dashboardSaveResponse, response, err := client.Dashboard.Create(dashboard)
	if err != nil {
		t.Fatal(err)
	}

	validHttpStatus(t, 200, response)

	if _, err = client.Dashboard.Delete(dashboardSaveResponse.Slug); err != nil {
		t.Errorf("unable to cleanup dashboard %s", dashboardSaveResponse.Slug)
	}
}

func TestGetDashboard(t *testing.T) {
	client := NewDefault(os.Getenv(HgToken))

	dashboard, err := dashboardTest("create.json")
	if err != nil {
		t.Fatal(err)
	}

	dashboardSaveResponse, _, err := client.Dashboard.Create(dashboard)
	if err != nil {
		t.Fatal(err)
	}

	_, response, err := client.Dashboard.Get(dashboardSaveResponse.Slug)
	if err != nil {
		t.Fatal(err)
	}

	validHttpStatus(t, 200, response)

	if _, err = client.Dashboard.Delete(dashboardSaveResponse.Slug); err != nil {
		t.Errorf("unable to cleanup dashboard %s", dashboardSaveResponse.Slug)
	}
}

func TestUpdateDashboard(t *testing.T) {
	client := NewDefault(os.Getenv(HgToken))

	dashboard, err := dashboardTest("create.json")
	if err != nil {
		t.Fatal(err)
	}

	updatedDashboard, err := dashboardTest("update.json")
	if err != nil {
		t.Fatal(err)
	}

	dashboardSaveResponse, _, err := client.Dashboard.Create(dashboard)
	if err != nil {
		t.Fatal(err)
	}

	dashboardSaveResponse, response, err := client.Dashboard.Update(updatedDashboard)
	if err != nil {
		t.Fatal(err)
	}

	dashboardResponse, _, err := client.Dashboard.Get(dashboardSaveResponse.Slug)
	if err != nil {
		t.Fatal(err)
	}

	validHttpStatus(t, 200, response)

	if len(dashboardResponse.Dashboard.Panels) != 7 {
		t.Errorf("expected 7 panels got %d", len(dashboardResponse.Dashboard.Panels))
	}

	if _, err = client.Dashboard.Delete(dashboardSaveResponse.Slug); err != nil {
		t.Errorf("unable to cleanup dashboard %s", dashboardSaveResponse.Slug)
	}
}

func validHttpStatus(t *testing.T, expected int, response *http.Response) {
	if response.StatusCode != expected {
		t.Errorf("http status code expected %d got %d", expected, response.StatusCode)
	}
}

func dashboardTest(file string) (*Dashboard, error) {
	var dashboard DashboardDefinition
	byteJSON, err := ioutil.ReadFile(fmt.Sprintf("./test_data/%s", file))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteJSON, &dashboard)
	if err != nil {
		return nil, err
	}

	return &Dashboard{Dashboard: dashboard}, nil
}
