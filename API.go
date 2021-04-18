package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type responseResults struct {
	Sunrise string `json:"sunrise"`
	Sunset  string `json:"sunset"`
}

type response struct {
	Result responseResults `json:"results"`
}

func shouldIHaveMyLightsOn(w http.ResponseWriter, r *http.Request) {

	req, _ := http.NewRequest("GET", "https://api.sunrise-sunset.org/json?lat=36.7201600&lng=2.376776", nil)
	client := &http.Client{}
	resp, _ := client.Do(req)

	b, _ := ioutil.ReadAll(resp.Body)
	defer r.Body.Close()

	var msg response
	json.Unmarshal(b, &msg)
	msg.Result.Sunset = strings.Split(msg.Result.Sunset, ":")[0]
	msg.Result.Sunrise = strings.Split(msg.Result.Sunrise, ":")[0]
	heureLevee, _ := strconv.Atoi(msg.Result.Sunrise)
	heureCouche, _ := strconv.Atoi(msg.Result.Sunset)
	heureCouche += 12
	if (time.Now().Hour() >= heureCouche) || (time.Now().Hour() <= heureLevee) {
		fmt.Fprint(w, "yes")
		return
	}
	fmt.Fprint(w, "no")
}
