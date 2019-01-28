package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ChimeraCoder/anaconda"

	"github.com/gorilla/mux"
)

func main() {
	log.SetOutput(os.Stdout)

	router := mux.NewRouter()
	// router.HandleFunc("/hitec/crawl/tweets/mention/{account_name}/history-in-days/{days}/lang/{lang}", getTweetsFromAccountByDays).Methods("GET")
	// router.HandleFunc("/hitec/crawl/tweets/mention/{account_name}/from/{date}/lang/{lang}", getTweetsFromDate).Methods("GET")
	router.HandleFunc("/hitec/crawl/tweets/mention/{account_name}/lang/{lang}", getTweetsInLang).Methods("GET")
	router.HandleFunc("/hitec/crawl/tweets/mention/{account_name}/lang/{lang}/fast", getTweetsInLangFast).Methods("GET")

	router.HandleFunc("/hitec/crawl/tweets/hashtag/{hashtag}/lang/{lang}", getTweetsWithHashtagInLang).Methods("GET")

	log.Fatal(http.ListenAndServe(":9624", router))
}

func getTweetsFromAccountByDays(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountName := params["account_name"]
	lang := params["lang"]
	days, err := strconv.Atoi(params["days"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var anacondaTweets []anaconda.Tweet
	timeFrame := TimeFrameFromDays(days)
	if timeFrame.IsValid() {
		anacondaTweets = Crawl(lang, timeFrame, accountName, false)
	}
	tweets := TweetsFromAnacondaCrawler(anacondaTweets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

func writeToJSONFile(tweets []Tweet) {
	fmt.Printf("----\nwe collected %v tweets\n---\n", len(tweets))
	jsonData, err := json.Marshal(tweets)
	jsonFile, err := os.Create("./Tweets.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}

func getTweetsFromDate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountName := params["account_name"]
	since := params["date"]
	lang := params["lang"]

	var anacondaTweets []anaconda.Tweet
	timeFrame := TimeFrameFromSince(since)
	if timeFrame.IsValid() {
		anacondaTweets = Crawl(lang, timeFrame, accountName, false)
	}
	tweets := TweetsFromAnacondaCrawler(anacondaTweets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

func getTweetsInLang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountName := params["account_name"]
	lang := params["lang"]

	var anacondaTweets []anaconda.Tweet
	timeFrame := TimeFrame{}
	anacondaTweets = Crawl(lang, timeFrame, accountName, false)
	tweets := TweetsFromAnacondaCrawler(anacondaTweets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

func getTweetsInLangFast(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountName := params["account_name"]
	lang := params["lang"]

	var anacondaTweets []anaconda.Tweet
	timeFrame := TimeFrame{}
	anacondaTweets = Crawl(lang, timeFrame, accountName, false)
	tweets := TweetsFromAnacondaCrawler(anacondaTweets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}

func getTweetsWithHashtagInLang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hashtag := params["hashtag"]
	lang := params["lang"]

	anacondaTweets := CrawlForHashtags(lang, []string{hashtag})
	tweets := TweetsFromAnacondaCrawler(anacondaTweets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
