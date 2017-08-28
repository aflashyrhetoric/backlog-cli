package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	"net/url"
	"os"
	"strings"
)

var cfgFile string
var hc = http.Client{}
var path string = "/Users/wdkevo/Nulab/cacoo-blog"
var formContentType string = "Content-Type:application/x-www-form-urlencoded"

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "backlog-cli",
	Short: "Use Backlog from the command line.",
	Long:  `to quickly create a Cobra application.`,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// FIXME: Temporary getter for Project Key
func ProjectKey() string {
	return viper.GetString("PROJECT_KEY")
}

// FIXME: Temporary getter for repository name
func Repo() string {
	return viper.GetString("REPOSITORY_NAME")
}

func initConfig() {
	if cfgFile != "" {
		fmt.Println("Config found. Loading...")
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Println("Config not found. Setting defaults...")

		viper.SetConfigName(".backlog-cli")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		// viper.SetEnvPrefix("backlog")

		// read in environment variables that match
		viper.AutomaticEnv()
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

// Endpoint returns an endpoint
func Endpoint(apiUrl string) string {
	baseURL := viper.GetString("BASE_URL")
	key := "?apiKey=" + viper.GetString("API_KEY")
	endpoint := baseURL + apiUrl + key
	return endpoint
}

func get(endpoint string) []byte {
	// func get(endpoint) string {
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

func post(endpoint string, form url.Values) []byte {
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//fmt.Printf("form was %v", form)
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

func printResponse(responseData []byte) {
	fmt.Println(string(responseData[:]))
}
