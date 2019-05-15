package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var router *mux.Router
var tweets []interface{}

func TestMain(m *testing.M) {
	fmt.Println("--- Start Tests")
	setupRouter()

	// run the test cases defined in this file
	retCode := m.Run()

	// call with result of m.Run()
	os.Exit(retCode)
}

func setupRouter() {
	fmt.Println("SetUp")

	router = mux.NewRouter()
	// router.HandleFunc("/hitec/crawl/tweets/mention/{account_name}/history-in-days/{days}/lang/{lang}", getTweetsFromAccountByDays).Methods("GET")
	// router.HandleFunc("/hitec/crawl/tweets/mention/{account_name}/from/{date}/lang/{lang}", getTweetsFromDate).Methods("GET")
	router.HandleFunc("/hitec/crawl/tweets/mention/{account_name}/lang/{lang}/fast", getTweetsInLangFast).Methods("GET")
	router.HandleFunc("/hitec/crawl/tweets/hashtag/{hashtag}/lang/{lang}", getTweetsWithHashtagInLang).Methods("GET")
}

func buildRequest(method, endpoint string, payload io.Reader, t *testing.T) *http.Request {
	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}

	return req
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

// func TestGetTweetsFromAccountByDays(t *testing.T) {
// 	fmt.Println("start TestTweetsFromAccountByDays")
// 	var method = "GET"
// 	var endpoint = "/hitec/crawl/tweets/mention/%s/history-in-days/%s/lang/%s"

// 	/*
// 	 * test for Success
// 	 */
// 	endpointSucess := fmt.Sprintf(endpoint, "VodafoneUK", "5", "en")
// 	req := buildRequest(method, endpointSucess, nil, t)
// 	rr := executeRequest(req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
// 	}

// 	err := json.NewDecoder(rr.Body).Decode(&tweets)
// 	if err != nil {
// 		t.Errorf("Did not receive a proper formed json")
// 	}

// 	if len(tweets) <= 0 {
// 		t.Errorf("response length differs. Expected %s .\n Got %d instead", "number greater than 0", len(tweets))
// 	}
// }
// func TestGetTweetsFromDate(t *testing.T) {
// 	fmt.Println("start TestTweetsFromDate")
// 	var method = "GET"
// 	var endpoint = "/hitec/crawl/tweets/mention/%s/from/%s/lang/%s"

// 	/*
// 	 * test for Success
// 	 */
// 	endpointSucess := fmt.Sprintf(endpoint, "WindItalia", "2018-05-03", "it")
// 	req := buildRequest(method, endpointSucess, nil, t)
// 	rr := executeRequest(req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
// 	}

// 	err := json.NewDecoder(rr.Body).Decode(&tweets)
// 	if err != nil {
// 		t.Errorf("Did not receive a proper formed json")
// 	}

// 	if len(tweets) <= 0 {
// 		t.Errorf("response length differs. Expected %s .\n Got %d instead", "number greater than 0", len(tweets))
// 	}
// }

func TestGetTweetsInLangFast(t *testing.T) {
	fmt.Println("start TestGetTweetsInLangFast")
	var method = "GET"
	var endpoint = "/hitec/crawl/tweets/mention/%s/lang/%s/fast"

	/*
	 * test for Success
	 */
	endpointSucess := fmt.Sprintf(endpoint, "WindItalia", "it")
	req := buildRequest(method, endpointSucess, nil, t)
	rr := executeRequest(req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	err := json.NewDecoder(rr.Body).Decode(&tweets)
	if err != nil {
		t.Errorf("Did not receive a proper formed json")
	}

	if len(tweets) <= 0 {
		t.Errorf("response length differs. Expected %s .\n Got %d instead", "number greater than 0", len(tweets))
	}
}

func TestGetTweetsWithHashtagInLang(t *testing.T) {
	fmt.Println("start TestGetTweetsWithHashtagInLang")
	var method = "GET"
	var endpoint = "/hitec/crawl/tweets/hashtag/%s/lang/%s"

	/*
	 * test for Success
	 */
	endpointSucess := fmt.Sprintf(endpoint, "coffee", "en")
	req := buildRequest(method, endpointSucess, nil, t)
	rr := executeRequest(req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	err := json.NewDecoder(rr.Body).Decode(&tweets)
	if err != nil {
		t.Errorf("Did not receive a proper formed json")
	}

	if len(tweets) <= 0 {
		t.Errorf("response length differs. Expected %s .\n Got %d instead", "number greater than 0", len(tweets))
	}
}
func TestGetAccountNameExists(t *testing.T) {
	fmt.Println("start TestGetAccountNameExists")
	var method = "GET"
	var endpoint = "/hitec/crawl/tweets/%s/exists"

	/*
	 * test for Success
	 */
	endpointSucess := fmt.Sprintf(endpoint, "VodafoneUK")
	req := buildRequest(method, endpointSucess, nil, t)
	rr := executeRequest(req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	/*
	 * test for Failure
	 */
	endpointFailure := fmt.Sprintf(endpoint, "VodafoneUK")
	req = buildRequest(method, endpointFailure, nil, t)
	rr = executeRequest(req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusInternalServerError, status)
	}
}
