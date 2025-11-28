package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/m4nsur/weather-lrn/geo"
)

func GetWeather(geo geo.GeoData, format int) (string, error){
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
	return "", err
	}

	body, err := io.ReadAll(resp.Body) 
	defer resp.Body.Close()
	if err != nil {
	return "", err
	}

	return string(body), nil
}