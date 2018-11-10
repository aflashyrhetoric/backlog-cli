package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cfgFile string
var hc = http.Client{}

// Endpoint returns an endpoint
func Endpoint(apiURL string) string {
	baseURL := viper.GetString("BASE_URL")
	key := "?apiKey=" + viper.GetString("API_KEY")
	endpoint := baseURL + apiURL + key
	return endpoint
}

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

func Post(endpoint string, form url.Values) []byte {
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//fmt.Printf(rform was %v", form)
	//fmt.Println(endpoint)

	// Fetch
	response, err := hc.Do(req)
	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
