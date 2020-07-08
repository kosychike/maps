package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Point struct {
	Latitude float64
	Longitude float64
}

type Polygon struct {
	Points []Point
}


func (point *Point) toString() string {
	return fmt.Sprintf("%f,%f", point.Latitude, point.Longitude)
}

func (polygon *Polygon) toString() string {
	var result string
	var points []string
	for _, point := range polygon.Points {
		points = append(points, point.toString())
	}
	result = strings.Join(points, ",")
	return result
}

func GetMapWithPolygon (output string, geometry string) {
	endpoint, _ := url.Parse("https://image.maps.ls.hereapi.com/mia/1.6/region?apiKey=nG7eRnSlGEFHGSCrzEFFd0YfCKM13P6eYdCmvXCkfPw")
	queryParams := endpoint.Query()
	queryParams.Set("app_id", "Oy28e0bFNkm9ranSVF0Q")
	queryParams.Set("app_code", "nG7eRnSlGEFHGSCrzEFFd0YfCKM13P6eYdCmvXCkfPw")
	queryParams.Set("ppi", "320")
	queryParams.Set("w", "2048")
	queryParams.Set("h", "1080")
	queryParams.Set("z", "100")
	queryParams.Set("u", "1k")
	queryParams.Set("a0", geometry)

	endpoint.RawQuery = queryParams.Encode()

	response, err := http.Get(endpoint.String())
	if err != nil {
		fmt.Printf("HTTP Request failed %s\n", err)
	} else {
		f, _ := os.Create(output)
		data, _ := ioutil.ReadAll(response.Body)
		f.Write(data)
		defer f.Close()
	}
}

func main () {
	fmt.Println("Starting Application... ")

	polygon := Polygon{
		Points: []Point {
			{Latitude:6.45407, Longitude:3.39467},
			{Latitude:7.15571, Longitude:3.34509},
			{Latitude:7.401962, Longitude:3.917313},
			{Latitude:6.45407, Longitude:3.39467},
		},
	}
	GetMapWithPolygon("map.jpg", polygon.toString())
}




