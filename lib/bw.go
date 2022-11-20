package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Bw struct {
	Session string
	Folder  Folder
	Flags   *cobra.Command
}

func (bw *Bw) ListFolders() error {
	result := bw.ExecuteCommand("list folders")

	var folders []Folder
	json.Unmarshal([]byte(result), &folders)

	templates := &promptui.SelectTemplates{

		Label:    "Select folder where ssh keys are stored",
		Active:   "  {{ .Name | green | bold }}",
		Inactive: "{{ .Name | bgBlack }}",
		Selected: "{{ .Name | green }}",
	}
	prompt := promptui.Select{
		Items:        folders,
		Templates:    templates,
		Size:         6,
		HideSelected: true,
	}
	i, _, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	bw.Folder = folders[i]
	viper.Set("folder_id", folders[i].Id)
	viper.WriteConfig()
	return nil
}

func (bw *Bw) Unlock() error {
	verbose, _ := bw.Flags.Flags().GetBool("verbose")

	prompt := promptui.Prompt{
		Label:       "Master Password",
		HideEntered: true,
		Mask:        ' ',
	}
	password, _ := prompt.Run()
	if verbose {
		fmt.Println(color.MagentaString("command:"), "Master Password")
		fmt.Println(color.CyanString("output:"), password)
	}
	session := bw.ExecuteCommand("unlock --raw --passwordenv BW_PASSWORD", "BW_PASSWORD="+password)
	if session == "Invalid master password." {
		color.Red((session))
		return errors.New(session)
	}

	viper.Set("BW_SESSION", session)
	viper.WriteConfig()
	bw.CheckStatus()
	bw.Session = session
	return nil
}

func (bw *Bw) CheckStatus() string {
	result := bw.ExecuteCommand("status")
	var bwStatus map[string]interface{}
	json.Unmarshal([]byte(result), &bwStatus)

	if bwStatus["status"].(string) != "unlocked" {
		viper.Set("BW_SESSION", "")
		viper.WriteConfig()
	}

	return bwStatus["status"].(string)
}

func (bw Bw) ExecuteCommand(command string, env ...string) string {
	verbose, _ := bw.Flags.Flags().GetBool("verbose")

	cmd := exec.Command("bw", strings.Split(command, " ")...)
	cmd.Env = os.Environ()
	bwSession := viper.Get("BW_SESSION")
	if bwSession != nil {
		cmd.Env = append(cmd.Env, "BW_SESSION="+bwSession.(string))
	}

	for _, element := range env {
		cmd.Env = append(cmd.Env, element)
	}

	cmdOutput, _ := cmd.CombinedOutput()
	result := string(cmdOutput[:])
	if verbose {
		fmt.Println(color.MagentaString("command:"), command)
		fmt.Println(color.CyanString("output:"), result)
	}
	return result
}
