package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	// 各種キーの読み込み
	keys := ReadConfigFile()

	// Twitter client のセットアップ
	config := oauth1.NewConfig(keys.APIKey, keys.APISecretKey)
	token := oauth1.NewToken(keys.AccessToken, keys.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, _, err := client.Statuses.Update(time.Now().String(), nil)
	if err != nil {
		log.Println(err)
	}

	// Sample Stream
	params := &twitter.StreamSampleParams{
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Sample(params)
	if err != nil {
		log.Fatalln(err)
	}
	for mes := range stream.Messages {
		fmt.Println(mes)
	}
}
