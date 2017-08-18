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
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.backlog-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")



	// fmt.Println(RootCmd.PersistentFlags().Lookup(key))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		fmt.Println("Config found. Loading...")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Println("Config not found. Setting defaults...")

		// Search config in home directory with name ".backlog-cli" (without extension).
		viper.SetConfigName(".backlog-cli")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		// viper.SetEnvPrefix("backlog")
		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			// fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

// Endpoint returns an endpoint
func Endpoint(apiUrl string) string {
	baseURL:= viper.GetString("API_KEY") 
	key:= "?apiKey=" + viper.GetString("API_KEY") 
	endpoint := baseURL + apiUrl + key 
	return endpoint
}

