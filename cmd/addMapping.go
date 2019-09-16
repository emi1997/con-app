package cmd

import (
	// "fmt"
	//"log"
	//"context"

	"github.com/emi1997/con-app/client"
	"github.com/spf13/cobra"
	//"github.com/olivere/elastic"
	//"client"
)

//calls on rootCmd from root.go with AddDocument function and passes it newly created command as argument
func init() {
	rootCmd.AddCommand(addMapping)
}

//Defining new command
var addMapping = &cobra.Command{
	Use:   "newmap",
	Short: "Add a new mapping",
	Long:  `Use this command to add a new mapping to an index of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.AddMapping()
	},
}
