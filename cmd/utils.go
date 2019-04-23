package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

// PrintResponse .. Prints out a []byte
func PrintResponse(responseData []byte) {
	fmt.Println(string(responseData[:]))
}

// Endpoint .. returns an endpoint
func Endpoint(apiURL string) string {
	// FIXME: We should just take SpaceID and build the "baseURL" from that
	baseURL := viper.GetString("BASEURL")
	key := "?apiKey=" + viper.GetString("API_KEY")
	endpoint := baseURL + apiURL + key
	return endpoint
}

// ErrorCheck .. Checks for error != nil
func ErrorCheck(err error) {
	if err != nil {
		fmt.Printf("#%v", err)
	}
}

// ErrorPanic .. Checks for errors, panics if found
func ErrorPanic(err error) {
	if err != nil {
		fmt.Printf("#%v", err)
		panic(err)
	}
}
