package cmd

import (
	//"fmt"

	"github.com/emi1997/con-app/client"
	"github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init() {
	rootCmd.AddCommand(deleteIndex)
}

//Defining new command
var deleteIndex = &cobra.Command{
	Use:   "delind",
	Short: "delete index from a given index",
	Long:  `Use this command to delete an index .`,
	Run: func(cmd *cobra.Command, args []string) {
		client.DeleteIndex()
	},
}
