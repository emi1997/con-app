package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init (){
  rootCmd.AddCommand(deleteDocument)
}

var deleteDocument = &cobra.Command{
  Use: "deldocument",
  Short: "delete document from a given index",
  Long:  `Use this command to delete a document from the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}