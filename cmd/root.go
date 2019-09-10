package cmd

import (
	"fmt"
	// "net"
	"os"

	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "school",
	Short: "This is an app to store topics",
	Long: `This is an app to store topics. It stores them in an Elasticsearch index
		and can query them.`,
}

//Execute - Adds all defined commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
