package cmd

import (
	"con-app/go/src/client"
	// "fmt"
	//"log"
	//"context"
	//"bufio"

	"github.com/spf13/cobra"
	//"github.com/olivere/elastic"
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
