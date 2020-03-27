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

	defer resp.Body.Close()

	if err != nil {
		panic(err)
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

	defer resp.Body.Close()

	if err != nil {
		panic(err)
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
	fmt.Println("--- END INFORMATION ---\n")
}

// Global func : display total information for Total World Wide Information
func Global(g *GlobalPop) {
	fmt.Println()
	fmt.Println("--- Global Corona Virus Information ---")
	fmt.Println("Confirmed Cases:", g.Global.All.Confirmed)
	fmt.Println("Deaths:", g.Global.All.Deaths)
	fmt.Println("Recovered:", g.Global.All.Recovered)
	fmt.Println("--- END INFORMATION ---\n")
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

// State : This struct handles US state JSON info
type State struct {
	US struct {
		All struct {
			Confirmed int `json:"confirmed"`
			Recovered int `json:"recovered"`
			Deaths    int `json:"deaths"`
		} `json:"All"`
		Mississippi struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Mississippi"`
		GrandPrincess struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Grand Princess"`
		Oklahoma struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Oklahoma"`
		Delaware struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Delaware"`
		Minnesota struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Minnesota"`
		Illinois struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Illinois"`
		Arkansas struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Arkansas"`
		NewMexico struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"New Mexico"`
		Indiana struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Indiana"`
		Louisiana struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Louisiana"`
		Texas struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Texas"`
		Wisconsin struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Wisconsin"`
		Kansas struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Kansas"`
		Recovered struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Recovered"`
		AmericanSamoa struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"American Samoa"`
		Connecticut struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Connecticut"`
		VirginIslands struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Virgin Islands"`
		California struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"California"`
		PuertoRico struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Puerto Rico"`
		Georgia struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Georgia"`
		NorthDakota struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"North Dakota"`
		Pennsylvania struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Pennsylvania"`
		WestVirginia struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"West Virginia"`
		Alaska struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Alaska"`
		Tennessee struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Tennessee"`
		Missouri struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Missouri"`
		SouthDakota struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"South Dakota"`
		Colorado struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Colorado"`
		NewJersey struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"New Jersey"`
		Guam struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Guam"`
		Washington struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Washington"`
		NewYork struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"New York"`
		Nevada struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Nevada"`
		NorthernMarianaIslands struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Northern Mariana Islands"`
		DiamondPrincess struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Diamond Princess"`
		Maryland struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Maryland"`
		Idaho struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Idaho"`
		Wyoming struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Wyoming"`
		Arizona struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Arizona"`
		Iowa struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Iowa"`
		Michigan struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Michigan"`
		Utah struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Utah"`
		Virginia struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Virginia"`
		Oregon struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Oregon"`
		Montana struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Montana"`
		NewHampshire struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"New Hampshire"`
		Massachusetts struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Massachusetts"`
		SouthCarolina struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"South Carolina"`
		Vermont struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Vermont"`
		Florida struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Florida"`
		Hawaii struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Hawaii"`
		Kentucky struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Kentucky"`
		RhodeIsland struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Rhode Island"`
		Nebraska struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Nebraska"`
		Ohio struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Ohio"`
		Alabama struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Alabama"`
		NorthCarolina struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"North Carolina"`
		DistrictOfColumbia struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"District of Columbia"`
		Maine struct {
			Lat       string `json:"lat"`
			Long      string `json:"long"`
			Confirmed int    `json:"confirmed"`
			Recovered int    `json:"recovered"`
			Deaths    int    `json:"deaths"`
			Updated   string `json:"updated"`
		} `json:"Maine"`
	} `json:"US"`
}
