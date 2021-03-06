// Copyright © 2017 The Things Network. Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "packet-forwarder",
	Short: "The Things Network LoRa Packet Forwarder",
	Long: `The Things Network LoRa Packet Forwarder

Every build is configured to interact with a kind of
LoRa concentrator.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default \"$HOME/.pktfwd.yml\")")
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName(".pktfwd")
	viper.AddConfigPath("$HOME")
	viper.SetEnvPrefix("pktfwd")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error when reading config file:", err)
		os.Exit(1)
	}
}
