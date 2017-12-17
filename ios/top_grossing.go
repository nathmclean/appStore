package ios

import (
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"strconv"
)
type TopGrossing struct {
	Feed TopGrossingFeed `json:"feed"`
}

type TopGrossingFeed struct {
	Results []App `json:"results"`
}

type App struct {
	ArtistName string `json:"artistName"`
	Name string `json:"name"`
	Genres []Genre `json:"genres"`
	Id string `json:"id"`
}

type Genre struct {
	Name string `json:"name"`
}

// Retrieves a list of Top Grossing App with a filter
// An empty filter slice implies no filter
// results is the number of top grossing to search through,
// not necessarily the number returned (if a filter is used)
func (c* Client) GetTopGrossingApps(genres []string, numberOfApps int) ([]App, error) {
	var topGrossingApps TopGrossing
	var filteredApps []App

	numString := strconv.Itoa(numberOfApps)

	resp, err := http.Get(c.getTopGrossingUrl(numString))
	if err != nil {
		return topGrossingApps.Feed.Results, err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return topGrossingApps.Feed.Results, err
	}

	err = json.Unmarshal(contents, &topGrossingApps)
	if err != nil {
		return topGrossingApps.Feed.Results, err
	}

	if len(genres) != 0 {
		for _, app := range topGrossingApps.Feed.Results {
			if applyFilter(genres, app) {
				filteredApps = append(filteredApps, app)
			}
		}
	}

	return filteredApps, nil
}

func applyFilter(filters []string, app App) bool {

	if len(filters) == 0 {
		return true
	}

	for _, filter := range filters {
		for _, genre := range app.Genres {
			if strings.ToLower(filter) == strings.ToLower(genre.Name) {
				return true
			}
		}
	}

	return false
}

func (c* Client) getTopGrossingUrl(quantity string) string {
	url := []string{c.TopGrossingEndpoint, quantity, c.TopGrossingExtension}
	log.Println(strings.Join(url, "/"))
	return strings.Join(url, "/")
}