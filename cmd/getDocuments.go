package cmd

import (
	"github.com/emi1997/con-app/client"
	"github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init() {
	rootCmd.AddCommand(getDocument)
}

//Defining new command
var getDocument = &cobra.Command{
	Use:   "getdoc",
	Short: "get document from a given index",
	Long:  `Use this command to get a document from the index of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.GetDocument()
	},
}
