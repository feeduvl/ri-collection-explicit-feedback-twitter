package main

import (
	"reflect"
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

var anacondaTweetMock anaconda.Tweet
var tweetMock Tweet
var tweetMockShouldFail Tweet

func init() {
	anacondaTweetMock.CreatedAt = "Thu May 03 20:01:03 +0000 2018"
	anacondaTweetMock.FavoriteCount = 0
	anacondaTweetMock.RetweetCount = 1
	anacondaTweetMock.FullText = "This is a test tweet"
	anacondaTweetMock.IdStr = "001"
	anacondaTweetMock.InReplyToScreenName = "WindItalia"
	anacondaTweetMock.Lang = "it"

	tweetMock.CreatedAt = 20180503
	tweetMock.CreatedAtFull = "Thu May 03 20:01:03 +0000 2018"
	tweetMock.FavoriteCount = 0
	tweetMock.RetweetCount = 1
	tweetMock.StatusID = "001"
	tweetMock.Text = "This is a test tweet"
	tweetMock.InReplyToScreenName = "WindItalia"
	tweetMock.Lang = "it"

	tweetMockShouldFail.CreatedAt = 20170503
	tweetMock.CreatedAtFull = "Thu May 03 20:01:03 +0000 2018"
	tweetMockShouldFail.FavoriteCount = 0
	tweetMockShouldFail.RetweetCount = 1
	tweetMockShouldFail.StatusID = "001"
	tweetMockShouldFail.Text = "This is a test tweet"
	tweetMockShouldFail.InReplyToScreenName = "WindItalia"
	tweetMockShouldFail.Lang = "it"
}

func TestTweetFromAnacondaCrawler(t *testing.T) {
	type args struct {
		anacondaTweet anaconda.Tweet
	}
	tests := []struct {
		name string
		args args
		want Tweet
	}{
		{"this should be identical", args{anacondaTweet: anacondaTweetMock}, tweetMock},
		// {"this should be different", args{anacondaTweet: anacondaTweetMock}, tweetMockShouldFail},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TweetFromAnacondaCrawler(tt.args.anacondaTweet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TweetFromAnacondaCrawler() = %v, want %v", got, tt.want)
			}
		})
	}
}
