/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/Eddayy/bwssh/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup authentication and designated ssh-agent folder ",

	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			color.Cyan("Bwssh settings")
			settings := viper.GetViper().AllSettings()
			formattedSetting, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(formattedSetting))
		}
		bw := lib.Bw{Flags: cmd}

		// check if login already
		bwStatus := bw.CheckStatus()
		if bwStatus == "locked" {
			if err := bw.Unlock(); err != nil {
				cobra.CheckErr(err)
				return
			}
		} else if bwStatus == "unauthenticated" {
			color.Red(bwStatus)
			fmt.Println("Start login flow")
			var loginCommandBuffer bytes.Buffer
			loginCommand := exec.Command("bw", "login", "--raw")
			loginCommand.Stdin = os.Stdin
			loginCommand.Stdout = io.MultiWriter(&loginCommandBuffer, os.Stdout)
			loginCommand.Stderr = io.MultiWriter(&loginCommandBuffer, os.Stderr)
			_ = loginCommand.Run()
			loginCommandOutput := strings.Split(string(loginCommandBuffer.String()), "\n")

			bw.Session = loginCommandOutput[len(loginCommandOutput)-1]
			color.Green("\nAuthenticated")

		} else if bwStatus != "unlocked" {
			cobra.CheckErr(bwStatus)
			return
		}

		if err := bw.SelectFolder(); err != nil {
			cobra.CheckErr(err)
			return
		}
		color.Green("Setup all done!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
