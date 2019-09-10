package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init (){
  rootCmd.AddCommand(addMapping)
}

//Defining new command
var addMapping = &cobra.Command{
  Use: "addmapping",
  Short: "Add new mapping to a given index",
  Long:  `Use this command to add a new mapping to the index of your choice.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Will be implemented shortly")
  },
}