package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)


func init (){
  rootCmd.AddCommand(readDocument)
}

var readDocument = &cobra.Command{
  Use: "deldocument",
  Short: "read document from a given index",
  Long:  `Use this command to read a document from the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}