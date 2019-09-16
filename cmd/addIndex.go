package cmd

import (
	"github.com/emi1997/con-app/client"
	"github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init() {
	rootCmd.AddCommand(addIndex)
}

//Defining new command
var addIndex = &cobra.Command{
	Use:   "newind",
	Short: "Add a new index",
	Long:  `Use this command to add a new index of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.AddIndex()
	},
}
