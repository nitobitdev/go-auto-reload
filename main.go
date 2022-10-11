package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"text/template"
)

func getValueForWind() (int, string) {
	min := 1
	max := 100

	wind := rand.Intn(max-min) + min
	fmt.Println("Wind: ", wind)

	if wind > 15 {
		fmt.Println("Status Wind: Bahaya")
		return wind, "Bahaya"
	} else if wind > 7 && wind <= 15 {
		fmt.Println("Status Wind: Siaga")
		return wind, "Siaga"
	} else {
		fmt.Println("Status Wind: Aman")
		return wind, "Aman"
	}
}

func getValueForWater() (int, string) {
	min := 1
	max := 100
	water := rand.Intn(max-min) + min
	fmt.Println("Water: ", water)
	if water > 8 {
		fmt.Println("Status Water: Bahaya")
		return water, "Bahaya"
	} else if water > 5 && water <= 8 {
		fmt.Println("Status Water: Siaga")
		return water, "Siaga"
	} else {
		fmt.Println("Status Water: Aman")
		return water, "Aman"
	}
}

type StatusKondisi struct {
	Wind        int    `json:"wind"`
	Water       int    `json:"water"`
	StatusWind  string `json:"status_wind"`
	StatusWater string `json:"status_water"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	wind, statusWind := getValueForWind()
	water, statusWater := getValueForWater()

	dataDisplay := StatusKondisi{
		Wind:        wind,
		Water:       water,
		StatusWind:  statusWind,
		StatusWater: statusWater,
	}

	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, dataDisplay)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", handler)
	panic(http.ListenAndServe(":8080", nil))
}
