package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GrabHTTPJSONState : Grab response from API and parse JSON
func GrabHTTPJSONState(p *State) {

	url := "https://covid-api.mmediagroup.fr/v1/cases"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {

		defer resp.Body.Close()
	}

	// read json http response
	JSONDataHTTP, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(JSONDataHTTP, &p)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}

// GrabHTTPJSONGlobal : Grab response from API and parse JSON for GLOBAL info
func GrabHTTPJSONGlobal(g *GlobalPop) {

	url := "https://covid-api.mmediagroup.fr/v1/cases"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
	}

	// read json http response
	JSONDataHTTP, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(JSONDataHTTP, &g)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}

// Tennessee func : display Tennessee State Corona Virus Information
func Tennessee(p *State) {

	fmt.Println("--- Tennessee Corona Virus Information ---")
	fmt.Println("Confirmed Cases:", p.US.Tennessee.Confirmed)
	fmt.Println("Recovered:", p.US.Tennessee.Recovered)
	fmt.Println("Deaths:", p.US.Tennessee.Deaths)
	fmt.Println("Updated:", p.US.Tennessee.Updated)
	fmt.Println("--- END INFORMATION ---")
}

// Global func : display total information for Total World Wide Information
func Global(g *GlobalPop) {
	fmt.Println()
	fmt.Println("--- Global Corona Virus Information ---")
	fmt.Println("Confirmed Cases:", g.Global.All.Confirmed)
	fmt.Println("Deaths:", g.Global.All.Deaths)
	fmt.Println("Recovered:", g.Global.All.Recovered)
	fmt.Println("--- END INFORMATION ---")
}

// PollAPIState : Poll the API every so often to display updated Corona Virus info, this func can be used in GoRoutine & for loop to stay alive as a server.
func PollAPIState(s *State, g *GlobalPop) {

	GrabHTTPJSONState(s)
	GrabHTTPJSONGlobal(g)
	Global(g)
	Tennessee(s)
}

// Main Func
func main() {

	var s State
	var g GlobalPop

	GrabHTTPJSONState(&s)
	PollAPIState(&s, &g)

}

// GlobalPop : This struct Handles Global JSON info
type GlobalPop struct {
	Global struct {
		All struct {
			Population int64 `json:"population"`
			Confirmed  int   `json:"confirmed"`
			Recovered  int   `json:"recovered"`
			Deaths     int   `json:"deaths"`
		} `json:"All"`
	} `json:"Global"`
}
