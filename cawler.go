package main

import (
	"fmt"
	"log"
	"math"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

var (
	consumerKey    = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")
	accessKey      = os.Getenv("ACCESS_KEY")
	accessSecret   = os.Getenv("ACCESS_SECRET")
)

// Crawl returns a list of tweets that are addressed to a certain account in a given timeframe and language
// paginate==false ->  no pagination. Crawl until we get blocked or the end of the 'page' is reached
func Crawl(l string, t TimeFrame, accountName string, paginate bool) []anaconda.Tweet {
	var tweets []anaconda.Tweet
	var tweetsSize int
	var tweetMaxID int64 = math.MaxInt64 // used for pagination
	canPaginateTweets := true            // used for pagination
	var crawlLimitExceeded int           // used for pagination

	api := anaconda.NewTwitterApiWithCredentials(accessKey, accessSecret, consumerKey, consumerSecret)
	for canPaginateTweets { // pagination
		query, vals := buildQuery(l, 450, t, accountName, tweetMaxID)
		log.Printf("query: %v\n", query)
		searchResult, err := api.GetSearch(query, vals)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("start crawling\n")
		for _, tweet := range searchResult.Statuses {
			tweets = append(tweets, tweet)
			if tweet.Id < tweetMaxID {
				tweetMaxID = tweet.Id
			}
		}
		canPaginateTweets = paginate
		if len(tweets) == tweetsSize { // exit condition, no new tweets added
			if crawlLimitExceeded <= 1 && canPaginateTweets {
				crawlLimitExceeded++
				log.Printf("sleep 30 seconds\n")
				time.Sleep(30 * time.Second) // sleep for 30 seconds in case we reached our crawl limit
			} else {
				canPaginateTweets = false
			}
		} else { // continue crawling
			tweetsSize = len(tweets)
		}
	}
	return tweets
}

func buildQuery(l string, count int, t TimeFrame, accountName string, tweetMaxID int64) (string, map[string][]string) {
	var query string
	if t.IsValid() {
		query = fmt.Sprintf("(to:%s) since:%s until:%s&src=typed_query", accountName, t.since, t.until)
	} else {
		query = fmt.Sprintf("to:%s max_id:%d", accountName, tweetMaxID)
	}

	vals := url.Values{}
	vals.Set("count", strconv.Itoa(count))
	vals.Set("l", l)
	vals.Set("lang", l)

	return query, vals
}

func CrawlForHashtags(l string, hashtags []string) []anaconda.Tweet {
	var tweets []anaconda.Tweet

	api := anaconda.NewTwitterApiWithCredentials(accessKey, accessSecret, consumerKey, consumerSecret)

	query, vals := buildHashtagQuery(l, 450, hashtags)
	searchResult, err := api.GetSearch(query, vals)
	if err != nil {
		log.Fatal(err)
	}
	tweets = append(tweets, searchResult.Statuses...)

	return tweets
}

func buildHashtagQuery(l string, count int, hashtags []string) (string, map[string][]string) {
	var query string
	for index, hashtag := range hashtags {
		if string(hashtag[0]) != "#" {
			hashtag = "#" + hashtag
		}
		if index == 0 {
			query = hashtag
		} else {
			query = query + " OR " + hashtag
		}
	}

	vals := url.Values{}
	vals.Set("count", strconv.Itoa(count))
	vals.Set("l", l)
	vals.Set("lang", l)

	return query, vals
}

func AccountNameExists(accountName string) bool {
	api := anaconda.NewTwitterApiWithCredentials(accessKey, accessSecret, consumerKey, consumerSecret)
	users, err := api.GetUsersLookup(accountName, url.Values{})

	if err != nil || len(users) == 0 {
		return false
	}
	return true
}
