/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/Eddayy/bwssh/lib"
	"github.com/fatih/color"
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

		var session string
		bw := lib.Bw{Flags: cmd}

		// check if login already

		isLogin, err := bw.CheckLogin()
		if err != nil {
			return
		}

		if isLogin {
			prompt := promptui.Prompt{
				Label: "Master password",
				Mask:  ' ',
			}
			password, _ := prompt.Run()
			if verbose {
				fmt.Println("Password: " + password)
			}
			cmd_line := "unlock --raw --passwordenv BW_PASSWORD"
			cmd := exec.Command("bw", strings.Split(cmd_line, " ")...)
			cmd.Env = os.Environ()
			cmd.Env = append(cmd.Env, "BW_PASSWORD="+password)
			cmdOutput, _ := cmd.CombinedOutput()
			session = string(cmdOutput[:])
			if session == "Invalid master password." {
				color.Red((session))
				return
			}
			if verbose {
				fmt.Println("command:", cmd_line)
				fmt.Println("output:", session)
			}
		} else {
			var loginCommandBuffer bytes.Buffer
			loginCommand := exec.Command("bw", "login", "--raw")
			loginCommand.Stdin = os.Stdin
			loginCommand.Stdout = io.MultiWriter(&loginCommandBuffer, os.Stdout)
			loginCommand.Stderr = io.MultiWriter(&loginCommandBuffer, os.Stderr)
			_ = loginCommand.Run()
			loginCommandOutput := strings.Split(string(loginCommandBuffer.String()), "\n")

			session = loginCommandOutput[len(loginCommandOutput)-1]
		}
		bw.Session = session

		fmt.Println("test", bw.ListFolders())
		// bw_login.Stdin = os.Stdin
		// bw_login.Stdout = os.Stdout
		// bw_login.Stderr = os.Stderr

		// Validate if keys are correct
		//#color.Red("client_id or client_secret was incorrect!")
	},
}

func prompt(arg string) string {
	prompt := promptui.Prompt{
		Label: arg,
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
