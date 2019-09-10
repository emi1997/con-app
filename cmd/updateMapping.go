package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)


func init (){
  rootCmd.AddCommand(updateMapping)
}

var updateMapping = &cobra.Command{
  Use: "delmapping",
  Short: "update mapping from a given index",
  Long:  `Use this command to update a Mapping from the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}