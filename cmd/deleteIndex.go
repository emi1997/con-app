package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init (){
  rootCmd.AddCommand(deleteIndex)
}

//Defining new command
var deleteIndex = &cobra.Command{
  Use: "delindex",
  Short: "delete index from a given index",
  Long:  `Use this command to delete an index .`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}