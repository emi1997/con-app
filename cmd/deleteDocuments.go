package cmd

import (
	"github.com/emi1997/con-app/client"
	"github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init() {
	rootCmd.AddCommand(deleteDocument)
}

//Defining new command
var deleteDocument = &cobra.Command{
	Use:   "deldoc",
	Short: "delete document from a given index",
	Long:  `Use this command to delete a document from the index of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.DeleteDocument()
	},
}
