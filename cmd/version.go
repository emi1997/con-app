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
  Short: "Print the version number of school",
  Long:  `All software has versions. This is school's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("version: 0.0.2")
  },
}