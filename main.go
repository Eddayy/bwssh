/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/Eddayy/bwssh/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.SetConfigFile(home + "/.bwssh")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	cmd.Execute()
}
