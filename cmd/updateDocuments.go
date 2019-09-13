package cmd

import (
  // "fmt"

  "github.com/spf13/cobra"
  //"client"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init (){
  rootCmd.AddCommand(updateDocument)
}

//Defining new command
var updateDocument = &cobra.Command{
  Use: "updoc",
  Short: "update document from a given index",
  Long:  `Use this command to update a document in the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    //client.UpdateDocument()
  },
}