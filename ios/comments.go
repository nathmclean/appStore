package ios

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type AppComments struct {
	Feed AppFeed `json:"feed"`
}

type AppFeed struct {
	Entries []Entry `json:"entry"`
}

type AuthorField struct {
	Label string `json:"label"`
}

type Entry struct {
	Author Author `json,omitifempty:"author"`
	Title TitleField `json:"title"`
	Content Content `json,omitifempty:"content"`
	Id IdLabel `json:"id"`
	Link Link `json:"link"`
}

type Link struct {
	Attributes LinkAttributes `json:"attributes"`
}

type LinkAttributes struct {
	Rel string `json:"rel"`
	Href string `json:"href"`
}

type IdLabel struct {
	Label string `json:"label"`
}

type Author struct {
	Name AuthorField `json:"name"`
	Uri AuthorField `json:"uri"`
}

type TitleField struct {
	Label string `json:"label"`
}

type Content struct {
	Label string `json:"label"`
	Attributes ContentAttributes `json:"attributes"`
}

type ContentAttributes struct {
	Type string `json:"type"`
}



func (c* Client) GetCommentsForApp(appId string) ([]Entry, error) {
	var app AppComments

	resp, err := http.Get(c.getCommentsURL(appId))
	if err != nil {
		log.Println(err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return app.Feed.Entries, err
	}

	err = json.Unmarshal(contents, &app)
	if err != nil {
		return app.Feed.Entries, err
	}

	// Find entries with no comments
	var emptyEntries []int
	for i, e := range app.Feed.Entries {
		if e.Content.Label == "" {
			emptyEntries = append(emptyEntries, i)
		}
	}

	// reverse slice so we're removing later index first
	for i := len(emptyEntries)/2-1; i >= 0; i-- {
		opp := len(emptyEntries)-1-i
		emptyEntries[i], emptyEntries[opp] = emptyEntries[opp], emptyEntries[i]
	}
	//remove entries
	for _, i := range emptyEntries {
		app.Feed.Entries = append(app.Feed.Entries[:i], app.Feed.Entries[i+1:]...)
	}
	return app.Feed.Entries, nil
}

func (c* Client) getCommentsURL(appId string) string {
	url := []string{c.CommentEndpoint, c.Country, c.CommentsPath, "id=" + appId, c.Extension }
	return strings.Join(url, "/")
}