package controllers

import (
	"encoding/json"
	httpLib "net/http"

	"github.com/goravel/framework/contracts/http"
)

type MyipController struct {
	// Dependent services
}

func NewMyipController() *MyipController {
	return &MyipController{
		// Inject services
	}
}

func (r *MyipController) Index(ctx http.Context) http.Response {
	ip := ctx.Request().Ip()

	data, err := fetchGeoIP(ip)

	if err != nil {
		return ctx.Response().Json(500, http.Json{
			"error":   err.Error(),
			"message": "Internal Server Error",
		})
	}

	return ctx.Response().View().Make("myip.tmpl", map[string]any{
		"IP": ip,
		"Geo": map[string]any{
			"Country": data.Country,
			"Region":  data.Region,
			"City":    data.City,
			"ISP":     data.ISP,
			"AS":      data.AS,
		},
	})
}

type GeoIP struct {
	Country string `json:"country"`
	Region  string `json:"regionName"`
	City    string `json:"city"`
	ISP     string `json:"isp"`
	AS      string `json:"as"`
}

func fetchGeoIP(ip string) (*GeoIP, error) {
	resp, err := httpLib.Get("https://ip-api.com/json/" + ip)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data GeoIP
	json.NewDecoder(resp.Body).Decode(&data)

	return &data, nil
}
