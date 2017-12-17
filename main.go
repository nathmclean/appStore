package main

import (
	"github.com/nathmclean/appStore/ios"
	"log"
)

func main()  {
	client ,_ := ios.NewClient()

	filter := []string{
		"Games",
	}

	log.Println("fetch")
	topGrossing, err := client.GetTopGrossingApps(filter, 100)
	if err != nil {
		log.Println(err)
	}
	log.Println("print all")
	for i, app := range topGrossing {
		log.Println(i+1, app.Id)
	}

	for _, app := range topGrossing {
		log.Println("Fetching comments for", app.Id, app.Name)
		comments, err := client.GetCommentsForApp(app.Id)
		if err != nil {
			log.Println(err)
		}
		log.Println("Comments", len(comments))
	}

	//comments, err := client.GetCommentsForApp("553834731")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(len(comments))
	//
	//comprehendClient := comprehend.NewClient()
	//
	//neu := 0
	//pos := 0
	//neg := 0
	//mix := 0
	//
	//for i, c := range comments {
	//	comment := c.Content.Label
	//	log.Println("Lookup sentiment", i)
	//	sentiment, err := comprehendClient.GetSentiment(comment)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	switch sentiment.SentimentClass {
	//	case "NEUTRAL":
	//		neu = neu +1
	//	case "POSITIVE":
	//		pos = pos +1
	//	case "NEGATIVE":
	//		neg = neg +1
	//	case "MIXED":
	//		mix = mix +1
	//	default:
	//		log.Println("not found")
	//	}
	//	log.Println("sentiment:", sentiment.SentimentClass)
	//	entities, err := comprehendClient.GetEntities(comment)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	log.Println("entities:", entities)
	//
	//	phrases, _ := comprehendClient.GetKeyPhrases(comment)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	log.Println("phrases:", phrases)
	//	log.Println()
	//}
	//log.Println("neu", neu, "pos", pos, "neg", neg, "mix", mix)
}
