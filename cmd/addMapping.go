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
var addMapping = &cobra.Command{
  Use: "newmap",
  Short: "Add a new mapping",
  Long:  `Use this command to add a new mapping to an index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    client.AddMapping()
  },
}