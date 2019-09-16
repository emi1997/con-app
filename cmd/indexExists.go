package cmd

import (
	"con-app/go/src/client"
  //"fmt"

  "github.com/spf13/cobra"

)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init (){
  rootCmd.AddCommand(indexExist)
}
//Defining new command
var indexExist = &cobra.Command{
  Use: "exist",
  Short: "Check if index exists",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    client.IndexExist()
  },
}