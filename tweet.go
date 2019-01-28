package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

const createdAtFormat = "20060102"

// Tweet model
type Tweet struct {
	CreatedAt           int      `json:"created_at"`
	CreatedAtFull       string   `json:"created_at_full"`
	FavoriteCount       int      `json:"favorite_count"`
	RetweetCount        int      `json:"retweet_count"`
	Text                string   `json:"text"`
	StatusID            string   `json:"status_id"`
	UserName            string   `json:"user_name"`
	InReplyToScreenName string   `json:"in_reply_to_screen_name"`
	Hashtags            []string `json:"hashtags"`
	Lang                string   `json:"lang"`
	TweetClass          string   `json:"tweet_class"`
}

// TweetFromAnacondaCrawler is a function that transform a tweet from anaconda.Tweet to the Tweet model
func TweetFromAnacondaCrawler(anacondaTweet anaconda.Tweet) Tweet {
	//////////////////////
	///// convert anaconda.Tweet date
	//////////////////////
	var date int
	anacondaTweetCreatedAtlayout := "Mon Jan 02 15:04:05 -0700 2006"
	t, err := time.Parse(anacondaTweetCreatedAtlayout, anacondaTweet.CreatedAt)
	if err != nil {
		fmt.Println(err)
		date = -1
	}
	date, _ = strconv.Atoi(t.Format(createdAtFormat))
	var hashtags []string
	for _, entityTag := range anacondaTweet.Entities.Hashtags {
		hashtags = append(hashtags, entityTag.Text)
	}

	// create the tweet
	return Tweet{
		CreatedAt:           date,
		CreatedAtFull:       anacondaTweet.CreatedAt,
		FavoriteCount:       anacondaTweet.FavoriteCount,
		RetweetCount:        anacondaTweet.RetweetCount,
		Text:                anacondaTweet.FullText,
		StatusID:            anacondaTweet.IdStr,
		InReplyToScreenName: anacondaTweet.InReplyToScreenName,
		Hashtags:            hashtags,
		UserName:            anacondaTweet.User.ScreenName,
		Lang:                anacondaTweet.Lang,
	}
}

// TweetsFromAnacondaCrawler is a function that transform a list of tweets from anaconda.Tweet to the Tweet model
func TweetsFromAnacondaCrawler(anacondaTweets []anaconda.Tweet) []Tweet {
	var tweets []Tweet
	for _, anacondaTweet := range anacondaTweets {
		tweets = append(tweets, TweetFromAnacondaCrawler(anacondaTweet))
	}

	return tweets
}
