package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/viper"
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

// truncate .. truncates a string
func Truncate(s string, max int) string {
	if len(s) >= max {
		return s[0:max]
	}
	return s
}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

// Prints if debug mode is on
func debugPrint(format string, a ...interface{}) {
	if viper.GetString("DEBUG_MODE") == "true" {
		fmt.Printf(format, a...)
	}
}
