# Hosted graphite API
Simple SDK to manage dashboard in Hosted graphite.

## Example
Basic usage for a dashboard
```go
var dashboard hgApi.Dashboard
client := hgApi.NewDefault("Your token")

file, _ := ioutil.ReadFile("dashboard.json")
_ = json.Unmarshal(file, &dashboard)

dashboardResp, resp, err := client.Dashboard.Create(&dashboard)
dashboardResp, resp, err = client.Dashboard.Update(&dashboard)
_, resp, err = client.Dashboard.Get("slug")
resp, err = client.Dashboard.Delete("slug")
```

## Test
You need a hosted graphite token in order to run the acceptance test.

```
HG_TOKE="your token" go test ./..
```

## Export dashboard
It's possible to export a dashboard directly from Hosted graphite. The procedure is the same as for [Grafana](https://grafana.com/docs/grafana/latest/dashboards/export-import/).