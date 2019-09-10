package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init (){
  rootCmd.AddCommand(updateDocument)
}

var updateDocument = &cobra.Command{
  Use: "deldocument",
  Short: "update document from a given index",
  Long:  `Use this command to update a document in the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}