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
	"context"
	"fmt"
	"os"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/wire"
	"data-collection-hub-server/pkg/utils/check"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type cliFlags struct {
	configFile  *string // config file path (default is $HOME/.data-collection-hub-server.prob.yaml)
	port        *string // port to listen on (default is 3000)
	host        *string // host to listen on (default is localhost)
	logLevel    *string // log level (default is info)
	tls         *bool   // enable tls (default is false)
	tlsCertFile *string // tls cert file path (default is "")
	tlsKeyFile  *string // tls key file path (default is "")
}

var (
	flags = &cliFlags{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "data-collection-hub-server",
	Short: "Data Collection Hub Server",
	Long:  `Data Collection Hub Server designed to collect instruction data in Stanford Alpaca format.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		// Start the app
		fiberApp, err := wire.InitializeApp(ctx)
		if err != nil {
			panic(err)
		}
		fiberApp.Logger.SetTagInContext(ctx, logging.SystemTag)
		// Start the app
		if fiberApp.Config.BaseConfig.EnableTls {
			fiberApp.Logger.Logger.Info(
				"Starting server with TLS enabled",
				zap.String("host", fiberApp.Config.BaseConfig.AppHost),
				zap.String("port", fiberApp.Config.BaseConfig.AppPort),
				zap.String("tls_cert_file", fiberApp.Config.BaseConfig.TlsCertFile),
				zap.String("tls_key_file", fiberApp.Config.BaseConfig.TlsKeyFile),
			)
			if err := fiberApp.App.ListenTLS(
				fmt.Sprintf("%s:%s", fiberApp.Config.BaseConfig.AppHost, fiberApp.Config.BaseConfig.AppPort),
				fiberApp.Config.BaseConfig.TlsCertFile,
				fiberApp.Config.BaseConfig.TlsKeyFile,
			); err != nil {
				panic(err)
			}
		} else {
			fiberApp.Logger.Logger.Info(
				"Starting server",
				zap.String("host", fiberApp.Config.BaseConfig.AppHost),
				zap.String("port", fiberApp.Config.BaseConfig.AppPort),
			)
			if err := fiberApp.App.Listen(
				fmt.Sprintf("%s:%s", fiberApp.Config.BaseConfig.AppHost, fiberApp.Config.BaseConfig.AppPort),
			); err != nil {
				panic(err)
			}
		}
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

	rootCmd.PersistentFlags().StringVar(
		flags.configFile,
		"config",
		"",
		"config file (default is $HOME/.data-collection-hub-server.prob.yaml)",
	)
	rootCmd.PersistentFlags().StringVarP(
		flags.port,
		"port",
		"p",
		"",
		"port to listen on (default is 8080)",
	)
	rootCmd.PersistentFlags().StringVarP(
		flags.host,
		"host",
		"H",
		"",
		"host to listen on (default is localhost)",
	)
	rootCmd.PersistentFlags().StringVarP(
		flags.logLevel,
		"log-level",
		"l",
		"info",
		"log level (default is info)",
	)
	rootCmd.PersistentFlags().BoolVar(
		flags.tls,
		"tls",
		false,
		"enable tls (default is false)",
	)
	rootCmd.PersistentFlags().StringVar(
		flags.tlsCertFile,
		"tls-cert-file",
		"",
		"tls cert file path (default is \"\")",
	)
	rootCmd.PersistentFlags().StringVar(
		flags.tlsKeyFile,
		"tls-key-file",
		"",
		"tls key file path (default is \"\")",
	)
}

// initConfig creates a new config object and initializes it with the values from the config file and flags.
// priority: flags > config file > default values
func initConfig() {
	cfg, err := config.New()
	cobra.CheckErr(err)
	if flags.configFile != nil {
		// Use config file from the flag.
		viper.SetConfigFile(*flags.configFile)
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
	cobra.CheckErr(err)
	err = viper.Unmarshal(cfg)
	cobra.CheckErr(err)
	if flags.port != nil {
		if !check.IsValidAppPort(*flags.port) {
			cobra.CheckErr(fmt.Errorf("invalid port: %s", *flags.port))
		}
		cfg.BaseConfig.AppPort = *flags.port
	}
	if flags.host != nil {
		if !check.IsValidAppHost(*flags.host) {
			cobra.CheckErr(fmt.Errorf("invalid host: %s", *flags.host))
		}
		cfg.BaseConfig.AppHost = *flags.host
	}
	if flags.logLevel != nil {
		if !check.IsValidLogLevel(*flags.logLevel) {
			cobra.CheckErr(fmt.Errorf("invalid log level: %s", *flags.logLevel))
		}
		cfg.ZapConfig.Level = *flags.logLevel
	}
	if flags.tls != nil {
		cfg.BaseConfig.EnableTls = *flags.tls
	}
	if flags.tlsCertFile != nil {
		// Check if the tls cert file exists
		if _, err := os.Stat(*flags.tlsCertFile); os.IsNotExist(err) {
			cobra.CheckErr(fmt.Errorf("tls cert file does not exist: %s", *flags.tlsCertFile))
		}
		cfg.BaseConfig.TlsCertFile = *flags.tlsCertFile
	}
	if flags.tlsKeyFile != nil {
		// Check if the tls key file exists
		if _, err := os.Stat(*flags.tlsKeyFile); os.IsNotExist(err) {
			cobra.CheckErr(fmt.Errorf("tls key file does not exist: %s", *flags.tlsKeyFile))
		}
		cfg.BaseConfig.TlsKeyFile = *flags.tlsKeyFile
	}
	config.Update(cfg)
}
