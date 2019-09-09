package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"./cmd"
)

func main()  {
	cmd.Execute()
	fmt.Println("When I´m grown upp, I´ll be an application for your console!")
}