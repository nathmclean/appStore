package ios

const commentsEndpoint = "https://itunes.apple.com"
const topGrossingEndpoint = "https://rss.itunes.apple.com/api/v1/us/ios-apps/top-grossing/all"
const topGrossingExtension = "explicit.json"
const country = "us"
const extension  = "json"
const commentsPath =  "rss/customerreviews"

type Client struct {
	CommentEndpoint string
	TopGrossingEndpoint string
	TopGrossingExtension string
	Country         string
	CommentsPath    string
	Extension       string
}

func NewClient() (Client, error) {
	conf := Client{
		CommentEndpoint: commentsEndpoint,
		TopGrossingEndpoint: topGrossingEndpoint,
		TopGrossingExtension: topGrossingExtension,
		Country:         country,
		CommentsPath:    commentsPath,
		Extension:       extension,
	}

	return conf, nil
}