package cmd

import (
  // "fmt"
  //"log"
  //"context"

  "github.com/spf13/cobra"
  //"github.com/olivere/elastic"

  "client"
)

//Defining new command
var addIndex = &cobra.Command{
  Use: "addindex",
  Short: "Add a new index",
  Long:  `Use this command to add a new index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    client.AddIndex()
  },
}