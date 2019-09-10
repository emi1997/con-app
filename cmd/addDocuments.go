package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init (){
  rootCmd.AddCommand(addDocument)
}

var addDocument = &cobra.Command{
  Use: "adddocument",
  Short: "Add new document to a given index",
  Long:  `Use this command to add a new document to the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}