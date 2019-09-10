package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)
func init (){
  rootCmd.AddCommand(addIndex)
}

var addIndex = &cobra.Command{
  Use: "addindext",
  Short: "Add a new index",
  Long:  `Use this command to add a new index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}