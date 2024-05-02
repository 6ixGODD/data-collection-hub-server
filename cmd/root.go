package cmd

/*
Copyright © 2024 BoChen SHEN 6godddddd@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"os"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/utils/check"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile  string // config file path (default is $HOME/.data-collection-hub-server.prob.yaml)
	port        string // port to listen on (default is 3000)
	host        string // host to listen on (default is localhost)
	logLevel    string // log level (default is info)
	tls         bool   // enable tls (default is false)
	tlsCertFile string // tls cert file path (default is "")
	tlsKeyFile  string // tls key file path (default is "")
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "data-collection-hub-server",
	Short: "Data Collection Hub Server",
	Long:  `Data Collection Hub Server designed to collect instruction data in Stanford Alpaca format.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Add your logic here
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&configFile,
		"config",
		"",
		"config file (default is $HOME/.data-collection-hub-server.prob.yaml)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&port,
		"port",
		"p",
		"",
		"port to listen on (default is 8080)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&host,
		"host",
		"H",
		"",
		"host to listen on (default is localhost)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&logLevel,
		"log-level",
		"l",
		"info",
		"log level (default is info)",
	)
	rootCmd.PersistentFlags().BoolVar(
		&tls,
		"tls",
		false,
		"enable tls (default is false)",
	)
	rootCmd.PersistentFlags().StringVar(
		&tlsCertFile,
		"tls-cert-file",
		"",
		"tls cert file path (default is \"\")",
	)
	rootCmd.PersistentFlags().StringVar(
		&tlsKeyFile,
		"tls-key-file",
		"",
		"tls key file path (default is \"\")",
	)
}

// initConfig creates a new config object and initializes it with the values from the config file and flags.
// priority: flags > config file > default values
func initConfig() {
	cfg, err := config.NewConfig()
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		var home string
		home, err = os.UserHomeDir()

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".data-collection-hub-server")
	}

	// If a config file is found, read it in.
	err = viper.ReadInConfig()
	err = viper.Unmarshal(&cfg)
	cobra.CheckErr(err)
	if !(check.IsValidAppHost(host) && check.IsValidAppPort(port)) || (host == "" && port == "") {
		cobra.CheckErr("Invalid host or port")
	}
}
