package cmd

import (
  //"fmt"

  "github.com/spf13/cobra"
  "client"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init (){
  rootCmd.AddCommand(bulkIndex)
}
//Defining new command
var bulkIndex = &cobra.Command{
  Use: "bulk",
  Short: "bulk index some Documents",
  Long:  `let´s you index multiple documents at once`,
  Run: func(cmd *cobra.Command, args []string) {
    client.BulkIndex()
  },
}