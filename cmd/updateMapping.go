package cmd

import (
	"github.com/emi1997/con-app/client"
	"github.com/spf13/cobra"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init() {
	rootCmd.AddCommand(updateMapping)
}

//Defining new command
var updateMapping = &cobra.Command{
	Use:   "upmap",
	Short: "update mapping from a given index",
	Long:  `Use this command to update a Mapping from the index of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.UpdateMapping()
	},
}
