package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init (){
  rootCmd.AddCommand(deleteIndex)
}

var deleteIndex = &cobra.Command{
  Use: "delindex",
  Short: "delete index from a given index",
  Long:  `Use this command to delete an index .`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}