package lib

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Bw struct {
	Session string
	Flags   *cobra.Command
}

func (bw Bw) ListFolders() string {
	result := bw.ExecuteCommand("list folders --session " + bw.Session)

	return result
}

func (bw Bw) CheckLogin() (bool, error) {
	loginStatus := bw.ExecuteCommand("login --check")
	if loginStatus == "You are logged in!" {
		color.Green(loginStatus)
	} else if loginStatus == "You are not logged in." {
		color.Red(loginStatus)
	} else {
		color.Red(loginStatus)
		return false, errors.New(loginStatus)
	}
	return loginStatus == "You are logged in!", nil
}

func (bw Bw) ExecuteCommand(command string) string {
	verbose, _ := bw.Flags.Flags().GetBool("verbose")
	cmd := exec.Command("bw", strings.Split(command, " ")...)
	cmdOutput, _ := cmd.CombinedOutput()
	result := string(cmdOutput[:])
	if verbose {
		fmt.Println("command:", command)
		fmt.Println("output:", result)
	}
	return result
}
