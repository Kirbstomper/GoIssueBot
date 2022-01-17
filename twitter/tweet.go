package twitter

import (
	"bytes"
	"log"
	"os"

	"github.com/dghubble/oauth1"
)

// Using twitter API to post a tweet to the bot account

var (
	CONSUMER_KEY,
	CONSUMER_SECRET,
	API_KEY,
	API_SECRET string
)

func init() {
	CONSUMER_KEY = os.Getenv("CONSUMER_KEY")
	CONSUMER_SECRET = os.Getenv("CONSUMER_SECRET")
	API_KEY = os.Getenv("API_KEY")
	API_SECRET = os.Getenv("API_SECRET")
}

func SendTweet(text string) {
	config := oauth1.NewConfig(CONSUMER_KEY, CONSUMER_SECRET)
	token := oauth1.NewToken(API_KEY, API_SECRET)

	httpClient := config.Client(oauth1.NoContext, token)

	req := `{
		"text": "` + text + `"
	}`
	resp, _ := httpClient.Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewBufferString(req))
	defer resp.Body.Close()

	log.Println("Done!")
}
