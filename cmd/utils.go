package cmd

import (
	"fmt"
)

// PrintResponse .. Prints out a []byte
func PrintResponse(responseData []byte) {
	fmt.Println(string(responseData[:]))
}

// Endpoint .. returns an endpoint
func Endpoint(apiURL string) string {
	// FIXME: We should just take SpaceID and build the "baseURL" from that
	endpoint := fmt.Sprintf("%s%s?apiKey=%s", GlobalConfig.BaseURL, apiURL, GlobalConfig.APIKey)
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
