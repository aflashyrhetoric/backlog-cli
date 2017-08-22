package cmd

import (
"fmt"
"os"

"github.com/spf13/cobra"
"github.com/spf13/viper"
)

var cfgFile string
var apiKey string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "backlog-cli",
	Short: "Use Backlog from the command line.",
	Long: `to quickly create a Cobra application.`,
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
	}
}

func Prep() {
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Endpoint returns an endpoint
func Endpoint(apiUrl string) string {
	baseURL:= viper.GetString("API_KEY") 
	key:= "?apiKey=" + viper.GetString("API_KEY") 
	endpoint := baseURL + apiUrl + key 
	return endpoint
}

