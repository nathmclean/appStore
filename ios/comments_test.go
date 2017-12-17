package ios

import (
	"testing"
)

func TestGetCommentsForApp(t *testing.T)  {

	client, _ := NewClient()

	cases := []struct {
		AppId       string
		FixtureFile string
	}{
		{
			AppId: "1058526204",
		},
	}

	for i, c := range cases {
		t.Run("Get Comments", func(t *testing.T) {
			comments, err := client.GetCommentsForApp(c.AppId)
			if err != nil {
				t.Fatal(i, err)
			}
			t.Log(i, "-", len(comments), "comments returned")
			t.Log(comments)
		})
	}
}

func TestGetCommentsURL(t *testing.T)  {
	client, _ := NewClient()

	url := client.getCommentsURL("1058526204")
	expected := "https://itunes.apple.com/us/rss/customerreviews/id=1058526204/json"

	if url != expected {
		t.Fatal("Expected", expected, "got", url)
	}
}
