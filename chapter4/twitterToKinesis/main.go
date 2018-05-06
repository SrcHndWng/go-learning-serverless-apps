package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

const region = "ap-northeast-1"

func main() {
	fmt.Println("twitterToKinesis start...")

	api := createTwitterAPI()
	k := createKinesisClient()

	v := url.Values{}
	v.Set("locations", "122.87,24.84,153.01,46.80")

	ts := api.PublicStreamFilter(v)
	for {
		x := <-ts.C
		switch tweet := x.(type) {
		case anaconda.Tweet:
			fmt.Println("-----------")
			putRecord(k, tweet.Text)
		case anaconda.StatusDeletionNotice:
			// pass
		default:
			fmt.Printf("unknown type(%T) : %v \n", x, x)
		}
	}
}

func putRecord(k *kinesis.Kinesis, text string) error {
	fmt.Println(text)
	record := &kinesis.PutRecordInput{
		Data:         []byte(text),
		PartitionKey: aws.String("filter"),
		StreamName:   aws.String("twitter-to-kinesis-stream"),
	}
	out, err := k.PutRecord(record)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}

func createKinesisClient() *kinesis.Kinesis {
	s := session.New(&aws.Config{Region: aws.String(region)})
	return kinesis.New(s)
}

func createTwitterAPI() *anaconda.TwitterApi {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	return anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}

func getAccessKeys() (consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) {
	consumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	return
}
