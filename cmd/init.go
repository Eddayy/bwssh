/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "First command to run, will setup authentication for future access",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		fmt.Println("Enter client_id and client_secret")
		fmt.Println("Learn more: https://bitwarden.com/help/personal-api-key/")

		client_id := prompt("client_id")
		client_secret := prompt("client_secret")

		if verbose {
			fmt.Println("client_id:", client_id)
			fmt.Println("client_secret:", client_secret)
		}

		// Validate if keys are correct
		
	},
}

func prompt(arg string) string {
	prompt := promptui.Prompt{
		Label: arg,
		Mask:  '*',
	}
	result, _ := prompt.Run()
	return result
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
