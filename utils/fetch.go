package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var configFile string
var hc = http.Client{}

// Error .. an error returned from the API
type Error struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

// ErrorsList .. a list of errors returned from the API
type ErrorsList struct {
	errorList []Error
}

// GetParam ... is an array of string parameters
type GetParam map[string]int

// Get .. issues an HTTP GET request
func Get(endpoint string) []byte {
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Response
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

// GetWithParams .. issues an HTTP Get request with query parameters
func GetWithParams(endpoint string, queryParams map[string]int) []byte {
	queryString := ""

	for k, v := range queryParams {
		queryString = fmt.Sprintf("%s&%s=%d", queryString, k, v)
	}

	endpoint = endpoint + queryString

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Response
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

// Post .. issues an HTTP POST request
func Post(endpoint string, form url.Values) ([]byte, error) {

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Fetch
	response, err := hc.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData, err
}
