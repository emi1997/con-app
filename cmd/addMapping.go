package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init (){
  rootCmd.AddCommand(addMapping)
}

var addMapping = &cobra.Command{
  Use: "addmapping",
  Short: "Add new mapping to a given index",
  Long:  `Use this command to add a new mapping to the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}