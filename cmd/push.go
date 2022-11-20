/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Eddayy/bwssh/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Upload ssh keys to bitwarden",

	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			color.Cyan("Bwssh settings")
			settings := viper.GetViper().AllSettings()
			formattedSetting, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(formattedSetting))
		}
		bw := lib.Bw{Flags: cmd}
		bwStatus := bw.CheckStatus()
		if bwStatus == "locked" {
			if err := bw.Unlock(); err != nil {
				cobra.CheckErr(err)
				return
			}
		} else if bwStatus == "unauthenticated" {
			cobra.CheckErr(errors.New("Unauthenticated, please run bwssh init first"))
		}

		bw.ValidateFolder()

	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
