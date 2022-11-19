package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Version number of bwssh",
  Long:  `Version number of bwssh`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("0.1")
  },
}